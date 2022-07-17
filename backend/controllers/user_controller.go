package controllers

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
)

type Message struct {
	Name    string
	Message string
}

type MainUser struct {
	User models.User									`json:"user"`
	OrganizationAuthority models.OrganizationAuthority 	`json:"organizationAuthority"`
	Projects []models.Project							`json:"projects"`
}


type User struct{}

// セッションが有効であるか確認
func (_ *User) Show(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	// セッションを確認し、ユーザー情報を作成
	mainUser, err := services.CreateMainUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	uJson, _ := json.Marshal(mainUser)
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}

// ユーザー情報の更新
func (h *User) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	// リクエストで送られてきたユーザー情報のjsonを変換
	err := services.UpdateMainUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJson)
		return
	}

	// 新しい変更ないようでセッションを作成
	mainUser, err := services.CreateMainUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	uJson, _ := json.Marshal(mainUser)
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}
