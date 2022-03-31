package controller

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

type Comment struct{}

func (c *Comment)Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    comment, err := models.NewComment(r);if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	err = comment.Create(); if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	task := models.Task{
		ID: comment.TaskID,
	}
	err = task.GetTask(); if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}

	s, _ := models.CheckSession(r)
	activity := models.Activity{
		UserID: s.UserID,
		ProjectID: task.ProjectID,
		ContentID: 3,
	}

	err = activity.Create(); if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}

	taskJson, _ := json.Marshal(task)
	w.WriteHeader(http.StatusCreated)
	w.Write(taskJson)
}

func (c *Comment)Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    comment, err := models.NewComment(r);if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	infolog.Println(comment.Replies)
	err = comment.Update(); if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	task := models.Task{
		ID: comment.TaskID,
	}
	err = task.GetTask(); if err != nil {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	taskJson, _ := json.Marshal(task)
	w.WriteHeader(http.StatusCreated)
	w.Write(taskJson)
}
