package config

import (
	"os"
)

// 本番環境の設定
func configureProdSettings() {
	App.Addr = ":8500"
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