package controllers

import (
	"backend/services"
	"encoding/json"
	"errors"
	"net/http"
)



type CSV struct{}

var (
	csvNameList = []string{
		"activities.csv",
		"activity_contents.csv",
		"authorities.csv",
		"comments.csv",
		"fields.csv",
		"milestones.csv",
		"organization_authorities.csv",
		"organizations.csv",
		"priorities.csv",
		"project_authorities.csv",
		"projects.csv",
		"statuses.csv",
		"tasks.csv",
		"types.csv",
		"users.csv",
	}
)


// DBのデータをCSVに変換しダウンロード
func(*CSV)Download(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
	}

	b, err := services.ExportCSV(r);if err != nil {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
}

// CSVをアップロードし、DBを更新
func (*CSV) Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
	}

	errorArray, err := services.ImportCSV(r);if err != nil {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
	}

	errJson, _ := json.Marshal(errorArray)
	w.WriteHeader(http.StatusOK)
	w.Write(errJson)
}

// 受け取ったJSONを構造体に変換
func GetJson(r *http.Request) (Project, error) {
	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return project, err
	}
	return project, nil
}

