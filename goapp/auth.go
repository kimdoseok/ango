package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type (
	AuthService struct {
		db *gorm.DB
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Authkey  string `json:"authkey"`
		Valid    bool   `json:"valid"`
	}
)

var secretKey = []byte("lovemeanseverhavingtosayyouaresorry")

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (auth *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	// Implementation for login
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// check login in database below
	validlogin := false

	token := ""
	if validlogin == true {
		token, err = createToken(loginReq.Email)
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			return
		}

	}

	loginData := LoginRequest{
		Email:    loginReq.Email,
		Password: "",
		Authkey:  token,
		Valid:    validlogin,
	}
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(jsonData))

}

func (auth *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	// Implementation for logout
	loginData := LoginRequest{
		Email:    "",
		Password: "",
		Authkey:  "",
		Valid:    false,
	}
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}
	w.Write([]byte(jsonData))
}
