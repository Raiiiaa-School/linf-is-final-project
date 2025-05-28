# REST API Endpoints

This document outlines the REST API endpoints for the Pokémon and User management system. These endpoints can be used by the frontend to interact with the backend services, which are connected to a MongoDB database and use RabbitMQ for messaging. The API is served at the address specified by the `LISTEN_ADDR` environment variable (default: `:8081`).

## Base URL

The base URL for all endpoints is `http://<host>:<port>` (e.g., `http://localhost:8081`).

## Endpoints

### Health Check

-   **Endpoint**: `/health`
-   **Method**: `GET`
-   **Description**: Checks the health status of the REST service.
-   **Response**:
    -   **Status Code**: `200 OK`
    -   **Body**: Plain text response: `"REST Service is healthy"`
-   **Usage**: Use this endpoint to verify that the REST service is running.

### Pokémon Endpoints

#### Create Pokémon

-   **Endpoint**: `/pokemons`
-   **Method**: `POST`
-   **Description**: Creates a new Pokémon in the database.
-   **Request Body**:
    ```json
    {
      "name": "string",
      "type": ["string"],
      "abilities": ["string"],
      "height": number,
      "weight": number,
      "pokedex_id": integer
    }
    ```
    Example:
    ```json
    {
        "name": "Pikachu",
        "type": ["Electric"],
        "abilities": ["Static", "Lightning Rod"],
        "height": 0.4,
        "weight": 6.0,
        "pokedex_id": 25
    }
    ```
-   **Response**:
    -   **Status Code**: `201 Created` on success, `400 Bad Request` on invalid input, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the created Pokémon details.
    ```json
    {
      "id": "string",
      "name": "string",
      "type": ["string"],
      "abilities": ["string"],
      "height": number,
      "weight": number,
      "pokedex_id": integer
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "Pikachu",
        "type": ["Electric"],
        "abilities": ["Static", "Lightning Rod"],
        "height": 0.4,
        "weight": 6.0,
        "pokedex_id": 25
    }
    ```
-   **Usage**: Send a JSON payload with Pokémon details to add a new Pokémon.

#### Get All Pokémon

-   **Endpoint**: `/pokemons`
-   **Method**: `GET`
-   **Description**: Retrieves a list of all Pokémon from the database.
-   **Query Parameters**: None
-   **Response**:
    -   **Status Code**: `200 OK` on success, `500 Internal Server Error` on failure.
    -   **Body**: JSON array of Pokémon objects.
    ```json
    [
      {
        "id": "string",
        "name": "string",
        "type": ["string"],
        "abilities": ["string"],
        "height": number,
        "weight": number,
        "pokedex_id": integer
      }
    ]
    ```
    Example:
    ```json
    [
        {
            "id": "507f1f77bcf86cd799439011",
            "name": "Pikachu",
            "type": ["Electric"],
            "abilities": ["Static", "Lightning Rod"],
            "height": 0.4,
            "weight": 6.0,
            "pokedex_id": 25
        },
        {
            "id": "507f1f77bcf86cd799439012",
            "name": "Bulbasaur",
            "type": ["Grass", "Poison"],
            "abilities": ["Overgrow"],
            "height": 0.7,
            "weight": 6.9,
            "pokedex_id": 1
        }
    ]
    ```
-   **Usage**: Fetch all Pokémon to display in a list or table.

#### Get Pokémon by ID

-   **Endpoint**: `/pokemons/{id}`
-   **Method**: `GET`
-   **Description**: Retrieves a Pokémon by its ID.
-   **Path Parameters**:
    -   `id`: The unique identifier of the Pokémon (MongoDB ObjectID, e.g., `507f1f77bcf86cd799439011`).
-   **Response**:
    -   **Status Code**: `200 OK` on success, `404 Not Found` if the Pokémon does not exist, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the Pokémon details or error message.
    ```json
    {
      "id": "string",
      "name": "string",
      "type": ["string"],
      "abilities": ["string"],
      "height": number,
      "weight": number,
      "pokedex_id": integer
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "Pikachu",
        "type": ["Electric"],
        "abilities": ["Static", "Lightning Rod"],
        "height": 0.4,
        "weight": 6.0,
        "pokedex_id": 25
    }
    ```
-   **Usage**: Fetch a specific Pokémon to display its details.

#### Update Pokémon

-   **Endpoint**: `/pokemons/{id}`
-   **Method**: `PUT`
-   **Description**: Updates an existing Pokémon by its ID.
-   **Path Parameters**:
    -   `id`: The unique identifier of the Pokémon (MongoDB ObjectID).
-   **Request Body**:
    ```json
    {
      "name": "string",
      "type": ["string"],
      "abilities": ["string"],
      "height": number,
      "weight": number,
      "pokedex_id": integer
    }
    ```
    Example:
    ```json
    {
        "name": "Pikachu",
        "type": ["Electric"],
        "abilities": ["Static", "Lightning Rod"],
        "height": 0.5,
        "weight": 6.5,
        "pokedex_id": 25
    }
    ```
-   **Response**:
    -   **Status Code**: `200 OK` on success, `400 Bad Request` on invalid input, `404 Not Found` if the Pokémon does not exist, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the updated Pokémon details.
    ```json
    {
      "id": "string",
      "name": "string",
      "type": ["string"],
      "abilities": ["string"],
      "height": number,
      "weight": number,
      "pokedex_id": integer
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "Pikachu",
        "type": ["Electric"],
        "abilities": ["Static", "Lightning Rod"],
        "height": 0.5,
        "weight": 6.5,
        "pokedex_id": 25
    }
    ```
-   **Usage**: Send updated Pokémon data to modify an existing Pokémon.

#### Delete Pokémon

-   **Endpoint**: `/pokemons/{id}`
-   **Method**: `DELETE`
-   **Description**: Deletes a Pokémon by its ID.
-   **Path Parameters**:
    -   `id`: The unique identifier of the Pokémon (MongoDB ObjectID).
-   **Response**:
    -   **Status Code**: `204 No Content` on success, `404 Not Found` if the Pokémon does not exist, or `500 Internal Server Error` on failure.
    -   **Body**: Empty on success, or JSON error message on failure.
-   **Usage**: Remove a Pokémon from the database.

### User Endpoints

#### Create User

-   **Endpoint**: `/users`
-   **Method**: `POST`
-   **Description**: Creates a new user in the database.
-   **Request Body**:
    ```json
    {
        "name": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "name": "ashketchum",
        "password": "pikachu123"
    }
    ```
-   **Response**:
    -   **Status Code**: `201 Created` on success, `400 Bad Request` on invalid input, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the created user details.
    ```json
    {
        "id": "string",
        "name": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "ashketchum",
        "password": "pikachu123"
    }
    ```
-   **Usage**: Register a new user by sending their details.

#### Get User by ID

-   **Endpoint**: `/users/{id}`
-   **Method**: `GET`
-   **Description**: Retrieves a user by their ID.
-   **Path Parameters**:
    -   `id`: The unique identifier of the user (MongoDB ObjectID, e.g., `507f1f77bcf86cd799439011`).
-   **Response**:
    -   **Status Code**: `200 OK` on success, `404 Not Found` if the user does not exist, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the user details or error message.
    ```json
    {
        "id": "string",
        "name": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "ashketchum",
        "password": "pikachu123"
    }
    ```
-   **Usage**: Fetch a specific user’s details for profile display.

#### Get User by Username

-   **Endpoint**: `/users/username/{username}`
-   **Method**: `GET`
-   **Description**: Retrieves a user by their username.
-   **Path Parameters**:
    -   `username`: The unique username of the user (e.g., `ashketchum`).
-   **Response**:
    -   **Status Code**: `200 OK` on success, `404 Not Found` if the user does not exist, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the user details or error message.
    ```json
    {
        "id": "string",
        "name": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "ashketchum",
        "password": "pikachu123"
    }
    ```
-   **Usage**: Search for a user by their username.

### Authentication Endpoints

#### Login

-   **Endpoint**: `/login`
-   **Method**: `POST`
-   **Description**: Authenticates a user and returns a JWT token.
-   **Request Body**:
    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "username": "ashketchum",
        "password": "pikachu123"
    }
    ```
-   **Response**:
    -   **Status Code**: `200 OK` on success, `401 Unauthorized` on invalid credentials, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the authentication token.
    ```json
    {
        "token": "string"
    }
    ```
    Example:
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```
-   **Usage**: Log in a user to obtain a token for authenticated requests.

#### Register

-   **Endpoint**: `/register`
-   **Method**: `POST`
-   **Description**: Registers a new user and stores their credentials.
-   **Request Body**:
    ```json
    {
        "name": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "name": "misty",
        "password": "waterpulse"
    }
    ```
-   **Response**:
    -   **Status Code**: `201 Created` on success, `400 Bad Request` on invalid input, or `500 Internal Server Error` on failure.
    -   **Body**: JSON object containing the created user details.
    ```json
    {
        "id": "string",
        "name": "string",
        "password": "string"
    }
    ```
    Example:
    ```json
    {
        "id": "507f1f77bcf86cd799439012",
        "name": "misty",
        "password": "waterpulse"
    }
    ```
-   **Usage**: Register a new user with authentication credentials.

## Notes for Frontend Developers

-   **Base URL**: Ensure the frontend uses the correct host and port (default `:8081` unless overridden by `LISTEN_ADDR`).
-   **Authentication**: The `/login` endpoint provides a JWT token that should be included in the `Authorization` header (e.g., `Bearer <token>`) for protected endpoints, if implemented in `handlers`.
-   **Error Handling**: Handle HTTP status codes like `400`, `401`, `404`, and `500` appropriately in the frontend to display user-friendly error messages.
-   **Request Format**: All endpoints expect JSON payloads for `POST` and `PUT` requests. Ensure the `Content-Type` header is set to `application/json`.
-   **Environment Variables**: The backend relies on `MONGO_URI`, `MONGO_DB_NAME`, `RABBITMQ_URI`, and `LISTEN_ADDR`. Ensure these are configured correctly in the deployment environment.
-   **Pagination**: The `/pokemons` GET endpoint may return a large number of results. Consider adding query parameters like `limit` and `offset` in the future for pagination, if supported by `handlers.GetAllPokemons`.
-   **Security Note**: The `password` field is included in the `User` response for simplicity, but in a production environment, avoid returning sensitive fields like passwords in API responses.

## Example Usage (Frontend)

Here’s an example of how to call the `/pokemons` endpoint to create a Pokémon using JavaScript `fetch`:

```javascript
async function createPokemon(pokemon) {
    try {
        const response = await fetch("http://localhost:8081/pokemons", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer <your-jwt-token>", // If required
            },
            body: JSON.stringify(pokemon),
        });
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const data = await response.json();
        console.log("Pokémon created:", data);
    } catch (error) {
        console.error("Error creating Pokémon:", error);
    }
}

createPokemon({
    name: "Pikachu",
    type: ["Electric"],
    abilities: ["Static", "Lightning Rod"],
    height: 0.4,
    weight: 6.0,
    pokedex_id: 25,
});
```

This documentation should help frontend developers integrate with the REST API effectively.

### Instructions to Copy

1. **Copy the Text**: Highlight the entire markdown text above (from `# REST API Endpoints` to the end) and copy it to your clipboard (Ctrl+C or Cmd+C).
2. **Save as Markdown**: Paste the text into a text editor (e.g., VS Code, Notepad) and save it with a `.md` extension (e.g., `api_endpoints.md`).
3. **Use in Frontend**: Share the markdown file with your frontend team or render it in a markdown viewer (e.g., GitHub, VS Code, or a documentation tool like MkDocs).

### Additional Notes

-   **Why You Couldn't Copy Before**: The previous response was embedded in an artifact tag, which might have been rendered in a way that made copying difficult (e.g., due to UI constraints or formatting). Providing the markdown as plain text avoids this issue.
-   **Content Details**: The markdown is unchanged from the previous version, as it already incorporated the correct `Pokemon`, `User`, `LoginRequest`, and `LoginResponse` structs. It includes precise request/response schemas, example JSON, and frontend guidance.
-   **Next Steps**: If you need the markdown in a specific format (e.g., hosted on a GitHub repo, converted to HTML, or integrated into a larger documentation system), let me know, and I can assist further. Alternatively, if you want to test the API endpoints, I can provide a sample client script or curl commands.
