package routes

import (
	"backend/config"
	"backend/controller"
	"net/http"
	"golang.org/x/net/websocket"
)

var home controller.Home
var auth controller.Auth
var task controller.Task
var comment controller.Comment
var project controller.Project
var projectAuthority controller.ProjectAuthority
var data controller.Data
var infolog = config.App.InfoLog
var allowOrigin = config.App.AllowOrigin
var projectenv = config.App.Project

func Routes() http.Handler{
	mux := http.NewServeMux()

	// static
	staticFiles := http.FileServer(http.Dir(config.App.Static))
	uploadFiles := http.FileServer(http.Dir(config.App.Media))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.Handle("/media/", http.StripPrefix("/media/", uploadFiles))

	// normal
	mux.HandleFunc("/api/login", auth.Login)
	// mux.HandleFunc("/sign-up", auth.Register)
	mux.HandleFunc("/oauth/callback", auth.GitHubLogin)
	mux.HandleFunc("/register/verification", auth.Confirm)

	// Auth
	mux.Handle("/api/home", Auth(http.HandlerFunc(home.Index)))
	mux.Handle("/api/user/update", Auth(http.HandlerFunc(home.Update)))
	mux.Handle("/api/project", Auth(http.HandlerFunc(project.Index)))
	mux.Handle("/api/project/edit", Auth(http.HandlerFunc(project.Edit)))
	mux.Handle("/api/project/create", Auth(http.HandlerFunc(project.Create)))
	mux.Handle("/api/project/update", Auth(http.HandlerFunc(project.Update)))
	mux.Handle("/api/project-authority/create", Auth(http.HandlerFunc(projectAuthority.Create)))
	mux.Handle("/api/project-authority/update", Auth(http.HandlerFunc(projectAuthority.Update)))
	mux.Handle("/api/project-authority/delete", Auth(http.HandlerFunc(projectAuthority.Delete)))
	mux.Handle("/api/task", Auth(http.HandlerFunc(task.Index)))
	mux.Handle("/api/task/create", Auth(http.HandlerFunc(task.Create)))
	mux.Handle("/api/task/update", Auth(http.HandlerFunc(task.Update)))
	mux.Handle("/api/comment/create", Auth(http.HandlerFunc(comment.Create)))
	mux.Handle("/api/comment/show", Auth(http.HandlerFunc(comment.Show)))
	mux.Handle("/api/comment/update", Auth(http.HandlerFunc(comment.Update)))
	mux.Handle("/api/comment/delete", Auth(http.HandlerFunc(comment.Delete)))
	mux.Handle("/api/session", Auth(http.HandlerFunc(home.Show)))
	mux.Handle("/api/logout", http.HandlerFunc(auth.Logout))
	mux.Handle("/api/invite", http.HandlerFunc(auth.InviteUser))
	
	mux.Handle("/api/data/upload", http.HandlerFunc(data.Upload))
	mux.Handle("/api/data/download", http.HandlerFunc(data.Download))

	// JSON
	mux.HandleFunc("/task", task.Index)
	
	// websocket
	mux.Handle("/chat", websocket.Handler(home.Chat))
	mux.Handle("/websocket/project/create", websocket.Handler(home.Chat))

	wrappedMux := NewLogger(mux)
	return wrappedMux
}
