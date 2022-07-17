package services

import (
	"backend/config"
	"backend/modules/session"
	"errors"
	"net/http"
)

var DB = config.DB

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