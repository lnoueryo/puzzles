package controller

import (
	"backend/models"
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
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

    p, err := models.NewProject(r)
    if err != nil {
		errorlog.Print(err)
        return
    }

	if p.ImageData != "" {
		fileName, err := StoreImage("projects", p.ImageData); if err != nil {
			errorlog.Print(err);
			errMap := map[string]string{"message": "couldn't save the image"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotFound)
			w.Write(errJson)
			return
		}
		p.Image = fileName
	}
	err = p.Create(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "couldn't create project"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	s, _ := GetSession(r)
	activity := models.Activity{
		UserID: s.UserID,
		ProjectID: p.ID,
		ContentID: 6,
	}

	err = activity.Create(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}

	mainUser, err := CreateMainUser(r); if err != nil {
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

func (*Project)Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	query := r.URL.Query()
    idSlice, ok := query["id"]; if !ok {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
    }
	id, err := strconv.Atoi(idSlice[0])
	if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	var p models.Project
	s, _ := GetSession(r)

	err = p.GetEditProject(id, s.UserID); if err != nil {
		errorlog.Print(err)
		message := "bad connection"
		errorlog.Print(message)
		errMap := map[string]string{"message": message}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	if p.OrganizationID != s.Organization {
		message := "not authorized"
		errorlog.Print(message)
		errMap := map[string]string{"message": message}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusForbidden)
		w.Write(sessionJson)
		return
	}
	pJson, _ := json.Marshal(p)
	w.WriteHeader(http.StatusOK)
	w.Write(pJson)
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

    pur, err := models.GetProjectUpdateRequestJson(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	err = pur.BulkUpdateProject(); if err != nil {
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	s, _ := GetSession(r)
	activity := models.Activity{
		UserID: s.UserID,
		ProjectID: pur.Project.ID,
		ContentID: 6,
	}

	err = activity.Create(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	projectJson, _ := json.Marshal(pur.Project)
	w.WriteHeader(http.StatusAccepted)
	w.Write(projectJson)
}

