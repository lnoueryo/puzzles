package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
)

type Task struct{}

const PAGE_NUM = 0

// タスクの取得
func (*Task) Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	ids, page, err := services.GetIDs(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	if page < 0 {
		errorlog.Print("page num is wrong")
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	w.WriteHeader(http.StatusOK)
	RespondTasks(w, r, ids[0], page)

}

// タスクの作成
func (_ *Task)Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	t, err := services.CreateTask(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
	}

	w.WriteHeader(http.StatusOK)
	RespondTasks(w, r, t.ProjectID, PAGE_NUM)

}

// タスクの更新
func (_ *Task)Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	t, err := services.UpdateTask(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
	}

	w.WriteHeader(http.StatusOK)
	RespondTasks(w, r, t.ProjectID, PAGE_NUM)
}