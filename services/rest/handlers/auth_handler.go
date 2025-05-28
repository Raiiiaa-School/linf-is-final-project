package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"project-is/pkg/db"
	"project-is/pkg/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := db.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Printf("Error finding user during login: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash the password: %v", err)
	}

	if string(hash) != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	placeholderToken := "generated.jwt.token"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{Token: placeholderToken})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash the password: %v", err)
	}
	user.Password = string(hash)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.CreateUser(ctx, &user)
	if err != nil {
		log.Printf("Error creating user during registration: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User registered successfully", "id": res.InsertedID})
}
