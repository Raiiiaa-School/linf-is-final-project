package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/services/graphql/graph"
	"project-is/services/graphql/graph/generated"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "GraphQL Service is healthy")
	}).Methods("GET")

	resolver := &graph.Resolver{}
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	})
	srv := handler.NewDefaultServer(schema)

	router.Handle("/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	server := &http.Server{
		Addr:         listenAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Printf("GraphQL service listening on %s", listenAddr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down GraphQL service...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("GraphQL service was stopped successfully")
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
		listenAddr = ":8082"
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