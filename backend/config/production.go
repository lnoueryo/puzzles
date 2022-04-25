package config

import (
	"os"
)

func configureProdSettings() {
	App.UseCache = true
	tc, err := CreateTemplateCache()
	if err != nil {
		errorlog.Fatal(err)
	}
	App.TemplateCache = tc
	App.Addr = ":8080"
	if os.Getenv("DB") == "CLOUDSQL" {
		DBSet := Database{
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Query:    os.Getenv("DB_QUERY"),
		}
		ConnectMysql(DBSet)
		return
	}
	ConnectSqlite3()
}