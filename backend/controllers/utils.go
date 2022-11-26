package controllers

import (
	"backend/config"
	"backend/models"
	"backend/modules/session"
	"backend/services"
	"encoding/json"
	"errors"
	"net/http"
)

var errorlog = config.App.ErrorLog
var projectID = config.App.ProjectID
var cookieKey = config.App.CookieKey

// 既存のセッションを取得
func GetSession(r *http.Request) (session.Session, error) {

	var s session.Session
	cookie, err := r.Cookie(cookieKey);if err != nil {
		err := errors.New("session is expired")
		return s, err
	}

	s, err = session.CheckSession(cookie.Value, projectID)
	if err != nil {
		return s, err
	}
	return s, nil
}

// ユーザーが参加しているプロジェクトをフィルター
func ProjectFilter(projects []models.Project, f func(int) bool) []models.Project {
	var userProjects []models.Project
	for _, project := range projects {
		isUser := false
		for _, user := range project.AuthorityUsers {
			if f(user.UserID) {
				isUser = true
				break
			}
		}
		if isUser {
			userProjects = append(userProjects, project)
		}
	}
	return userProjects
}

func RespondMainUser(w http.ResponseWriter, r *http.Request) {
	// セッションを確認し、ユーザー情報を作成
	mainUser, err := services.CreateMainUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	uJson, _ := json.Marshal(mainUser)
	w.Write(uJson)
}

func RespondTasks(w http.ResponseWriter, r *http.Request, id int, pageNum int) {

	var ses session.Session
	s := r.Context().Value(ses).(session.Session)
	isAllowed := services.FindProjectAuthority(id, s.UserID);if !isAllowed {
		errorlog.Printf("try to get tasks for different project: %v %v %v %v %v", r.Method, r.URL.Path, s.Name, s.Email, r.RemoteAddr)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	t, err := services.GetTask(id, pageNum); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	// tJson, _ := json.Marshal(t)

	w.Write(t)
}

func RespondComments(w http.ResponseWriter, r *http.Request, id int) {

	c, err := services.GetComment(id); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	cJson, _ := json.Marshal(c)
	w.Write(cJson)
}