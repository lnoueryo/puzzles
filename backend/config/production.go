package config

import (
	"os"
)

// 本番環境の設定
func configureProdSettings() {
	App.Addr = ":8080"
	App.AllowOrigin = os.Getenv("ALLOW_ORIGIN")
	App.Origin = "puzzles-api.jounetsism.biz"
	App.Project = "puzzles-345814"
	DBSet := Database{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     "3306",
		Query:    "parseTime=true",
	}
	_, err := ConnectMysql(DBSet);if err != nil {
		ConnectSqlite3()
	}
}