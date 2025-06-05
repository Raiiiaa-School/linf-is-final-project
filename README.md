# linf-is-final-project

# Pokémon and User Management System

This repository contains the implementation of a Pokémon and User management system with multiple service interfaces: REST, GraphQL, gRPC, and SOAP. Each service integrates with MongoDB for data storage and RabbitMQ for event messaging, ensuring consistency across the system. This README provides an overview of the services and instructions for setup, usage, and testing.

## Table of Contents
- [Overview](#overview)
- [Services](#services)
  - [REST API](#rest-api)
  - [GraphQL Service](#graphql-service)
  - [gRPC Service](#grpc-service)
  - [SOAP Service](#soap-service)
- [Setup and Configuration](#setup-and-configuration)
- [Testing](#testing)
- [Security Considerations](#security-considerations)
- [Future Improvements](#future-improvements)
- [Services Documentation](#)
    - [Rest](./documentation/rest-service.md)
    - [Soap](./documentation/soap-service.md)
    - [GRPC](./documentation/grpc-service.md)
    - [GraphQL](./documentation/graphql-service.md)
    - [RoadMap](./documentation/roadmap.md)


## Overview
The system provides APIs for managing Pokémon and User data, supporting CRUD operations (Create, Read, Update, Delete) and authentication. All services share the same MongoDB database (`pokedex`) and publish events to RabbitMQ for system-wide consistency.

## Services

### REST API
The REST API is served at `http://localhost:8081` (configurable via `LISTEN_ADDR`). It provides endpoints for Pokémon and User management, as well as authentication.

**Key Endpoints**:
- **Pokémon**: `/pokemons` (GET, POST), `/pokemons/{id}` (GET, PUT, DELETE)
- **User**: `/users` (POST), `/users/{id}` (GET), `/users/username/{username}` (GET)
- **Authentication**: `/login` (POST), `/register` (POST)

**Documentation**: [REST API Endpoints](docs/rest_api_endpoints.md)

**Example (Create Pokémon)**:
```bash
curl -X POST http://localhost:8081/pokemons \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Pikachu",
    "type": ["Electric"],
    "abilities": ["Static", "Lightning Rod"],
    "height": 0.4,
    "weight": 6.0,
    "pokedex_id": 25
  }'
```

### GraphQL Service
The GraphQL service runs on port `8082` with a single endpoint `/query` and a Playground interface at `/`.

**Key Operations**:
- **Queries**: `pokemons`, `pokemon(id)`, `user(id)`, `userByUsername(username)`
- **Mutations**: `createPokemon`, `updatePokemon`, `deletePokemon`, `createUser`, `login`, `register`

**Documentation**: [GraphQL Service Documentation](docs/graphql_service.md)

**Example (Create Pokémon)**:
```bash
curl -X POST http://localhost:8082/query \
  -H 'Content-Type: application/json' \
  -d '{
    "query": "mutation($input: CreatePokemonInput!) { createPokemon(input: $input) { id name type abilities height weight pokedexId } }",
    "variables": {
      "input": {
        "name": "Pikachu",
        "type": ["Electric"],
        "abilities": ["Static", "Lightning Rod"],
        "height": 0.4,
        "weight": 6.0,
        "pokedexId": 25
      }
    }
  }'
```

### gRPC Service
The gRPC service runs on port `8083` and provides `PokemonService` and `UserService` for managing Pokémon and User data.

**Key Methods**:
- **PokemonService**: `CreatePokemon`, `GetPokemon`, `ListPokemons`, `UpdatePokemon`, `DeletePokemon`
- **UserService**: `CreateUser`, `GetUser`, `GetUserByUsername`, `Login`

**Documentation**: [gRPC Service Documentation](docs/grpc_service.md)

**Example (Create Pokémon with grpcurl)**:
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

### SOAP Service
The SOAP service runs on port `8084` with endpoints `/soap/pokemon` and `/soap/user` for Pokémon and User management.

**Key Operations**:
- **Pokemon Service**: `create`, `get`, `update`, `delete`
- **User Service**: `create`, `get`

**Documentation**: [SOAP Service Documentation](docs/soap_service.md)

**Example (Create Pokémon)**:
```bash
curl -X POST http://localhost:8084/soap/pokemon \
  -H 'Content-Type: text/xml' \
  -d '<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <PokemonRequest>
      <action>create</action>
      <pokemon>
        <name>Pikachu</name>
        <type>Electric</type>
        <abilities>Static</abilities>
        <abilities>Lightning Rod</abilities>
        <height>0.4</height>
        <weight>6.0</weight>
        <pokedex_id>25</pokedex_id>
      </pokemon>
    </PokemonRequest>
  </soap:Body>
</soap:Envelope>'
```

## Setup and Configuration

### Prerequisites
- Go 1.24 or later
- MongoDB
- RabbitMQ
- Protocol Buffers compiler (protoc) for gRPC
- Docker (optional)

### Environment Variables
All services use the following environment variables:
```bash
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=pokedex
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
LISTEN_ADDR=:8081  # :8082 for GraphQL, :8083 for gRPC, :8084 for SOAP
```

### Building and Running
For each serviced:
1. **Build**:
   ```bash
   go build -o <service-name> main.go
   ```
2. **Run**:
   ```bash
   ./<service-name>
   ```

### Docker Support
Each service can be built and run using Docker:
```bash
# Build
docker build -t pokemon-<service>-service .
# Run
docker run -p <port>:<port> \
  -e MONGO_URI=mongodb://mongo:27017 \
  -e RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/ \
  pokemon-<service>-service
```

Replace `<service>` with `rest`, `graphql`, `grpc`, or `soap`, and `<port>` with `8081`, `8082`, `8083`, or `8084`, respectively.

## Testing
- **REST**: Use `curl` or tools like Postman.
- **GraphQL**: Use the GraphQL Playground at `http://localhost:8082/` or `curl`.
- **gRPC**: Use `grpcurl` or a custom Go client.
- **SOAP**: Use `curl` or SOAP clients like SoapUI.

See the respective documentation files for detailed testing examples.

## Security Considerations
- Use HTTPS in production.
- Implement authentication middleware for protected endpoints.
- Avoid returning sensitive data (e.g., password hashes) in responses.
- Validate all input data to prevent injection attacks.
- Use proper error handling to avoid information leakage.

## Future Improvements
- Add field-level permissions and query complexity analysis (GraphQL).
- Implement proper JWT token generation and validation (gRPC, REST).
- Add WSDL generation and request validation (SOAP).
- Support real-time updates via subscriptions or streaming.
- Implement pagination, rate limiting, and metrics collection across services.
  
---

Trabalho realizado por:

- Rúben Alves | 210100338
- Rodrigo Ventura | 240001108
