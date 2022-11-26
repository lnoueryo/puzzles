package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
)

type Comment struct{}

// コメントの作成
func (*Comment)Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	c, err := services.CreateComment(r);if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	w.WriteHeader(http.StatusOK)
	RespondTasks(w, r, c.TaskID, PAGE_NUM)
}

// コメントの取得
func (c *Comment)Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	ids, _, err := services.GetIDs(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	w.WriteHeader(http.StatusCreated)
	RespondTasks(w, r, ids[0], PAGE_NUM)
}

// コメントの更新
func (*Comment)Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	c, err := services.UpdateComment(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
	}

	w.WriteHeader(http.StatusOK)
	RespondTasks(w, r, c.TaskID, PAGE_NUM)
}

// コメントの削除
func (_ *Comment)Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	err := services.DeleteComment(r);if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// RespondTasks(w, r, c.TaskID)
}
