package config

import (
	"os"
)

// 開発環境の設定
func configureLocalSettings() {
	App.Addr = "0.0.0.0:8080"
	App.AllowOrigin = "http://localhost:3000"
	App.Origin = "http://localhost:8080"
	App.ProjectID = "puzzles-345814"
	// DB接続
	DBSet := Database{
		Name:		os.Getenv("DB_NAME"),
		Host:		os.Getenv("DB_HOST"),
		User:		os.Getenv("DB_USER"),
		Password:	os.Getenv("DB_PASSWORD"),
		Port:		"3306",
		Query:		"parseTime=true",
	}
	ConnectMysql(DBSet)
	// ConnectSqlite3()
}