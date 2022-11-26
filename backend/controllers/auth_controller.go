package controllers

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
)

type Auth struct{}

// ログイン処理
func (au *Auth) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
	}

	r, err := services.Login(w, r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errJson)
		return
	}

	w.WriteHeader(http.StatusOK)
	RespondMainUser(w, r)
}

// ログアウト処理
func (au *Auth) Logout(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	err := services.Logout(&w, r);if err != nil {
		errorlog.Print(err)
		errMessage := "session is expired"
		errorlog.Print(errMessage)
		errMap := map[string]string{"message": errMessage}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errJson)
		return
	}

	message := "logout is successful"
	successMap := map[string]string{"message": message}
	successJson, _ := json.Marshal(successMap)
	w.WriteHeader(http.StatusOK)
	w.Write(successJson)
}

// ユーザー追加
func (*Auth) InviteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	// ユーザーを組織に招待
	err := services.InviteUser(r); if err != nil {
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

// ユーザーが本人であることの確認
func (au *Auth) Confirm(w http.ResponseWriter, r *http.Request) {

	// verificationコードが正しいか確認
	var oa models.OrganizationAuthority
	url, err := services.ValidateVerification(r, &oa);if err != nil {
		errorlog.Print(err)
		http.Redirect(w, r, url, http.StatusNotFound)
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
