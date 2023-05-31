package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
)



type Project struct {}

// プロジェクトの作成
func (*Project)Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := services.CreateProject(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "couldn't create project"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}

	w.WriteHeader(http.StatusCreated)
	RespondMainUser(w, r)
}

// プロジェクトの更新
func (*Project)Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    err := services.UpdateProject(r);if err != nil {
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