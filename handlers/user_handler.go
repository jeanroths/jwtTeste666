package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jeanroths/jwtTeste666/models"
)

var users []models.User

var jwtKey = []byte("aodjfeowjedsoj134trefdsas234t")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


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
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Falha ao decodificar as credenciais", http.StatusBadRequest)
		return
	}

	// Autenticação básica de exemplo
	if creds.Username != "admin" || creds.Password != "admin" {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao gerar token: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
