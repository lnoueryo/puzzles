package csv

import (
	"archive/zip"
	"backend/models"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	// "io"
	// "os"
	"strconv"
	"time"
)



func DLActivity(activities []models.Activity) *bytes.Buffer {
	rows := [][]string{{"id", "user_id", "project_id", "content_id", "created_at", "updated_at"}}
	for _, activity := range activities {
		row := []string{}
		ID := strconv.Itoa(activity.ID)
		UserID := strconv.Itoa(activity.UserID)
		ProjectID := strconv.Itoa(activity.ProjectID)
		ContentID := strconv.Itoa(activity.ContentID)
		CreatedAt := activity.CreatedAt.String()
		UpdatedAt := activity.UpdatedAt.String()
		row = append(row, ID, UserID, ProjectID, ContentID, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	buf := zipByte(b)
	return buf
}

func zipByte(b *bytes.Buffer) *bytes.Buffer {
    downloadedFileName := "hello.csv"
    // zipFileName := "a.zip"

    body, err := ioutil.ReadAll(b)
    buf := new(bytes.Buffer)
    w := zip.NewWriter(buf)
    fh := &zip.FileHeader{
        Name:     downloadedFileName,
        Modified: time.Now(),
        Method:   8,
    }
    f, err := w.CreateHeader(fh)
    if err != nil {
        fmt.Print(err)
    }
    if _, err := f.Write(body); err != nil {
        fmt.Print(err)
    }
    if err != nil {
		fmt.Print(err)
    }
    // file, err := os.Create(zipFileName)
    // if err != nil {
	// 	fmt.Print(err)
    // }
    // if _, err = io.Copy(file, buf); err != nil {
	// 	fmt.Print(err)
    // }
    // file.Close()
	w.Close()
	return buf
}