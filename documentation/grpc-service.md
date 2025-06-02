# gRPC Service Documentation

This document outlines the gRPC service implementation for the Pokémon and User management system. The service uses Protocol Buffers for service definition and data serialization, and integrates with MongoDB for data storage and RabbitMQ for event messaging.

## Service Definition

The service is defined in `proto/service.proto` and provides two main services:

- PokemonService: For managing Pokémon data
- UserService: For user management and authentication

### PokemonService Methods

1. **CreatePokemon**

    - Creates a new Pokémon entry
    - Request: `CreatePokemonRequest`

    ```protobuf
    message CreatePokemonRequest {
      string name = 1;
      repeated string type = 2;
      repeated string abilities = 3;
      double height = 4;
      double weight = 5;
      int32 pokedex_id = 6;
    }
    ```

2. **GetPokemon**

    - Retrieves a Pokémon by ID
    - Request: `GetPokemonRequest`

    ```protobuf
    message GetPokemonRequest {
      string id = 1;
    }
    ```

3. **ListPokemons**

    - Lists all Pokémon with optional pagination
    - Request: `ListPokemonsRequest`

    ```protobuf
    message ListPokemonsRequest {
      int32 page_size = 1;
      int32 page_number = 2;
    }
    ```

4. **UpdatePokemon**

    - Updates an existing Pokémon
    - Request: `UpdatePokemonRequest`

    ```protobuf
    message UpdatePokemonRequest {
      string id = 1;
      Pokemon pokemon = 2;
    }
    ```

5. **DeletePokemon**
    - Deletes a Pokémon by ID
    - Request: `DeletePokemonRequest`
    ```protobuf
    message DeletePokemonRequest {
      string id = 1;
    }
    ```

### UserService Methods

1. **CreateUser**

    - Registers a new user
    - Request: `CreateUserRequest`

    ```protobuf
    message CreateUserRequest {
      string username = 1;
      string password = 2;
    }
    ```

2. **GetUser**

    - Retrieves a user by ID
    - Request: `GetUserRequest`

    ```protobuf
    message GetUserRequest {
      string id = 1;
    }
    ```

3. **GetUserByUsername**

    - Finds a user by username
    - Request: `GetUserByUsernameRequest`

    ```protobuf
    message GetUserByUsernameRequest {
      string username = 1;
    }
    ```

4. **Login**
    - Authenticates a user and returns a JWT token
    - Request: `LoginRequest`
    ```protobuf
    message LoginRequest {
      string username = 1;
      string password = 2;
    }
    ```

## Setup and Installation

### Prerequisites

- Go 1.24 or later
- Protocol Buffers compiler (protoc)
- MongoDB
- RabbitMQ

### Environment Variables

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=pokedex
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
LISTEN_ADDR=:8083
```

### Generating Protocol Buffers Code

1. Install the required Go plugins:

    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

2. Run the generation script:
    ```bash
    ./generate.sh
    ```

### Building and Running

1. Build the service:

    ```bash
    go build -o grpc-service main.go
    ```

2. Run the service:
    ```bash
    ./grpc-service
    ```

## Using the Service

### Client Examples

Here's a Go client example for creating a Pokémon:

```go
package main

import (
    "context"
    "log"
    "time"
    pb "project-is/services/grpc/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewPokemonServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    pokemon, err := client.CreatePokemon(ctx, &pb.CreatePokemonRequest{
        Name:      "Pikachu",
        Type:      []string{"Electric"},
        Abilities: []string{"Static", "Lightning Rod"},
        Height:    0.4,
        Weight:    6.0,
        PokedexId: 25,
    })
    if err != nil {
        log.Fatalf("Could not create pokemon: %v", err)
    }
    log.Printf("Created Pokemon: %v", pokemon)
}
```

### Error Handling

The service uses standard gRPC status codes:

- `INVALID_ARGUMENT`: Invalid input data
- `NOT_FOUND`: Resource not found
- `ALREADY_EXISTS`: Duplicate resource (e.g., username)
- `INTERNAL`: Server-side errors
- `UNAUTHENTICATED`: Invalid credentials

### Events

The service publishes events to RabbitMQ for:

- Pokemon creation
- Pokemon updates
- Pokemon deletion

Events are published to the `pokemon_events` exchange with the following routing keys:

- `pokemon.created`
- `pokemon.updated`
- `pokemon.deleted`

## Security

- Passwords are hashed using bcrypt before storage
- The service supports JWT-based authentication (placeholder implementation)
- Transport security should be configured in production using TLS

## Docker Support

The service can be built and run using Docker:

```bash
# Build the image
docker build -t pokemon-grpc-service .

# Run the container
docker run -p 8083:8083 \
  -e MONGO_URI=mongodb://mongo:27017 \
  -e RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/ \
  pokemon-grpc-service
```

## Integration with Other Services

The gRPC service shares the same MongoDB database and RabbitMQ message broker with other services (REST, GraphQL, SOAP) to maintain data consistency across the system.

## Future Improvements

1. Implement proper JWT token generation and validation
2. Add support for streaming endpoints (e.g., real-time updates)
3. Implement request rate limiting
4. Add metrics and monitoring
5. Implement proper TLS configuration
6. Add support for bulk operations

## Testing

To test the service, you can use:

- `grpcurl` for manual testing
- The provided Go client examples
- Integration tests (to be implemented)
- Load testing tools like `ghz`

Example using grpcurl:

```bash
# List services
grpcurl -plaintext localhost:8083 list

# Create a Pokemon
grpcurl -plaintext -d '{"name": "Pikachu", "type": ["Electric"], "abilities": ["Static"], "height": 0.4, "weight": 6.0, "pokedex_id": 25}' localhost:8083 grpc.PokemonService/CreatePokemon
```

---

## General grpcurl Commands

First, you can list all services or methods of a specific service:

1. List all available services:

```bash
grpcurl -plaintext localhost:8083 list
```

(Expected to show grpc.PokemonService and grpc.UserService among others like reflection service)

2. Describe a service (e.g., PokemonService):

```bash
grpcurl -plaintext localhost:8083 describe grpc.PokemonService
```

3. Describe a specific method (e.g., CreatePokemon):

```bash
grpcurl -plaintext localhost:8083 describe grpc.PokemonService.CreatePokemon
```

PokemonService Methods

1. CreatePokemon

```bash
grpcurl -plaintext -d '{
"name": "Pikachu",
"type": ["Electric"],
"abilities": ["Static", "Lightning Rod"],
"height": 0.4,
"weight": 6.0,
"pokedex_id": 25
}' localhost:8083 grpc.PokemonService/CreatePokemon
```

(Save the returned id from the response for use in other commands)

2. GetPokemon
   Replace <pokemon_id_here> with an actual Pokémon ID.

```bash
grpcurl -plaintext -d '{
"id": "<pokemon_id_here>"
}' localhost:8083 grpc.PokemonService/GetPokemon
```

3. ListPokemons

- Without pagination parameters (server might use defaults or return all):

```bash
grpcurl -plaintext -d '{}' localhost:8083 grpc.PokemonService/ListPokemons
```

- With pagination parameters:

```bash
grpcurl -plaintext -d '{
  "page_size": 5,
      "page_number": 1
}' localhost:8083 grpc.PokemonService/ListPokemons
```

4. UpdatePokemon
   Replace <pokemon_id_to_update_here> with an actual Pokémon ID. The nested pokemon object should contain all fields for the Pokémon as you want them to be after the update.

```bash
grpcurl -plaintext -d '{
"id": "<pokemon_id_to_update_here>",
"pokemon": {
"id": "<pokemon_id_to_update_here>", # Often, this nested ID is ignored or validated against the outer ID
"name": "Pikachu Libre",
"type": ["Electric", "Fighting"],
"abilities": ["Static", "Lightning Rod", "Tough Claws"],
"height": 0.45,
"weight": 6.5,
"pokedex_id": 25
}
}' localhost:8083 grpc.PokemonService/UpdatePokemon
```

5. DeletePokemon
   Replace <pokemon_id_to_delete_here> with an actual Pokémon ID.

```bash
grpcurl -plaintext -d '{
"id": "<pokemon_id_to_delete_here>"
}' localhost:8083 grpc.PokemonService/DeletePokemon
```

UserService Methods

1. CreateUser

```bash
grpcurl -plaintext -d '{
"username": "ashketchum",
"password": "Pikachu<3"
}' localhost:8083 grpc.UserService/CreateUser
```

(Save the returned id from the response for use in GetUser)

2. GetUser
   Replace <user_id_here> with an actual User ID.

```bash
grpcurl -plaintext -d '{
"id": "<user_id_here>"
}' localhost:8083 grpc.UserService/GetUser
```

3. GetUserByUsername

```bash
grpcurl -plaintext -d '{
"username": "ashketchum"
}' localhost:8083 grpc.UserService/GetUserByUsername
```

4. Login

```bash
grpcurl -plaintext -d '{
"username": "ashketchum",
"password": "Pikachu<3"
}' localhost:8083 grpc.UserService/Login
```

(This should return a token if the credentials are correct and JWT generation is implemented)
