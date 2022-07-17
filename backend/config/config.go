package config

import (
	"backend/modules/mail"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	UseCache 		bool
	TemplateCache 	map[string]*template.Template
	InfoLog			*log.Logger
	ErrorLog		*log.Logger
	InProduction	bool
	Addr			string
	Static			string
	Media			string
	Host			string
	Origin			string
	AllowOrigin		string
	Email			mail.Mail
	ProjectID		string
}

var App AppConfig
var infolog *log.Logger
var errorlog *log.Logger

// envによる環境の決定
func init() {

    err := godotenv.Load(".env.dev"); if err != nil {
		configureProdSettings()
	} else {
		configureLocalSettings()
	}
	commonSettings()
}

// 本番、開発環境共通のセッティング
func commonSettings() {
	// log
	infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App.InfoLog = infolog
	App.ErrorLog = errorlog

	// file path
	App.Static = "public"
	App.Media = "upload"

	// email
	App.Email.From = os.Getenv("EMAIL_FROM")
	App.Email.Username = os.Getenv("EMAIL_USERNAME")
	App.Email.Password = os.Getenv("EMAIL_PASSWORD")
}
