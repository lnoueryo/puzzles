package controller

import (
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/image"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	Name    string
	Message string
}

type MainUser struct {
	User models.User 							`json:"user"`
	Organization models.OrganizationAuthority 	`json:"organization"`
	Projects []models.Project 					`json:"projects"`
}


type Home struct{}

func (_ *Home) Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
	}
	var org models.Organization
	s, _ := GetSession(r)
	err := org.GetOrganization(s.Organization); if err != nil {
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
	}
	orgJson, _ := json.Marshal(org)
	w.WriteHeader(http.StatusOK)
	w.Write(orgJson)
}

func (_ *Home) Show(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	var u models.User
	s, err := GetSession(r)
	if err != nil {
		errMap := map[string]string{"message": "session is expired"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotModified)
		w.Write(errJson)
		return
	}
	infolog.Print(s)
	err = u.GetMainUser(s.UserID, s.Organization); if err != nil {
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	var mainUser MainUser
	mainUser.Projects = ProjectFilter(u.Organizations[0].Organization.Projects, func(userID int) bool {
        return userID == u.ID
    })
	u.Organizations[0].Organization.Projects = nil
	mainUser.Organization = u.Organizations[0]
	u.Organizations = nil
	mainUser.User = u
	uJson, _ := json.Marshal(mainUser)
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}

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

func (h *Home) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}
	u, err := models.GetUserJson(r); if err != nil {
		errMap := map[string]string{"message": "bad connection"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJson)
		return
	}
	if u.ImageData != "" && u.Image != "" {
		fileName, err := StoreImage("users", u.ImageData); if err != nil {
			errorlog.Print(err);
			errMap := map[string]string{"message": "couldn't save the image"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotFound)
			w.Write(errJson)
			return
		}
		err = os.Remove("./upload/users/" + u.Image); if err != nil {
			errorlog.Print(err);
			errMap := map[string]string{"message": "couldn't save the image"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotFound)
			w.Write(errJson)
			return
		}
		u.Image = fileName
	}
	if u.Image == "" {
		u.Image, _ = crypto.MakeRandomStr(20)
		buf, err := image.CreateImage(u.Name, u.Image); if err != nil {
			errorlog.Print(err);
			errMap := map[string]string{"message": "couldn't save the image"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotFound)
			w.Write(errJson)
			return
		}
		path := "users/" + u.Image
		if credentialsPath != "" {
			err = StoreImageToGCS(buf.Bytes(), path)
		} else {
			err = StoreBinaryImage(buf.Bytes(), path)
		}
		if err != nil {
			errorlog.Print(err);
			errMap := map[string]string{"message": "couldn't save the image"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotFound)
			w.Write(errJson)
			return
		}
	}
	err = u.Update(); if err != nil {
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	uJson, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}

func (h *Home) Chat(ws *websocket.Conn) {
	ws_array = append(ws_array, ws)
	data_receive(ws)
}


func data_receive(ws *websocket.Conn) {
	for {
		var message models.Project
		type Count struct{ID int}
		if err := websocket.JSON.Receive(ws, &message); err != nil {
			log.Println("Receive error:", err)
			break
		} else {
			for _, con := range ws_array {
				con := con
				c := make(chan string)
				go func() {
					for {
						msg, ok := <-c
						if ok {
							err = websocket.JSON.Send(con, msg)
						}
					}
				}()
				c <- "start"
				time.Sleep(time.Second)
				c <- "half"
				time.Sleep(time.Second)
				c <- "{\"id\": hello}"
			}
		}
	}
}
