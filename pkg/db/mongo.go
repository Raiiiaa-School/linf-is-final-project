package db

import (
	"context"
	"fmt"
	"log"
	"project-is/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	database *mongo.Database
	pokemonsCollection *mongo.Collection
	usersCollection *mongo.Collection
)

func Connect(mongoURI string, dbName string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("Failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		if disconnectErr := client.Disconnect(context.Background()); disconnectErr != nil {
			log.Printf("Error disconnecting from MongoDB after ping failure: %v", disconnectErr)
		}
		return fmt.Errorf("Failed to ping MongoDB: %w", err)
	}
	log.Println("Connected to MongoDB!")

	database = client.Database(dbName)
	pokemonsCollection = database.Collection("pokemons")
	usersCollection = database.Collection("users")

	return nil
}

func Disconnect() error {
	if client == nil {
		return fmt.Errorf("MongoDB client is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}

// CreatePokemon inserts a new Pokemon into the database.
func CreatePokemon(ctx context.Context, pokemon *models.Pokemon) (*mongo.InsertOneResult, error) {
	res, err := pokemonsCollection.InsertOne(ctx, pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to create pokemon: %w", err)
	}
	return res, nil
}

// GetPokemonByID retrieves a Pokemon by its ID.
func GetPokemonByID(ctx context.Context, id primitive.ObjectID) (*models.Pokemon, error) {
	var pokemon models.Pokemon
	err := pokemonsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&pokemon)
	if err == mongo.ErrNoDocuments {
		return nil, nil // Pokemon not found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get pokemon by ID: %w", err)
	}
	return &pokemon, nil
}

// GetPokemonByName retrieves a Pokemon by its name.
func GetPokemonByName(ctx context.Context, name string) (*models.Pokemon, error) {
	var pokemon models.Pokemon
	err := pokemonsCollection.FindOne(ctx, bson.M{"name": name}).Decode(&pokemon)
	if err == mongo.ErrNoDocuments {
		return nil, nil // Pokemon not found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get pokemon by name: %w", err)
	}
	return &pokemon, nil
}


// GetAllPokemons retrieves all Pokemons from the database.
func GetAllPokemons(ctx context.Context) ([]models.Pokemon, error) {
	var pokemons []models.Pokemon
	cursor, err := pokemonsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all pokemons: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &pokemons); err != nil {
		return nil, fmt.Errorf("failed to decode pokemons: %w", err)
	}
	return pokemons, nil
}

// UpdatePokemon updates an existing Pokemon.
func UpdatePokemon(ctx context.Context, id primitive.ObjectID, updatedPokemon *models.Pokemon) (*mongo.UpdateResult, error) {
	res, err := pokemonsCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedPokemon})
	if err != nil {
		return nil, fmt.Errorf("failed to update pokemon: %w", err)
	}
	return res, nil
}

// DeletePokemon deletes a Pokemon by its ID.
func DeletePokemon(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	res, err := pokemonsCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to delete pokemon: %w", err)
	}
	return res, nil
}

// --- User CRUD Operations ---

// CreateUser inserts a new User into the database.
func CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	res, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return res, nil
}

// GetUserByID retrieves a User by their ID.
func GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := usersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // User not found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}

// GetUserByUsername retrieves a User by their username.
func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // User not found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}
