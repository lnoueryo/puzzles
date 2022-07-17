package controllers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"
)



type OrganizationAuthority struct {}

// 組織権限の変更を行う
func (*OrganizationAuthority)Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		errorlog.Print("put method required")
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    oa, err := models.GetOrganizationAuthorityJson(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }

	// 権限の変更
	err = oa.ChangeAuthority(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": err.Error()}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
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

func (*OrganizationAuthority)Delete(w http.ResponseWriter, r *http.Request) {
	// 未実装
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