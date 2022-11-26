package routes

import (
	"backend/modules/session"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

//Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// type httpContextStruct struct {
//     user  string
//     wazaa string
// }

// var httpContext httpContextStruct
// var ses session.Session
// セッションの確認
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieKey);if err != nil {
			infolog.Printf("no cookie %s %s %v", r.Method, r.URL.Path, r.RemoteAddr)
			errMap := map[string]string{"message": "session is expired"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotModified)
			w.Write(errJson)
			return
		}
		s, err := session.CheckSession(cookie.Value, projectID);if err != nil {
			errMap := map[string]string{"message": "session is expired"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotModified)
			w.Write(errJson)
			return
		}
		infolog.Printf("%v %v %v %v %v", r.Method, r.URL.Path, s.Name, s.Email, r.RemoteAddr)
        ses := session.Session{}
        r = r.WithContext(context.WithValue(
            r.Context(),
            ses,
            s,
        ))
		next.ServeHTTP(w, r)
	})
}

//　リクエストのたびにログを作成
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// クロスオリジン用にセット
	w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")

	// CSVのインポートエクスポートの場合
	if r.URL.Path == "/api/data/download" {
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=a.zip")

	// JSONの場合
	} else {
		w.Header().Set("Content-Type", "application/json")
	}

	// preflight用に200でいったん返す
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// セッション情報をログに吐き出す
	start := time.Now()
	infolog.Printf("%s %s %v %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	l.handler.ServeHTTP(w, r)
}

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
