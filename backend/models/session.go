package models

import (
	"backend/modules/crypto"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

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

// Checks if the user is logged in and has a session, if not err is not nil
func CheckSession(r *http.Request) (Session, bool) {
	cookie, err := r.Cookie("_cookie")
	s := Session{}
	if err == nil {
		filepath := fmt.Sprintf("./session/%v.txt", cookie.Value)
		err = s.readSession(filepath)
		return s, IsSession(filepath)
	}
	return s, false
}

func GetSession(r *http.Request) Session {
	cookie, _ := r.Cookie("_cookie")
	s := Session{}
	filepath := fmt.Sprintf("./session/%v.txt", cookie.Value)
	s.readSession(filepath)
	return s
}

func DeliverSession(r *http.Request) Session {
	s := GetSession(r)
	err := s.GenerateCSRFToken(r)
	if err != nil {
		log.Print(err)
	}
	return s
}

func (s *Session) GenerateCSRFToken(r *http.Request) error {
	filepath := fmt.Sprintf("./session/%v.txt", s.ID)
	s.CSRFToken, _ = crypto.MakeRandomStr(32)
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(f)
	defer f.Close()

	if err := enc.Encode(&s); err != nil {
		return err
	}
	return nil
}

func (s *Session) CheckCSRFToken(r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		log.Print(err, "Cannot find user")
	}
	if s.CSRFToken != r.Form.Get("csrf_token") {
		return false
	}
	return true
}

// Checks if the user is logged in and has a session, if not err is not nil
func (s *Session) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("./session/%v.txt", cookie.Value)
	isSession := IsSession(filename)
	if isSession {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	DeleteCookie(w, r)
	return nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return err
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	return nil
}

func IsSession(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	} else {
		return os.IsExist(err)
	}
}

func (s *Session) readSession(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(f)
	defer f.Close()
	if err := dec.Decode(&s); err != nil {
		return err
	}
	return nil
}

func FormCookie(w http.ResponseWriter, r *http.Request) {
	jsonByte, _ := json.Marshal(&r)
	j64Value := base64.StdEncoding.EncodeToString(jsonByte)
	expiryTime := time.Now().Add(time.Second * 15)
	cookie := http.Cookie{
		Name:     "_form",
		Value:    j64Value,
		HttpOnly: true,
		Secure:   true,
		Path:     "/users",
		Expires:  expiryTime,
	}
	http.SetCookie(w, &cookie)
}

func CheckFormCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_form")
	if err == nil {
		var form http.Request
		ds, _ := base64.StdEncoding.DecodeString(cookie.Value)
		json.Unmarshal(ds, &form)
		for key, values := range form.Form { // range over map
			for _, value := range values { // range over []string
				fmt.Println(key, value)
			}
		}
	}
}
