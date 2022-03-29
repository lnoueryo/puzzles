package controller

import (
	"backend/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"golang.org/x/net/websocket"
)

type Message struct {
	Name    string
	Message string
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
	s, _ := models.CheckSession(r)
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
	s, _ := models.CheckSession(r)
	err := u.GetMainUser(s); if err != nil {
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
