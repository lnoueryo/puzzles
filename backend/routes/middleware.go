package routes

import (
	"backend/modules/session"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// Auth checks if it's valid session
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_cookie");if err != nil {
			errMap := map[string]string{"message": "session is expired"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotModified)
			w.Write(errJson)
			return
		}
		_, err = session.CheckSession(cookie.Value, projectenv)
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

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if allowOrigin == "" {
		allowOrigin = "https://puzzles.jounetsism.biz"
	}
	w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	if r.URL.Path == "/api/data/download" {
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=a.zip")
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	start := time.Now()
	cookie, err := r.Cookie("_cookie");if err != nil {
		infolog.Printf("%s %s %v %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		l.handler.ServeHTTP(w, r)
		return
	}
	fmt.Print(cookie)
	s, err := session.CheckSession(cookie.Value, projectenv)
	if err != nil {
		infolog.Printf("%s %s %v %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	} else {
		infolog.Printf("%v %v %v %v %v %v", r.Method, r.URL.Path, s.Name, s.Email, r.RemoteAddr, time.Since(start))
	}
	l.handler.ServeHTTP(w, r)
}

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
