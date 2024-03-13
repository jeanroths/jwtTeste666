package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeanroths/jwtTeste666/models"
)

var users []models.User

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Aqui você implementaria a lógica de login e geração de token JWT
	// Por simplicidade, vamos apenas retornar um status 200 OK
	w.WriteHeader(http.StatusOK)
}
