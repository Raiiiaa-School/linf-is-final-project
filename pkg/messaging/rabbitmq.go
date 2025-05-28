package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"project-is/pkg/models"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	connection *amqp.Connection
	channel    *amqp.Channel
)

const (
	PokemonEventsExchange = "pokemon_events"
)

// Connect establishes a connection to the RabbitMQ server.
func Connect(amqpURI string) error {
	var err error
	connection, err = amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err = connection.Channel()
	if err != nil {
		connection.Close()
		return fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}

	// Declare the exchange for Pokemon events (fanout sends to all bound queues)
	err = channel.ExchangeDeclare(
		PokemonEventsExchange, // name
		"fanout",              // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		channel.Close()
		connection.Close()
		return fmt.Errorf("failed to declare RabbitMQ exchange: %w", err)
	}

	log.Println("Connected to RabbitMQ!")
	return nil
}

// Disconnect closes the RabbitMQ connection.
func Disconnect() error {
	if channel != nil {
		if err := channel.Close(); err != nil {
			log.Printf("Error closing RabbitMQ channel: %v", err)
		}
	}
	if connection != nil {
		if err := connection.Close(); err != nil {
			log.Printf("Error closing RabbitMQ connection: %v", err)
		}
	}
	return nil
}

// PokemonEvent represents a change event for a Pokemon.
type PokemonEvent struct {
	Type      string             `json:"type"` // e.g., "created", "updated", "deleted"
	Pokemon   *models.Pokemon    `json:"pokemon,omitempty"`
	PokemonID primitive.ObjectID `json:"pokemon_id,omitempty"` // For deleted events
}

// PublishPokemonEvent sends a message to RabbitMQ about a Pokemon change.
func PublishPokemonEvent(eventType string, pokemon *models.Pokemon, pokemonID ...primitive.ObjectID) error {
	event := PokemonEvent{
		Type:    eventType,
		Pokemon: pokemon,
	}
	if len(pokemonID) > 0 {
		event.PokemonID = pokemonID[0]
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal pokemon event: %w", err)
	}

	err = channel.Publish(
		PokemonEventsExchange, // exchange
		"",                    // routing key (not used in fanout exchange)
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish pokemon event: %w", err)
	}

	log.Printf("Published Pokemon event: %s", eventType)
	return nil
}
