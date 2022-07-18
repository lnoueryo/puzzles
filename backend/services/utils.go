package services

import (
	"backend/config"
	"backend/modules/session"
	"errors"
	"net/http"
	"strconv"
)

var DB = config.DB
var projectID = config.App.ProjectID
var email = config.App.Email
var origin = config.App.Origin
var allowOrigin = config.App.AllowOrigin
var infolog = config.App.InfoLog

// 既存のセッションを取得
func GetSession(r *http.Request) (session.Session, error) {

	var s session.Session
	cookie, err := r.Cookie("_cookie");if err != nil {
		err := errors.New("session is expired")
		return s, err
	}

	s, err = session.CheckSession(cookie.Value, config.App.ProjectID)
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetIDs(r *http.Request) ([]int, error) {
	var ids []int
	query := r.URL.Query()
    idSlice, ok := query["id"]; if !ok {
		message := "couldn't get id"
		err := errors.New(message)
		return ids, err
    }

	for _, ID := range idSlice {
		id, err := strconv.Atoi(ID)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}