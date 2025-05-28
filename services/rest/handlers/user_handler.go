package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"project-is/pkg/db"
	"project-is/pkg/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	existingUser, err := db.GetUserByUsername(ctx, user.Username)
	if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	res, err := db.CreateUser(ctx, &user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created successfully", "id": res.InsertedID})
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := db.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Error getting user by ID %s: %v", idStr, err)
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := db.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Error getting user by username %s: %v", username, err)
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
