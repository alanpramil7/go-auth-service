package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alanpramil7/go-auth-service/internal/models"
)

func registerHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.RegisterUser(db, creds); err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registerd sucessfully"))
}
