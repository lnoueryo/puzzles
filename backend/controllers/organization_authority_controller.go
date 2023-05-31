package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
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

	err := services.UpdateOrganizationAuthority(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	w.WriteHeader(http.StatusOK)
	RespondMainUser(w, r)
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

	err := services.DeleteOrganizationAuthority(r); if err != nil {
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