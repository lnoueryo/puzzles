package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
)



type Organization struct {}

// 組織情報の更新
func (*Organization)Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := services.UpdateOrganization(r);if err != nil {
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

