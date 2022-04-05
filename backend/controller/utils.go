package controller

import (
	"backend/config"
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/mail"
	"bytes"
	"encoding/base64"
	"errors"
	"html/template"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/websocket"
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


func init() {
	infolog = config.App.InfoLog
	errorlog = config.App.ErrorLog
	DB = config.DB
	email = config.App.Email
	origin = config.App.Origin
	allowOrigin = config.App.AllowOrigin
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
