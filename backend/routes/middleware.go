package routes

import (
	"backend/models"
	"encoding/json"
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
		_, isSession := models.CheckSession(r)
		if !isSession {
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	start := time.Now()
	session, isSession := models.CheckSession(r)
	if !isSession {
		infolog.Printf("%s %s %v %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	} else {
		infolog.Printf("%v %v %v %v %v %v", r.Method, r.URL.Path, session.Name, session.Email, r.RemoteAddr, time.Since(start))
	}
	l.handler.ServeHTTP(w, r)
}

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
