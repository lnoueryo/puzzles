package models

import (
	"backend/config"
	"backend/modules/mail"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
var email mail.Mail
var origin string
var allowOrigin string

func init() {
	DB = config.DB
	email = config.App.Email
	origin = config.App.Origin
	allowOrigin = config.App.AllowOrigin
}

func timeToString(t time.Time) string {
	str := t.Format("20060102150405")
	return str
}
