package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/pkg/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	var pokemon models.Pokemon
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.CreatePokemon(ctx, &pokemon)
	if err != nil {
		log.Printf("Error creating pokemon: %v", err)
		http.Error(w, "Failed to create pokemon", http.StatusInternalServerError)
		return
	}

	if err := messaging.PublishPokemonEvent("created", &pokemon); err != nil {
		log.Printf("Error publishing pokemon created event to RabbitMQ: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Pokemon created successfully",
		"id":      res.InsertedID,
	})
}

func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pokemons, err := db.GetAllPokemons(ctx)
	if err != nil {
		log.Printf("Error getting all pokemons: %v", err)
		http.Error(w, "Failed to retrieve pokemons", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid Pokemon ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pokemon, err := db.GetPokemonByID(ctx, id)
	if err != nil {
		log.Printf("Error getting pokemon by ID %s: %v", idStr, err)
		http.Error(w, "Failed to retrieve pokemon", http.StatusInternalServerError)
		return
	}

	if pokemon == nil {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid Pokemon ID", http.StatusBadRequest)
		return
	}

	var pokemon models.Pokemon
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// You might want to set the ID from the URL parameters in the decoded object
	pokemon.ID = id

	res, err := db.UpdatePokemon(ctx, id, &pokemon)
	if err != nil {
		log.Printf("Error updating pokemon with ID %s: %v", idStr, err)
		http.Error(w, "Failed to update pokemon", http.StatusInternalServerError)
		return
	}

	if res.ModifiedCount == 0 {
		http.Error(w, "Pokemon not found or no changes made", http.StatusNotFound)
		return
	}

	// Publish event to RabbitMQ
	if err := messaging.PublishPokemonEvent("updated", &pokemon); err != nil {
		log.Printf("Error publishing pokemon updated event to RabbitMQ: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Pokemon updated successfully"})
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid Pokemon ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.DeletePokemon(ctx, id)
	if err != nil {
		log.Printf("Error deleting pokemon with ID %s: %v", idStr, err)
		http.Error(w, "Failed to delete pokemon", http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}

	// Publish event to RabbitMQ
	if err := messaging.PublishPokemonEvent("deleted", nil, id); err != nil {
		log.Printf("Error publishing pokemon deleted event to RabbitMQ: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Pokemon deleted successfully"})
}
