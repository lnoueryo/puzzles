package controllers

import (
	"backend/services"
	"encoding/json"

	"net/http"
)



type ProjectAuthority struct {}

// ユーザーをプロジェクトに追加する
func (*ProjectAuthority)Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := services.CreateProjectAuthority(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }

	w.WriteHeader(http.StatusOK)
	RespondMainUser(w, r)
}

// プロジェクトの権限変更
func (*ProjectAuthority)Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := services.UpdateProjectAuthority(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }

	w.WriteHeader(http.StatusOK)
	RespondMainUser(w, r)
}

// ユーザーをプロジェクトから除外
func (_ *ProjectAuthority)Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := services.DeleteProjectAuthority(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }

	w.WriteHeader(http.StatusOK)
	RespondMainUser(w, r)
}