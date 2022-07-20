package models

import (
	"backend/config"
	"backend/modules/mail"
	"time"
)

var email mail.Mail
var origin string
var allowOrigin string
var project string
var UploadToGCS = config.UploadToGCS
var DeleteImage = config.DeleteImage
var errorlog = config.App.ErrorLog

func init() {
	email = config.App.Email
	origin = config.App.Origin
	allowOrigin = config.App.AllowOrigin
	project = config.App.ProjectID
}

func timeToString(t time.Time) string {
	str := t.Format("20060102150405")
	return str
}
