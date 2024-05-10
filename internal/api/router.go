package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/register", makeHandler(db, registerHandler)).Methods("POST")

	return router
}

func makeHandler(db *sql.DB, fn func(http.ResponseWriter, *http.Request, *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
