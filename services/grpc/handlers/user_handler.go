package handlers

import (
	"context"
	"log"
	"project-is/pkg/db"
	"project-is/pkg/models"
	pb "project-is/services/grpc/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// Check if user already exists
	existingUser, err := db.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to check for existing user: %v", err)
	}
	if existingUser != nil {
		return nil, status.Error(codes.AlreadyExists, "Username already exists")
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash the password: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to process password: %v", err)
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hash),
	}

	res, err := db.CreateUser(ctx, user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return &pb.UserResponse{
		User: convertToUserProto(user),
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid User ID: %v", err)
	}

	user, err := db.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Error getting user by ID %s: %v", req.Id, err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user: %v", err)
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.UserResponse{
		User: convertToUserProto(user),
	}, nil
}

func (s *UserServer) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.UserResponse, error) {
	user, err := db.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Printf("Error getting user by username %s: %v", req.Username, err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user: %v", err)
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.UserResponse{
		User: convertToUserProto(user),
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := db.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Printf("Error finding user during login: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	if user == nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid username or password")
	}

	// TODO: Implement proper JWT token generation
	placeholderToken := "generated.jwt.token"

	return &pb.LoginResponse{
		Token: placeholderToken,
	}, nil
}

func convertToUserProto(user *models.User) *pb.User {
	return &pb.User{
		Id:       user.ID.Hex(),
		Username: user.Username,
		Password: "", // Don't send password back to client
	}
}