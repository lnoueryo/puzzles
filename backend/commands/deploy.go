package commands

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

var requiredDirectory = []string{"config", "controller", "credentials", "models", "modules", "public", "routes", "templates"}
var requiredFiles = []string{"Dockerfile", "go.mod", "go.sum", "gorm.db", ".env", "server/main.go", "cloudbuild.yaml"}

func Deploy() {
	for _, value := range requiredDirectory {
		Dir(value, "backend/"+value)
	}
	for _, value := range requiredFiles {
		if value == "server/main.go" {
			File(value, "backend/main.go")
		}
		File(value, "backend/"+value)
	}
	err := os.Chdir("./backend")
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(currDir)
	cmd := exec.Command("gcloud", "builds", "submit", "--config", "cloudbuild.yaml", ".")
	// cmd := exec.Command("gcloud", "beta", "run", "deploy", "--source", ".")
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		tmp := make([]byte, 1024)
		stdout.Read(tmp)
		fmt.Print(string(tmp))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
}

func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
