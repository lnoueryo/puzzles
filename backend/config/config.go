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
	CookieKey		string
}

var App AppConfig
var infolog *log.Logger
var errorlog *log.Logger
var DB_NAME string
var DB_HOST string
var DB_USER string
var DB_PASSWORD string
var SERVER_PORT string
var BUCKET_NAME string


// SESSION=DATASTORE
// APP_ORIGIN=puzzles-api.jounetsism.biz
// EMAIL_FROM=popo62520908@gmail.com
// EMAIL_USERNAME=popo62520908@gmail.com
// EMAIL_PASSWORD=xprcmxlrrfbnodux
// envによる環境の決定
func init() {
    godotenv.Load(".env")

	// log
	infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App.InfoLog = infolog
	App.ErrorLog = errorlog

	// file path
	App.Static = "public"
	App.Media = "upload"

	App.CookieKey = "_cookie"

	// email
	App.Email.From = os.Getenv("EMAIL_FROM")
	infolog.Print("EMAIL_FROM: ", App.Email.From)
	App.Email.Username = os.Getenv("EMAIL_USERNAME")
	infolog.Print("EMAIL_USERNAME: ", App.Email.Username)
	App.Email.Password = os.Getenv("EMAIL_PASSWORD")
	infolog.Print("EMAIL_PASSWORD: ", App.Email.Password)
	DB_NAME = os.Getenv("DB_NAME");if DB_NAME == "" {
		DB_NAME = "puzzle"
	}
	infolog.Print("DB_NAME: ", DB_NAME)
	DB_HOST = os.Getenv("DB_HOST");if DB_HOST == "" {
		DB_HOST = "localhost"
	}
	infolog.Print("DB_HOST: ", DB_HOST)
	DB_USER = os.Getenv("DB_USER");if DB_USER == "" {
		DB_USER = "puzzles"
	}
	infolog.Print("DB_USER: ", DB_USER)
	DB_PASSWORD = os.Getenv("DB_PASSWORD");if DB_PASSWORD == "" {
		DB_PASSWORD = "password"
	}
	infolog.Print("DB_PASSWORD: ", DB_PASSWORD)
	App.Addr = os.Getenv("SERVER_PORT");if App.Addr == "" {
		App.Addr = "8080"
	}
	infolog.Print("SERVER_PORT: ", App.Addr)
	App.AllowOrigin = os.Getenv("ALLOW_ORIGIN");if App.AllowOrigin == "" {
		App.AllowOrigin = "*"
	}
	infolog.Print("ALLOW_ORIGIN: ", App.AllowOrigin)
	App.Host = os.Getenv("HOST");if App.Host == "" {
		App.Host = "http://localhost:8080"
	}
	infolog.Print("HOST: ", App.Host)
	App.ProjectID = "private-361516"
	BUCKET_NAME = "puzzles-media"
	// DB接続
	DBSet := Database{
		Name:		DB_NAME,
		Host:		DB_HOST,
		User:		DB_USER,
		Password:	DB_PASSWORD,
		Port:		"3306",
		Query:		"parseTime=true",
	}
	ConnectMysql(DBSet)
}

// 本番、開発環境共通のセッティング
func commonSettings() {
	// mode := os.Getenv("MODE");if mode == "" {
	// 	mode = "develop"
	// }

	// if mode == "production" {
	// 	App.Addr = ":8080"
	// 	App.Origin = "puzzles-api.jounetsism.biz"
	// 	App.ProjectID = "puzzles-345814"
	// }

}
