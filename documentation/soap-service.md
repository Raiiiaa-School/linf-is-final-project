# SOAP Service Documentation

This document outlines the SOAP service implementation for the Pokémon and User management system. The service provides XML-based SOAP endpoints for managing Pokémon and User data, integrating with MongoDB for storage and RabbitMQ for event messaging.

## Service Overview

The SOAP service runs on port 8084 and provides two main endpoints:
- `/soap/pokemon` - For Pokémon management operations
- `/soap/user` - For user management operations

A health check endpoint is also available at `/health`.

## SOAP Endpoints

### Pokemon Service

#### Create Pokemon

- **Endpoint**: `/soap/pokemon`
- **SOAP Action**: `create`
- **Request Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
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
</soap:Envelope>
```
- **Response Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <PokemonResponse>
      <status>success</status>
      <message>Pokemon created successfully</message>
      <pokemon>
        <id>507f1f77bcf86cd799439011</id>
        <name>Pikachu</name>
        <type>Electric</type>
        <abilities>Static</abilities>
        <abilities>Lightning Rod</abilities>
        <height>0.4</height>
        <weight>6.0</weight>
        <pokedex_id>25</pokedex_id>
      </pokemon>
    </PokemonResponse>
  </soap:Body>
</soap:Envelope>
```

#### Get Pokemon

- **Endpoint**: `/soap/pokemon`
- **SOAP Action**: `get`
- **Request Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <PokemonRequest>
      <action>get</action>
      <pokemon>
        <id>507f1f77bcf86cd799439011</id>
      </pokemon>
    </PokemonRequest>
  </soap:Body>
</soap:Envelope>
```
- **Response Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <PokemonResponse>
      <status>success</status>
      <pokemon>
        <id>507f1f77bcf86cd799439011</id>
        <name>Pikachu</name>
        <type>Electric</type>
        <abilities>Static</abilities>
        <abilities>Lightning Rod</abilities>
        <height>0.4</height>
        <weight>6.0</weight>
        <pokedex_id>25</pokedex_id>
      </pokemon>
    </PokemonResponse>
  </soap:Body>
</soap:Envelope>
```

#### Update Pokemon

- **Endpoint**: `/soap/pokemon`
- **SOAP Action**: `update`
- **Request Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <PokemonRequest>
      <action>update</action>
      <pokemon>
        <id>507f1f77bcf86cd799439011</id>
        <name>Raichu</name>
        <type>Electric</type>
        <abilities>Static</abilities>
        <height>0.8</height>
        <weight>30.0</weight>
        <pokedex_id>26</pokedex_id>
      </pokemon>
    </PokemonRequest>
  </soap:Body>
</soap:Envelope>
```

#### Delete Pokemon

- **Endpoint**: `/soap/pokemon`
- **SOAP Action**: `delete`
- **Request Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <PokemonRequest>
      <action>delete</action>
      <pokemon>
        <id>507f1f77bcf86cd799439011</id>
      </pokemon>
    </PokemonRequest>
  </soap:Body>
</soap:Envelope>
```

### User Service

#### Create User

- **Endpoint**: `/soap/user`
- **SOAP Action**: `create`
- **Request Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <UserRequest>
      <action>create</action>
      <user>
        <username>ashketchum</username>
        <password>pikachu123</password>
      </user>
    </UserRequest>
  </soap:Body>
</soap:Envelope>
```

#### Get User

- **Endpoint**: `/soap/user`
- **SOAP Action**: `get`
- **Request Example**:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <UserRequest>
      <action>get</action>
      <user>
        <id>507f1f77bcf86cd799439011</id>
      </user>
    </UserRequest>
  </soap:Body>
</soap:Envelope>
```

## Testing with curl

Here are some examples of how to test the SOAP endpoints using curl:

### Create Pokemon
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

### Create User
```bash
curl -X POST http://localhost:8084/soap/user \
  -H 'Content-Type: text/xml' \
  -d '<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <UserRequest>
      <action>create</action>
      <user>
        <username>ashketchum</username>
        <password>pikachu123</password>
      </user>
    </UserRequest>
  </soap:Body>
</soap:Envelope>'
```

## Error Handling

The service follows standard SOAP fault conventions for error responses:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <soap:Fault>
      <faultcode>Server</faultcode>
      <faultstring>Error message</faultstring>
      <detail>Detailed error information</detail>
    </soap:Fault>
  </soap:Body>
</soap:Envelope>
```

Common error scenarios:
- Invalid request format
- Database connection issues
- Resource not found
- Authentication/Authorization failures

## Setup and Configuration

### Environment Variables

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=pokedex
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
LISTEN_ADDR=:8084
```

### Building and Running

1. Build the service:
```bash
go build -o soap-service main.go
```

2. Run the service:
```bash
./soap-service
```

### Docker Support

Build and run the service using Docker:

```bash
# Build the image
docker build -t pokemon-soap-service .

# Run the container
docker run -p 8084:8084 \
  -e MONGO_URI=mongodb://mongo:27017 \
  -e RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/ \
  pokemon-soap-service
```

## Integration Notes

- The service shares the same MongoDB database with other services
- Pokemon events are published to RabbitMQ for system-wide consistency
- The service implements the same data models as other services for compatibility
- Passwords are hashed using bcrypt before storage

## Security Considerations

- Use HTTPS in production
- Implement authentication middleware
- Never return password hashes in responses
- Validate all input data
- Use proper error handling to avoid information leakage

## Future Improvements

1. Add WSDL generation
2. Implement request validation
3. Add authentication middleware
4. Support bulk operations
5. Add detailed logging
6. Implement metrics collection
7. Add rate limiting
8. Improve error messages and codes

The SOAP service provides a traditional XML-based interface for systems that require SOAP compatibility while maintaining consistency with the REST, GraphQL, and gRPC services in the system.