package config

import (
	"os"
)

// 開発環境の設定
func configureLocalSettings() {
	App.Addr = "0.0.0.0:8500"

	// DB接続
	DBSet := Database{
		Name: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port: os.Getenv("DB_PORT"),
		Query: os.Getenv("DB_QUERY"),
	}
	ConnectMysql(DBSet)
	// ConnectSqlite3()
}