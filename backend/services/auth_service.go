package services

import (
	"backend/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type LoginUser struct {
	Organization	string    `json:"organization"`
	Email			string    `json:"email"`
	Password		string    `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) error {
	// バリデーション
	u, err := tryToLogin(w, r);if err != nil {
		return err
	}

	// ユーザーが発見できた場合セッションを作成
	_, err = u.CreateSession(w);if err != nil {
		return err
	}

	return nil
}

// ログインのバリデーション
func tryToLogin(w http.ResponseWriter, r *http.Request) (*models.User, error) {

	var u *models.User
	l, err := GetLoginJson(r);if err != nil {
		return u, err
	}

	// ログインフォーム空白確認
	err = l.CheckLoginFormBlank();if err != nil {
		return u, err
	}

	// メールアドレスのフォームチェック
	err = models.CheckEmailFormat(l.Email);if err != nil {
		return u, err
	}

	// 全ての項目を踏まえログイン情報が正しい確認
	err = u.FindLoginUser(DB, l.Email, l.Password, l.Organization);if err != nil {
		return u, err
	}

	return u, nil
}

func (l *LoginUser) CheckLoginFormBlank() error {

	if l.Organization == "" {
		message := "organization is blank"
		err := errors.New(message)
		return err
	}

	if l.Email == "" {
		message := "email address is blank"
		err := errors.New(message)
		return err
	}

	if l.Password == "" {
		message := "password is blank"
		err := errors.New(message)
		return err
	}
	return nil
}

func GetLoginJson(r *http.Request) (LoginUser, error) {
	var login LoginUser
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return login, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &login)
		if err != nil {
			return login, err
		}
		return login, nil
	}

	message := "request body is empty"
	err = errors.New(message)
	return login, err
}