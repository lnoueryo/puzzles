package controllers

import (
	"backend/config"
	"backend/models"
	"backend/modules/mail"
	"backend/modules/session"
	"backend/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
	"gorm.io/gorm"
)
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]string
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Flash     string
	Warning   string
	Error     string
	JSON      []byte
	Users     []models.User
	Session   session.Session
}

var infolog *log.Logger
var errorlog *log.Logger
var DB *gorm.DB
var ws_array []*websocket.Conn // *websocket.Connを入れる配列
var email mail.Mail
var origin string
var allowOrigin string
var projectID string
var UploadToGCS = config.UploadToGCS
var DeleteImage = config.DeleteImage
var StoreImageToGCS = config.StoreImageToGCS
var StoreBinaryImage = config.StoreBinaryImage


func init() {
	infolog = config.App.InfoLog
	errorlog = config.App.ErrorLog
	DB = config.DB
	email = config.App.Email
	origin = config.App.Origin
	allowOrigin = config.App.AllowOrigin
	projectID = config.App.ProjectID
}

// 既存のセッションを取得
func GetSession(r *http.Request) (session.Session, error) {

	var s session.Session
	cookie, err := r.Cookie("_cookie");if err != nil {
		err := errors.New("session is expired")
		return s, err
	}

	s, err = session.CheckSession(cookie.Value, projectID)
	if err != nil {
		return s, err
	}
	return s, nil
}

// 新たにセッション作成
// func CreateMainUser(r *http.Request) (MainUser, error) {
// 	// BUG(inoueryo) 何かしらデータ更新を行う際に新しくセッションを作成し、古いものは削除されていない可能性あり
// 	var mainUser MainUser
// 	var u models.User
// 	s, err := GetSession(r)
// 	err = u.GetMainUser(s.UserID, s.Organization); if err != nil {
// 		errorlog.Print(err)
// 		return mainUser, err
// 	}

// 	mainUser.Projects = ProjectFilter(u.Organizations[0].Organization.Projects, func(userID int) bool {
// 		return userID == u.ID
//     })
// 	u.Organizations[0].Organization.Projects = nil
// 	mainUser.OrganizationAuthority = u.Organizations[0]
// 	u.Organizations = nil
// 	mainUser.User = u
// 	return mainUser, nil
// }

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

func RespondTasks(w http.ResponseWriter, r *http.Request, id int) {

	t, err := services.GetTask(id); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	tJson, _ := json.Marshal(t)
	w.Write(tJson)
}