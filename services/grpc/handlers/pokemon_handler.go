package handlers

import (
	"context"
	"fmt"
	"log"
	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/pkg/models"
	pb "project-is/services/grpc/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PokemonServer struct {
	pb.UnimplementedPokemonServiceServer
}

func NewPokemonServer() *PokemonServer {
	return &PokemonServer{}
}

func (s *PokemonServer) CreatePokemon(ctx context.Context, req *pb.CreatePokemonRequest) (*pb.PokemonResponse, error) {
	pokemon := &models.Pokemon{
		Name:      req.Name,
		Type:      req.Type,
		Abilities: req.Abilities,
		Height:    req.Height,
		Weight:    req.Weight,
		PokedexID: int(req.PokedexId),
	}

	res, err := db.CreatePokemon(ctx, pokemon)
	if err != nil {
		log.Printf("Error creating pokemon: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create pokemon: %v", err)
	}

	pokemon.ID = res.InsertedID.(primitive.ObjectID)

	// Publish event to RabbitMQ
	if err := messaging.PublishPokemonEvent("created", pokemon); err != nil {
		log.Printf("Error publishing pokemon created event to RabbitMQ: %v", err)
	}

	return &pb.PokemonResponse{
		Pokemon: convertToPokemonProto(pokemon),
	}, nil
}

func (s *PokemonServer) GetPokemon(ctx context.Context, req *pb.GetPokemonRequest) (*pb.PokemonResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Pokemon ID: %v", err)
	}

	pokemon, err := db.GetPokemonByID(ctx, id)
	if err != nil {
		log.Printf("Error getting pokemon by ID %s: %v", req.Id, err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve pokemon: %v", err)
	}

	if pokemon == nil {
		return nil, status.Error(codes.NotFound, "Pokemon not found")
	}

	return &pb.PokemonResponse{
		Pokemon: convertToPokemonProto(pokemon),
	}, nil
}

func (s *PokemonServer) ListPokemons(ctx context.Context, req *pb.ListPokemonsRequest) (*pb.ListPokemonsResponse, error) {
	pokemons, err := db.GetAllPokemons(ctx)
	if err != nil {
		log.Printf("Error getting all pokemons: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve pokemons: %v", err)
	}

	var protoPokemons []*pb.Pokemon
	for _, pokemon := range pokemons {
		protoPokemons = append(protoPokemons, convertToPokemonProto(&pokemon))
	}

	return &pb.ListPokemonsResponse{
		Pokemons: protoPokemons,
	}, nil
}

func (s *PokemonServer) UpdatePokemon(ctx context.Context, req *pb.UpdatePokemonRequest) (*pb.PokemonResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Pokemon ID: %v", err)
	}

	pokemon := &models.Pokemon{
		ID:        id,
		Name:      req.Pokemon.Name,
		Type:      req.Pokemon.Type,
		Abilities: req.Pokemon.Abilities,
		Height:    req.Pokemon.Height,
		Weight:    req.Pokemon.Weight,
		PokedexID: int(req.Pokemon.PokedexId),
	}

	res, err := db.UpdatePokemon(ctx, id, pokemon)
	if err != nil {
		log.Printf("Error updating pokemon with ID %s: %v", req.Id, err)
		return nil, status.Errorf(codes.Internal, "Failed to update pokemon: %v", err)
	}

	if res.ModifiedCount == 0 {
		return nil, status.Error(codes.NotFound, "Pokemon not found or no changes made")
	}

	// Publish event to RabbitMQ
	if err := messaging.PublishPokemonEvent("updated", pokemon); err != nil {
		log.Printf("Error publishing pokemon updated event to RabbitMQ: %v", err)
	}

	return &pb.PokemonResponse{
		Pokemon: convertToPokemonProto(pokemon),
	}, nil
}

func (s *PokemonServer) DeletePokemon(ctx context.Context, req *pb.DeletePokemonRequest) (*pb.DeletePokemonResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Pokemon ID: %v", err)
	}

	res, err := db.DeletePokemon(ctx, id)
	if err != nil {
		log.Printf("Error deleting pokemon with ID %s: %v", req.Id, err)
		return nil, status.Errorf(codes.Internal, "Failed to delete pokemon: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Pokemon not found")
	}

	// Publish event to RabbitMQ
	if err := messaging.PublishPokemonEvent("deleted", nil, id); err != nil {
		log.Printf("Error publishing pokemon deleted event to RabbitMQ: %v", err)
	}

	return &pb.DeletePokemonResponse{
		Success: true,
		Message: fmt.Sprintf("Pokemon with ID %s deleted successfully", req.Id),
	}, nil
}

func convertToPokemonProto(pokemon *models.Pokemon) *pb.Pokemon {
	return &pb.Pokemon{
		Id:        pokemon.ID.Hex(),
		Name:      pokemon.Name,
		Type:      pokemon.Type,
		Abilities: pokemon.Abilities,
		Height:    pokemon.Height,
		Weight:    pokemon.Weight,
		PokedexId: int32(pokemon.PokedexID),
	}
}