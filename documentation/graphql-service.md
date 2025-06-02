# GraphQL Service Documentation

This document outlines the GraphQL service implementation for the Pokémon and User management system. The service provides a single GraphQL endpoint for managing Pokémon and User data, integrating with MongoDB for storage and RabbitMQ for event messaging.

## Service Overview

The GraphQL service runs on port 8082 and provides a single endpoint `/query` for all operations. A GraphQL Playground interface is available at the root path `/` for interactive API exploration.

## Schema

The service defines the following types and operations:

### Types

#### Pokemon
```graphql
type Pokemon {
  id: ID!
  name: String!
  type: [String!]!
  abilities: [String!]!
  height: Float!
  weight: Float!
  pokedexId: Int!
}
```

#### User
```graphql
type User {
  id: ID!
  username: String!
}
```

#### AuthResponse
```graphql
type AuthResponse {
  token: String!
}
```

### Queries

#### Get All Pokemons
```graphql
query {
  pokemons {
    id
    name
    type
    abilities
    height
    weight
    pokedexId
  }
}
```

#### Get Pokemon by ID
```graphql
query {
  pokemon(id: "507f1f77bcf86cd799439011") {
    id
    name
    type
    abilities
    height
    weight
    pokedexId
  }
}
```

#### Get User by ID
```graphql
query {
  user(id: "507f1f77bcf86cd799439011") {
    id
    username
  }
}
```

#### Get User by Username
```graphql
query {
  userByUsername(username: "ashketchum") {
    id
    username
  }
}
```

### Mutations

#### Create Pokemon
```graphql
mutation {
  createPokemon(input: {
    name: "Pikachu"
    type: ["Electric"]
    abilities: ["Static", "Lightning Rod"]
    height: 0.4
    weight: 6.0
    pokedexId: 25
  }) {
    id
    name
    type
    abilities
    height
    weight
    pokedexId
  }
}
```

#### Update Pokemon
```graphql
mutation {
  updatePokemon(
    id: "507f1f77bcf86cd799439011",
    input: {
      name: "Raichu"
      type: ["Electric"]
      abilities: ["Static"]
      height: 0.8
      weight: 30.0
      pokedexId: 26
    }
  ) {
    id
    name
    type
    abilities
    height
    weight
    pokedexId
  }
}
```

#### Delete Pokemon
```graphql
mutation {
  deletePokemon(id: "507f1f77bcf86cd799439011")
}
```

#### Create User
```graphql
mutation {
  createUser(input: {
    username: "ashketchum"
    password: "pikachu123"
  }) {
    id
    username
  }
}
```

#### Login
```graphql
mutation {
  login(input: {
    username: "ashketchum"
    password: "pikachu123"
  }) {
    token
  }
}
```

#### Register
```graphql
mutation {
  register(input: {
    username: "misty"
    password: "starmie123"
  }) {
    id
    username
  }
}
```

## Testing with curl

Here are examples of how to test the GraphQL endpoints using curl:

### Query All Pokemons
```bash
curl -X POST http://localhost:8082/query \
  -H 'Content-Type: application/json' \
  -d '{"query": "{ pokemons { id name type abilities height weight pokedexId } }"}'
```

### Create Pokemon
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

### Login
```bash
curl -X POST http://localhost:8082/query \
  -H 'Content-Type: application/json' \
  -d '{
    "query": "mutation($input: LoginInput!) { login(input: $input) { token } }",
    "variables": {
      "input": {
        "username": "ashketchum",
        "password": "pikachu123"
      }
    }
  }'
```

## Setup and Configuration

### Environment Variables
```
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=pokedex
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
LISTEN_ADDR=:8082
```

### Building and Running

1. Build the service:
```bash
go build -o graphql-service main.go
```

2. Run the service:
```bash
./graphql-service
```

### Docker Support

Build and run the service using Docker:

```bash
# Build the image
docker build -t pokemon-graphql-service .

# Run the container
docker run -p 8082:8082 \
  -e MONGO_URI=mongodb://mongo:27017 \
  -e RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/ \
  pokemon-graphql-service
```

## Error Handling

GraphQL errors are returned in a standardized format:

```json
{
  "errors": [
    {
      "message": "Error message",
      "path": ["fieldName"],
      "extensions": {
        "code": "ERROR_CODE"
      }
    }
  ]
}
```

Common error scenarios:
- Invalid input data
- Database connection issues
- Resource not found
- Authentication/Authorization failures

## Integration Notes

- The service shares the same MongoDB database with other services
- Pokemon events are published to RabbitMQ for system-wide consistency
- The service implements the same data models as other services
- Passwords are hashed using bcrypt before storage

## Security Considerations

- Use HTTPS in production
- Implement authentication middleware
- Never return password hashes in responses
- Validate all input data
- Use proper error handling to avoid information leakage

## Future Improvements

1. Add field-level permissions
2. Implement DataLoader for N+1 query optimization
3. Add subscription support for real-time updates
4. Implement query complexity analysis
5. Add query depth limiting
6. Implement proper cursor-based pagination
7. Add field-level error handling
8. Implement query caching

The GraphQL service provides a flexible and powerful API interface while maintaining consistency with the REST, SOAP, and gRPC services in the system.