package routes

import (
	"backend/config"
	"backend/controller"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var home controller.Home
var auth controller.Auth
var task controller.Task
var comment controller.Comment
var project controller.Project
var infolog *log.Logger

func init() {
	infolog = config.App.InfoLog
}

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

	// Auth
	mux.Handle("/api/home", Auth(http.HandlerFunc(home.Index)))
	mux.Handle("/api/project", Auth(http.HandlerFunc(project.Index)))
	mux.Handle("/api/project/edit", Auth(http.HandlerFunc(project.Edit)))
	mux.Handle("/api/project/create", Auth(http.HandlerFunc(project.Create)))
	mux.Handle("/api/project/update", Auth(http.HandlerFunc(project.Update)))
	mux.Handle("/api/task", Auth(http.HandlerFunc(task.Index)))
	mux.Handle("/api/task/create", Auth(http.HandlerFunc(task.Create)))
	mux.Handle("/api/task/update", Auth(http.HandlerFunc(task.Update)))
	mux.Handle("/api/comment/create", Auth(http.HandlerFunc(comment.Create)))
	mux.Handle("/api/comment/update", Auth(http.HandlerFunc(comment.Update)))
	mux.Handle("/api/session", Auth(http.HandlerFunc(home.Show)))
	mux.Handle("/api/logout", http.HandlerFunc(auth.Logout))
	
	// JSON
	mux.HandleFunc("/task", task.Index)
	
	// websocket
	mux.Handle("/chat", websocket.Handler(home.Chat))
	mux.Handle("/websocket/project/create", websocket.Handler(home.Chat))

	wrappedMux := NewLogger(mux)
	return wrappedMux
}
