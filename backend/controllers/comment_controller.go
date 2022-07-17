package controllers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type Comment struct{}

// コメントの作成
func (c *Comment)Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    comment, err := models.NewComment(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	err = comment.Create(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	// infolog.Print(comment)
	// comments, err := models.GetComments(comment.TaskID); if err != nil {
	// 	errMap := map[string]string{"message": "not found"}
	// 	errJson, _ := json.Marshal(errMap)
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write(errJson)
	// 	return
	// }

	// s, _ := GetSession(r)
	// activity := models.Activity{
	// 	UserID: s.UserID,
	// 	ProjectID: task.ProjectID,
	// 	ContentID: 3,
	// }

	// err = activity.Create(); if err != nil {
	// 	errMap := map[string]string{"message": "not found"}
	// 	errJson, _ := json.Marshal(errMap)
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write(errJson)
	// }

	commentJson, _ := json.Marshal(comment)
	w.WriteHeader(http.StatusCreated)
	w.Write(commentJson)
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
	comments, _ := models.GetComments(id)
	commentJson, _ := json.Marshal(comments)
	w.WriteHeader(http.StatusCreated)
	w.Write(commentJson)
}

// コメントの更新
func (c *Comment)Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

    comment, err := models.NewComment(r);if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
        return
    }
	err = comment.Update(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
	}
	taskJson, _ := json.Marshal(comment)
	w.WriteHeader(http.StatusCreated)
	w.Write(taskJson)
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

	query := r.URL.Query()
    idSlice, ok := query["id[]"]; if !ok {
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

	c, err := models.DeleteComment(IDs); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	commentJson, _ := json.Marshal(c)
	w.WriteHeader(http.StatusNoContent)
	w.Write(commentJson)
}
