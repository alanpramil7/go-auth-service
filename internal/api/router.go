package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// Function to create new route
func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", makeHandler(db, registerHandler)).Methods("POST")
	router.HandleFunc("/login", makeHandler(db, loginHandler)).Methods("POST")
	router.Handle("/welcome", ValidateTokenMiddleware(http.HandlerFunc(welcomeHandler))).Methods("GET")
	return router
}

// Helper Function
func makeHandler(db *sql.DB, fn func(http.ResponseWriter, *http.Request, *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
