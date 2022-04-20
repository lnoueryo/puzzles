package controller

import (
	"backend/config"
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/mail"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"html/template"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"cloud.google.com/go/storage"
	"golang.org/x/net/websocket"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]string
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Flash     string
	Warning   string
	Error     string
	JSON      []byte
	Users     []models.User
	Session   models.Session
}

var infolog *log.Logger
var errorlog *log.Logger
var DB *gorm.DB
var ws_array []*websocket.Conn // *websocket.Connを入れる配列
var email mail.Mail
var origin string
var allowOrigin string
var credentialsPath string


func init() {
	infolog = config.App.InfoLog
	errorlog = config.App.ErrorLog
	DB = config.DB
	email = config.App.Email
	origin = config.App.Origin
	allowOrigin = config.App.AllowOrigin
	credentialsPath = config.App.CredentialsPath
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData) {
	var tc map[string]*template.Template
	if config.App.UseCache {
		tc = config.App.TemplateCache
	} else {
		tc, _ = config.CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		errorlog.Print("could not get template")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		errorlog.Print("could not get template")
	}
}

func StoreImage(dir string, base64Data string) (string, error) {
	if credentialsPath != "" {
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
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsPath))
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