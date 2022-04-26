package session

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Session struct {
	ID           string    `json:"id"`
	UserID       int       `json:"user_id"`
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Sex          string    `json:"sex"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	Image        string    `json:"image"`
	Description  string    `json:"description"`
	Organization string    `json:"organization"`
	Authority    string    `json:"authority"`
	CSRFToken    string    `json:"csrf_token"`
	CreatedAt    time.Time `json:"-"`
}

func timeToString(t time.Time) string {
	str := t.Format("20060102150405")
	return str
}

func (s *Session)GenerateSessionID() {
	sessionId := string(s.UserID) + timeToString(s.CreatedAt) + timeToString(time.Now())
	hashedByteSessionId := sha256.Sum256([]byte(sessionId))
	hashedSessionId := fmt.Sprintf("%x", (hashedByteSessionId))
	s.ID = hashedSessionId
}