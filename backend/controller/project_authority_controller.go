package controller

import (
	"backend/models"
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
)



type ProjectAuthority struct {}

func (*ProjectAuthority)Index(w http.ResponseWriter, r *http.Request) {
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

	err = p.GetProject(id, s.UserID); if err != nil {
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

func (*ProjectAuthority)Create(w http.ResponseWriter, r *http.Request) {
	// projectにprojectauthorityを入れてフロントから送る
	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	
    pa, err := models.GetProjectAuthorityJson(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }

	err = pa.Create(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "couldn't create project"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	s, _ := GetSession(r)
	activity := models.Activity{
		UserID: s.UserID,
		ProjectID: pa.ProjectID,
		ContentID: 6,
	}

	err = activity.Create(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}

	pJson, _ := json.Marshal(pa)
	w.WriteHeader(http.StatusCreated)
	w.Write(pJson)
}

func (*ProjectAuthority)Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    pa, err := models.GetProjectAuthorityJson(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	err = pa.Update(); if err != nil {
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	s, _ := GetSession(r)
	activity := models.Activity{
		UserID: s.UserID,
		ProjectID: pa.ProjectID,
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

	projectJson, _ := json.Marshal(pa)
	w.WriteHeader(http.StatusAccepted)
	w.Write(projectJson)
}

func (_ *ProjectAuthority)Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	query := r.URL.Query()
    idSlice, ok := query["id"]; if !ok {
		errorlog.Println(query)
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
    }
	var IDs []int
	for _, ID := range idSlice {
		id, err := strconv.Atoi(ID)
		if err != nil {
			errorlog.Print(err)
			errMap := map[string]string{"message": "bad connection"}
			sessionJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(sessionJson)
			return
		}
		IDs = append(IDs, id)
	}

	pa, err := models.DeleteProjectAuthority(IDs); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	commentJson, _ := json.Marshal(pa)
	w.WriteHeader(http.StatusNoContent)
	w.Write(commentJson)
}