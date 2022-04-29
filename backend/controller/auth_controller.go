package controller

import (
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/mail"
	"backend/modules/oauth"
	"backend/modules/session"
	"encoding/json"
	"net/http"
	"time"
)

type Auth struct{}

func (au *Auth) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
	}

	u, err := models.TryToLogin(w, r)
	if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errJson)
		return
	}
	s, err := u.CreateSession(w)
	if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errJson)
		return
	}
	sessionJson, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(sessionJson)
}

func (au *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	s, err := GetSession(r)
	if err != nil {
		errorlog.Print(err)
		errMessage := "session is expired"
		errorlog.Print(errMessage)
		errMap := map[string]string{"message": errMessage}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errJson)
		return
	}
	infolog.Println("delete")
	err = session.DeleteSession(s.ID, project)
	if err != nil {
		errorlog.Print(err)
	}
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		errorlog.Print(err)
		errMessage := "session is expired"
		errMap := map[string]string{"message": errMessage}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNoContent)
		w.Write(errJson)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	message := "logout is successful"
	successMap := map[string]string{"message": message}
	successJson, _ := json.Marshal(successMap)
	w.WriteHeader(http.StatusOK)
	w.Write(successJson)
}

func (au *Auth) GitHubLogin(w http.ResponseWriter, r *http.Request) {
	userInfo, err := oauth.GithubOAuth(w, r)
	if err != nil {
		errorlog.Print(err)
	}
	// databaseの処理Createを記載する↓↓
	http.Redirect(w, r, "/?access_token="+userInfo.AccessToken, http.StatusFound)
}

func (*Auth) InviteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := models.InviteUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	errMap := map[string]string{"message": "success"}
	errJson, _ := json.Marshal(errMap)
	w.WriteHeader(http.StatusOK)
	w.Write(errJson)
}

func (au *Auth) Confirm(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
	code, ok := q["code"];if !ok {
		url := allowOrigin + "/expiry"
		http.Redirect(w, r, url, 301)
		return
	}
	verification := code[0]
	if verification == "" {
		url := allowOrigin + "/expiry"
		http.Redirect(w, r, url, 301)
		return
	}
	var oa models.OrganizationAuthority
	err := oa.Find(verification);if err != nil {
		url := allowOrigin + "/expiry"
		http.Redirect(w, r, url, 301)
		return
	}

	oneDay := time.Now().Add(-time.Hour * 24)
	if !oneDay.Before(oa.UpdatedAt) {
		url := allowOrigin + "/expiry"
		http.Redirect(w, r, url, 301)
		return
	}

	err = oa.ChangeActive(); if err != nil {
		errorlog.Print(err)
		url := allowOrigin + "/bad-connection"
		http.Redirect(w, r, url, 404)
		return
	}

	url := allowOrigin + "/success"
	m := email
	m.Sub = "承認されました"
	m.To = oa.User.Email
	m.Message = "組織ID: " + oa.OrganizationID + "\nメールアドレス: " + oa.User.Email
	if oa.User.Name == "" {
		password, _ := crypto.MakeRandomStr(20)
		oa.User.Password = crypto.Encrypt(password)
		err = oa.User.Update(); if err !=nil {
			errorlog.Print(err)
			url := allowOrigin + "/bad-connection"
			http.Redirect(w, r, url, 404)
			return
		}
		m.Message += "\n初回パスワード: " + password
	}
	err = mail.SendEmail(m); if err !=nil {
		errorlog.Print(err)
		errorlog.Print(m)
		return
	}
	http.Redirect(w, r, url, 301)
}
