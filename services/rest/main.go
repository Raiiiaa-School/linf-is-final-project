package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/services/rest/handlers"

	"github.com/gorilla/mux"
)

var (
	mongoURI    string
	dbName      string
	rabbitmqURI string
	listenAddr  string
)

func main() {
	initialize()

	disconnectMongo := connectDB()
	defer disconnectMongo()

	disconnectRabbitMQ := connectRabbitMQ()
	defer disconnectRabbitMQ()

	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "REST Service is healthy")
	}).Methods("GET")

	r.HandleFunc("/pokemons", handlers.CreatePokemon).Methods("POST")
	r.HandleFunc("/pokemons", handlers.GetAllPokemons).Methods("GET")
	r.HandleFunc("/pokemons/{id}", handlers.GetPokemonByID).Methods("GET")
	r.HandleFunc("/pokemons/{id}", handlers.UpdatePokemon).Methods("PUT")
	r.HandleFunc("/pokemons/{id}", handlers.DeletePokemon).Methods("DELETE")

	// User Routes
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/users/username/{username}", handlers.GetUserByUsername).Methods("GET")

	// Authentication Routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	server := &http.Server{
		Addr:         listenAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run server in a goroutine so that it doesn't block other operations
	go func() {
		log.Printf("REST service listening on %s", listenAddr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until we receive our signal
	<-c
	log.Println("Shutting down REST service...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("REST service was stopped successfully")
	os.Exit(0)
}

func initialize() {
	mongoURI = os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	dbName = os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "pokedex"
	}

	rabbitmqURI = os.Getenv("RABBITMQ_URI")
	if rabbitmqURI == "" {
		log.Fatal("RABBITMQ_URI environment variable not set")
	}

	listenAddr = os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8081"
	}
}

func connectDB() func() {
	if err := db.Connect(mongoURI, dbName); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	return func() {
		if err := db.Disconnect(); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}

func connectRabbitMQ() func() {
	if err := messaging.Connect(rabbitmqURI); err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	return func() {
		if err := messaging.Disconnect(); err != nil {
			log.Printf("Error disconnecting from RabbitMQ: %v", err)
		}
	}
}
