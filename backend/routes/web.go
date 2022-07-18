package routes

import (
	"backend/config"
	"backend/controllers"
	"net/http"
	// "golang.org/x/net/websocket"
)

var user controllers.User
var auth controllers.Auth
var task controllers.Task
var comment controllers.Comment
var organization controllers.Organization
var organizationAuthority controllers.OrganizationAuthority
var project controllers.Project
var projectAuthority controllers.ProjectAuthority
var csv controllers.CSV
var infolog = config.App.InfoLog
var allowOrigin = config.App.AllowOrigin
var projectID = config.App.ProjectID

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
	mux.Handle("/api/user/update", Auth(http.HandlerFunc(user.Update)))
	mux.Handle("/api/organization/update", Auth(http.HandlerFunc(organization.Update)))
	mux.Handle("/api/organization-authority/update", Auth(http.HandlerFunc(organizationAuthority.Update)))
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
	mux.Handle("/api/session", Auth(http.HandlerFunc(user.Show)))
	mux.Handle("/api/logout", http.HandlerFunc(auth.Logout))
	mux.Handle("/api/invite", http.HandlerFunc(auth.InviteUser))
	
	mux.Handle("/api/data/upload", http.HandlerFunc(csv.Upload))
	mux.Handle("/api/data/download", http.HandlerFunc(csv.Download))

	// JSON
	mux.HandleFunc("/task", task.Index)
	
	// websocket
	// mux.Handle("/chat", websocket.Handler(user.Chat))
	// mux.Handle("/websocket/project/create", websocket.Handler(user.Chat))

	wrappedMux := NewLogger(mux)
	return wrappedMux
}
