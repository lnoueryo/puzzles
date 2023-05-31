package controllers

import (
	"backend/config"
	"net/http"
	"backend/modules/session"
	"encoding/json"
)

var user User
var auth Auth
var task Task
var comment Comment
var organization Organization
var organizationAuthority OrganizationAuthority
var project Project
var projectAuthority ProjectAuthority
var csv CSV
var infolog = config.App.InfoLog
var allowOrigin = config.App.AllowOrigin

func Routes() http.Handler{
	mux := http.NewServeMux()

	// static
	staticFiles := http.FileServer(http.Dir(config.App.Static))
	uploadFiles := http.FileServer(http.Dir(config.App.Media))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.Handle("/media/", http.StripPrefix("/media/", uploadFiles))

	// normal
	mux.HandleFunc("/api/login", auth.Login)
	mux.HandleFunc("/register/verification", auth.Confirm)

	// Auth
	mux.Handle("/api/user/update", authRoute(http.HandlerFunc(user.Update)))
	mux.Handle("/api/session", authRoute(http.HandlerFunc(user.Show)))
	mux.Handle("/api/organization/update", authRoute(http.HandlerFunc(organization.Update)))
	mux.Handle("/api/organization-authority/update", authRoute(http.HandlerFunc(organizationAuthority.Update)))
	mux.Handle("/api/project/create", authRoute(http.HandlerFunc(project.Create)))
	mux.Handle("/api/project/update", authRoute(http.HandlerFunc(project.Update)))
	mux.Handle("/api/project-authority/create", authRoute(http.HandlerFunc(projectAuthority.Create)))
	mux.Handle("/api/project-authority/update", authRoute(http.HandlerFunc(projectAuthority.Update)))
	mux.Handle("/api/project-authority/delete", authRoute(http.HandlerFunc(projectAuthority.Delete)))
	mux.Handle("/api/task", authRoute(http.HandlerFunc(task.Index)))
	mux.Handle("/api/task/create", authRoute(http.HandlerFunc(task.Create)))
	mux.Handle("/api/task/update", authRoute(http.HandlerFunc(task.Update)))
	mux.Handle("/api/comment/create", authRoute(http.HandlerFunc(comment.Create)))
	mux.Handle("/api/comment/show", authRoute(http.HandlerFunc(comment.Show)))
	mux.Handle("/api/comment/update", authRoute(http.HandlerFunc(comment.Update)))
	mux.Handle("/api/comment/delete", authRoute(http.HandlerFunc(comment.Delete)))
	mux.Handle("/api/logout", http.HandlerFunc(auth.Logout))
	mux.Handle("/api/invite", http.HandlerFunc(auth.InviteUser))
	
	mux.Handle("/api/data/upload", http.HandlerFunc(csv.Upload))
	mux.Handle("/api/data/download", http.HandlerFunc(csv.Download))

	// JSON
	mux.HandleFunc("/task", task.Index)

	wrappedMux := NewOringCheck(mux)
	return wrappedMux
}

type OringCheck struct {
	handler http.Handler
}

// セッションの確認
func authRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieKey);if err != nil {
			errMap := map[string]string{"message": "session is expired"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotModified)
			w.Write(errJson)
			return
		}
		_, err = session.CheckSession(cookie.Value, projectID)
		if err != nil {
			errMap := map[string]string{"message": "session is expired"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotModified)
			w.Write(errJson)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//　リクエストのたびにログを作成
func (l *OringCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// クロスオリジン用にセット
	w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")

	// CSVのインポートエクスポートの場合
	if r.URL.Path == "/api/data/download" {
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=a.zip")

	// JSONの場合
	} else {
		w.Header().Set("Content-Type", "application/json")
	}

	// preflight用に200でいったん返す
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	l.handler.ServeHTTP(w, r)
}

//NewLogger constructs a new Logger middleware handler
func NewOringCheck(handlerToWrap http.Handler) *OringCheck {
	return &OringCheck{handlerToWrap}
}
