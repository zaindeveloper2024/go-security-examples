package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("secret-key"))
)

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", HealthHandler).Methods("GET")
	r.HandleFunc("/api/login", loginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/logout", logoutHandler).Methods("POST", "OPTIONS")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port 8080")

	log.Fatal(srv.ListenAndServe())
}
