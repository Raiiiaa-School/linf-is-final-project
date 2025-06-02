package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/services/grpc/handlers"
	pb "project-is/services/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// Register our services
	pb.RegisterPokemonServiceServer(server, handlers.NewPokemonServer())
	pb.RegisterUserServiceServer(server, handlers.NewUserServer())

	// Register reflection service on gRPC server
	reflection.Register(server)

	// Run server in a goroutine
	go func() {
		log.Printf("gRPC server listening on %s", listenAddr)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal
	<-c
	log.Println("Shutting down gRPC server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Gracefully stop the server
	server.GracefulStop()
	if err := listener.Close(); err != nil {
		log.Printf("Error closing listener: %v", err)
	}

	<-ctx.Done()
	log.Println("gRPC server was stopped gracefully")
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
		listenAddr = ":50051"
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
