package commands

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"
	ses "backend/modules/session"
	"backend/models"
	"cloud.google.com/go/datastore"
)
type Entity struct {
	Value string
}
type Session struct {
	ID				string		`json:"id"`
	UserID			int			`json:"user_id"`
	Name			string		`json:"name"`
	Age				int			`json:"age"`
	Sex				string		`json:"sex"`
	Email			string		`json:"email"`
	Address			string		`json:"address"`
	Image			string		`json:"image"`
	Description		string		`json:"description"`
	Organization	string		`json:"organization"`
	Authority		string		`json:"authority"`
	CSRFToken		string		`json:"csrf_token"`
	CreatedAt		time.Time	`json:"-"`
}

func Practice() {
	// temp()
	var u models.User
	u.ID = 1
	var se ses.Session
	se.UserID = u.ID
	// se.CreateSession()
	return
	sessionId := string(1) + timeToString(time.Now())
	hashedByteSessionId := sha256.Sum256([]byte(sessionId))
	hashedSessionId := fmt.Sprintf("%x", (hashedByteSessionId))
	// session用ファイル作成
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, "puzzles-345814")
	if err != nil {
		// Handle error.
	}
	defer dsClient.Close()

	k := datastore.NameKey("Session", hashedSessionId, nil)
	s := new(Session)

	s.ID = hashedSessionId
	s.UserID = 1
	s.Name = "井上領"
	s.Age = 31
	s.Sex = "男性"
	s.Address = "東京都世田谷区松原"
	s.Email = "popo62520908@gmail.com"
	s.Description = "ハロー"
	if _, err := dsClient.Put(ctx, k, s); err != nil {
		fmt.Print(err,1)
		return
	}
	// 削除
	// key := datastore.NameKey("Session", "eaa86b02826364894589ec734055407ea3538399bd41c5cc50bfa859945dfc01", nil)
	// dsClient.Delete(ctx, key)
	// 取得
	// key := datastore.NameKey("Session", "e62fffac298be362f13385a03db9005823c13081064c2271040530c05d909fab", nil)
	// session := &Session{}
	// if err := dsClient.Get(ctx, key, session); err != nil {
	// 	// TODO: Handle error.
	// }
}

func timeToString(t time.Time) string {
	str := t.Format("20060102150405")
	return str
}

func temp() {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, "my-project")
	if err != nil {
		// Handle error.
	}
	defer dsClient.Close()

	k := datastore.NameKey("Entity", "stringID", nil)
	e := new(Entity)
	if err := dsClient.Get(ctx, k, e); err != nil {
		// Handle error.
	}

	old := e.Value
	e.Value = "Hello World!"

	if _, err := dsClient.Put(ctx, k, e); err != nil {
		// Handle error.
	}

	fmt.Printf("Updated value from %q to %q\n", old, e.Value)
}


// import(
// 	"log"
// 	"time"
// 	"golang.org/x/net/websocket"
// )

// func (h *Home) Chat(ws *websocket.Conn) {
// 	ws_array = append(ws_array, ws)
// 	data_receive(ws)
// }


// func data_receive(ws *websocket.Conn) {
// 	for {
// 		var message models.Project
// 		type Count struct{ID int}
// 		if err := websocket.JSON.Receive(ws, &message); err != nil {
// 			log.Println("Receive error:", err)
// 			break
// 		} else {
// 			for _, con := range ws_array {
// 				con := con
// 				c := make(chan string)
// 				go func() {
// 					for {
// 						msg, ok := <-c
// 						if ok {
// 							err = websocket.JSON.Send(con, msg)
// 						}
// 					}
// 				}()
// 				c <- "start"
// 				time.Sleep(time.Second)
// 				c <- "half"
// 				time.Sleep(time.Second)
// 				c <- "{\"id\": hello}"
// 			}
// 		}
// 	}
// }
