package controller

import (
	"backend/models"
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
)



type Organization struct {}

func (*Organization)Create(w http.ResponseWriter, r *http.Request) {
	// projectにprojectauthorityを入れてフロントから送る
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

	pa, err := models.GetProjectAuthority(p.ID, s.UserID); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "couldn't create project"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	pJson, _ := json.Marshal(pa)
	w.WriteHeader(http.StatusCreated)
	w.Write(pJson)
}

func (*Organization)Edit(w http.ResponseWriter, r *http.Request) {
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

// 組織情報の更新
func (*Organization)Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    o, err := models.GetOrganizationJson(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	err = o.Update(); if err != nil {
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	// s, _ := GetSession(r)
	// activity := models.Activity{
	// 	UserID: s.UserID,
	// 	ProjectID: o.ID,
	// 	ContentID: 6,
	// }

	// err = activity.Create(); if err != nil {
	// 	errorlog.Print(err)
	// 	errMap := map[string]string{"message": "not found"}
	// 	errJson, _ := json.Marshal(errMap)
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write(errJson)
	// 	return
	// }

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

