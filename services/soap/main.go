package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/services/soap/handlers"
	"syscall"
	"time"
)

var (
	mongoURI    string
	dbName      string
	rabbitmqURI string
	listenAddr  string
)

type SoapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SoapBody
}

type SoapBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *SoapFault
	Content interface{}
}

type SoapFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode"`
	String  string   `xml:"faultstring"`
	Detail  string   `xml:"detail"`
}

func main() {
	initialize()

	disconnectMongo := connectDB()
	defer disconnectMongo()

	disconnectRabbitMQ := connectRabbitMQ()
	defer disconnectRabbitMQ()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "SOAP Service is healthy")
	})

	// Pokemon SOAP endpoints
	http.HandleFunc("/soap/pokemon", handlers.HandlePokemonService)
	// http.HandleFunc("/soap/user", handlers.HandleUserService)

	server := &http.Server{
		Addr:         listenAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		log.Printf("SOAP service listening on %s", listenAddr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down SOAP service...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("SOAP service was stopped successfully")
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
		listenAddr = ":8084"
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
