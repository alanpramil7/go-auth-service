package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alanpramil7/go-auth-service/internal/models"
)

// Function to handle register route
func registerHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.RegisterUser(db, creds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registerd sucessfully"))
}

// Function for login handler
func loginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := models.AutheneticateUser(db, creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
	w.Write([]byte("User loged in sucessfully"))
}

// Welcome handler to return a greeting message with the username
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    claims, ok := r.Context().Value("userClaims").(*models.Claims)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    message := "Welcome, " + claims.UserName + "!"
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": message})
}
