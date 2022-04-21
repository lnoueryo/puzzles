package controller

import (
	"backend/config"
	"backend/models"
	"backend/modules/mail"
	"bytes"
	"html/template"
	"log"
	"net/http"
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
var credentialsPath string
var StoreImage = config.StoreImage
var StoreImageToGCS = config.StoreImageToGCS
var StoreBinaryImage = config.StoreBinaryImage


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
