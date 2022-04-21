package config

import (
	"backend/modules/crypto"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)


func StoreImage(dir string, base64Data string) (string, error) {
	if App.CredentialsPath != "" {
		filename, err := ToGCS(dir, base64Data)
		if err != nil {
			return filename, err
		}
		return filename, nil
	}
	randStr, _ := crypto.MakeRandomStr(20)
	// サーバー側に保存するために空ファイルを作成
    coI := strings.Index(base64Data, ",")
    rawImage := base64Data[coI+1:]

    // Encoded Image DataUrl //
    unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))

    res := bytes.NewReader(unbased)
	var fileName string
	switch strings.TrimSuffix(base64Data[5:coI], ";base64") {
	case "image/png":
		extension := ".png"
		fileName = randStr + extension
		saveImage, err := os.Create("./upload/" + dir + "/" + fileName)
		if err != nil {
			message := "could't upload the file"
			err := errors.New(message)
			return fileName, err
		}
		defer saveImage.Close()
		pngI, _ := png.Decode(res)
        png.Encode(saveImage, pngI)
	case "image/jpeg":
		extension := ".jpeg"
		fileName = randStr + extension
		saveImage, err := os.Create("./upload/" + dir + "/" + fileName)
		if err != nil {
			message := "could't upload the file"
			err := errors.New(message)
			return fileName, err
		}
		defer saveImage.Close()
		jpgI, _ := jpeg.Decode(res)
        jpeg.Encode(saveImage, jpgI, &jpeg.Options{Quality: 75})
	}
	return fileName, nil
}

func ToGCS(dir string, base64Data string) (string, error) {
	coI := strings.Index(base64Data, ",")
	randStr, _ := crypto.MakeRandomStr(20)
	filename := randStr + "."
	switch strings.TrimSuffix(base64Data[5:coI], ";base64") {
	case "image/png":
		filename += "png"
	case "image/jpeg":
		filename += "jpeg"
	}
	b64data := base64Data[strings.IndexByte(base64Data, ',')+1:]
    decodedImage, err := base64.StdEncoding.DecodeString(b64data)
    if err != nil {
        errorlog.Print(err)
		message := "could't upload the file"
		err := errors.New(message)
		return filename, err
    }
	objectPath := dir + "/" + filename
	err = StoreImageToGCS(decodedImage, objectPath)
	if err != nil {
		return filename, err
	}

	return filename, nil
}

func StoreImageToGCS(bImage []byte, path string) error {
    f := bytes.NewReader(bImage)
	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(App.CredentialsPath))
	if err != nil {
		errorlog.Print(err)
		message := "could't upload the file"
		err := errors.New(message)
		return err
	}
	// オブジェクトのReaderを作成
	bucketName := "puzzle-media"
	writer := client.Bucket(bucketName).Object(path).NewWriter(ctx)
	defer writer.Close()
	
	// 書き込み
    if _, err := io.Copy(writer, f); err != nil {
        errorlog.Print(err)
		message := "could't upload the file"
		err := errors.New(message)
		return err
    }
	return nil
}

func StoreBinaryImage(bImage []byte, path string) error {
    file, err := os.Create("upload/" + path)
    if err != nil {
        return err
    }
    defer file.Close()

    file.Write(bImage)
	return nil
}

func DeleteImage(name string, dir string) {
	if App.CredentialsPath != "" {
		DeleteOnGCS(name, dir)
		return
	}
	path := "upload/" + dir + "/" + name
	os.Remove(path)
}

func DeleteOnGCS(name string, dir string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(App.CredentialsPath))
	if err != nil {
		errorlog.Print(err)
	}
	// オブジェクトのReaderを作成
	bucketName := "puzzle-media"
	path := "/" + dir + "/" + name
	src := client.Bucket(bucketName).Object(path)
	if err := src.Delete(ctx); err != nil {
		errorlog.Print(err)
    }
}