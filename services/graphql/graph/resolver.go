package graph

import (
	"context"
	"fmt"
	"log"
	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/pkg/models" // Assuming models.Pokemon.ID and models.User.ID are e.g. primitive.ObjectID
	"project-is/services/graphql/graph/generated"
	"project-is/services/graphql/graph/model"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	// Any dependencies your resolvers need, e.g.:
	// DB *db.Client
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} } // <<< ADD THIS

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} } // <<< ADD THIS

// mutationResolver implements the MutationResolver interface.
type mutationResolver struct{ *Resolver }

// queryResolver implements the QueryResolver interface.
type queryResolver struct{ *Resolver }

// pokemonResolver implements the PokemonResolver interface.
type pokemonResolver struct{ *Resolver } // <<< ADD THIS STRUCT

// userResolver implements the UserResolver interface.
type userResolver struct{ *Resolver } // <<< ADD THIS STRUCT

// --- Implement PokemonResolver methods ---

// ID is the resolver for the ID field of the Pokemon type.
// It converts the internal ID (likely primitive.ObjectID) to a string.
func (r *pokemonResolver) ID(ctx context.Context, obj *models.Pokemon) (string, error) { // <<< ADD THIS METHOD
	if obj == nil {
		// Depending on schema (ID! vs ID), handle appropriately.
		// If ID!, an error or empty string might be expected if obj is unexpectedly nil.
		return "", fmt.Errorf("cannot resolve ID for nil Pokemon object")
	}
	return obj.ID.Hex(), nil // Assuming obj.ID is primitive.ObjectID
}

// Add other resolvers for Pokemon fields if models.Pokemon fields don't directly match schema
// or need custom logic. For example, if you had a 'fullName' field in GraphQL
// that was computed from 'firstName' and 'lastName' in your model.
// Based on your generated PokemonResolver, only ID is needed for now.

// --- Implement UserResolver methods ---

// ID is the resolver for the ID field of the User type.
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) { // <<< ADD THIS METHOD
	if obj == nil {
		return "", fmt.Errorf("cannot resolve ID for nil User object")
	}
	return obj.ID.Hex(), nil // Assuming obj.ID is primitive.ObjectID
}

// --- Your existing Query resolver methods (no changes needed here for THIS error) ---
func (r *queryResolver) Pokemons(ctx context.Context) ([]*models.Pokemon, error) {
	pokemons, err := db.GetAllPokemons(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemons: %v", err)
	}

	result := make([]*models.Pokemon, len(pokemons))
	for i := range pokemons {
		result[i] = &pokemons[i] // gqlgen will use PokemonResolver for fields of these items
	}
	return result, nil
}

func (r *queryResolver) Pokemon(ctx context.Context, id string) (*models.Pokemon, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid pokemon ID: %v", err)
	}
	pokemon, err := db.GetPokemonByID(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon: %v", err)
	}
	return pokemon, nil // gqlgen will use PokemonResolver for fields of this item
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}
	user, err := db.GetUserByID(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil // gqlgen will use UserResolver for fields of this item
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := db.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil // gqlgen will use UserResolver for fields of this item
}

// --- Your existing Mutation resolver methods (no changes needed here for THIS error) ---
// (CreatePokemon, UpdatePokemon, DeletePokemon, CreateUser, Login, Register)
// ... (rest of your mutationResolver methods) ...
func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.CreatePokemonInput) (*models.Pokemon, error) {
	pokemon := &models.Pokemon{
		ID:        primitive.NewObjectID(), // Ensure ID is initialized if needed before CreatePokemon call
		Name:      input.Name,
		Type:      input.Type,
		Abilities: input.Abilities,
		Height:    input.Height,
		Weight:    input.Weight,
		PokedexID: input.PokedexID,
	}

	res, err := db.CreatePokemon(ctx, pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to create pokemon: %v", err)
	}

	// If db.CreatePokemon doesn't populate pokemon.ID, but returns it in res:
	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to get inserted ID for pokemon")
	}
	pokemon.ID = insertedID


	if err := messaging.PublishPokemonEvent("created", pokemon); err != nil {
		log.Printf("Error publishing pokemon created event to RabbitMQ: %v", err)
	}

	return pokemon, nil // gqlgen will use PokemonResolver for fields of this item
}

// ... (ensure other mutation methods that return *models.Pokemon or *models.User are also consistent)
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &models.User{
		ID:       primitive.NewObjectID(), // Ensure ID is initialized
		Username: input.Username,
		Password: string(hash),
	}

	res, err := db.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to get inserted ID for user")
	}
	user.ID = insertedID
	return user, nil
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid pokemon ID: %v", err)
	}

	deleted, err := db.DeletePokemon(ctx, objectID)
	if err != nil {
		return false, fmt.Errorf("failed to delete pokemon: %v", err)
	}

	if deleted.DeletedCount != 0 {
		// Publish deletion event to RabbitMQ
		pokemon := &models.Pokemon{ID: objectID}
		if err := messaging.PublishPokemonEvent("deleted", pokemon); err != nil {
			log.Printf("Error publishing pokemon deleted event to RabbitMQ: %v", err)
			// Don't return error since deletion was successful
		}
	}

	return deleted.DeletedCount != 0, nil
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, id string, input model.UpdatePokemonInput) (*models.Pokemon, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid pokemon ID: %v", err)
	}

	// Get existing pokemon to update
	pokemon, err := db.GetPokemonByID(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon: %v", err)
	}

	// Update fields if provided in input
	if input.Name != nil {
		pokemon.Name = *input.Name
	}
	if input.Type != nil {
		pokemon.Type = *&input.Type
	}
	if input.Abilities != nil {
		pokemon.Abilities = input.Abilities
	}
	if input.Height != nil {
		pokemon.Height = *input.Height
	}
	if input.Weight != nil {
		pokemon.Weight = *input.Weight
	}
	if input.PokedexID != nil {
		pokemon.PokedexID = *input.PokedexID
	}

	// Update in database
	_,  err = db.UpdatePokemon(ctx, objectID, pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to update pokemon: %v", err)
	}

	// Publish update event
	if err := messaging.PublishPokemonEvent("updated", pokemon); err != nil {
		log.Printf("Error publishing pokemon updated event to RabbitMQ: %v", err)
	}

	return pokemon, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := db.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, fmt.Errorf("authentication failed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, fmt.Errorf("authentication failed")
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &model.AuthResponse{
		Token: tokenString,
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	// Check if username already exists
	existingUser, err := db.GetUserByUsername(ctx, input.Username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Create new user - reuse CreateUser logic
	return r.CreateUser(ctx, input)
}
