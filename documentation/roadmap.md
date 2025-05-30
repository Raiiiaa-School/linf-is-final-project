# Project Roadmap

This roadmap provides a step-by-step guideline for completing the Pokémon and User management system, covering both backend (services) and frontend (client) development. The backend is implemented in Go with REST, GraphQL, gRPC, and SOAP services, and the frontend will be built in Vue.js to provide a user-friendly interface for testing all endpoints.

---

## 1. Project Setup

-   **Version Control:** Ensure all code is tracked in Git.
-   **Directory Structure:** Keep backend services, frontend client, and documentation organized in separate folders.
-   **Environment Variables:** Use `.env` files or Docker Compose for configuration (MongoDB, RabbitMQ, JWT secrets, etc.).

---

## 2. Backend Services

### 2.1. REST API (Already Implemented)

-   Pokémon CRUD endpoints (`/pokemons`): create, read (all/by id), update, delete.
-   User endpoints: create, get by id, get by username.
-   Authentication: login and register endpoints with JWT support.
-   Health check endpoint.
-   **To Do:**
    -   Remove password from user responses for security.
    -   Add pagination to `/pokemons` GET if needed.
    -   Add/verify authentication middleware for protected endpoints.
    -   Write tests for handlers and integration.

### 2.2. GraphQL Service

-   Define schema for Pokémon and User.
-   Implement queries and mutations for CRUD operations.
-   Ensure authentication and error handling.

### 2.3. gRPC Service

-   Define protobuf files for Pokémon and User services.
-   Implement server and handlers for CRUD operations.
-   Ensure compatibility with other services.

### 2.4. SOAP Service

-   Implement SOAP endpoints for Pokémon operations.
-   Provide WSDL documentation.

### 2.5. Messaging & Database

-   Ensure MongoDB models and indexes are set up.
-   Use RabbitMQ for asynchronous tasks (e.g., notifications, logging).
-   Document message formats and queues.

---

## 3. Frontend Client (Vue.js)

### 3.1. Project Setup

-   Initialize Vue project (Vite or Vue CLI).
-   Install dependencies: Axios, Vue Router, Pinia/Vuex, UI library (e.g., Vuetify).

### 3.2. Authentication

-   Create login and registration pages.
-   Store JWT securely (preferably in memory or HttpOnly cookies).
-   Add route guards for protected pages.

### 3.3. Pokémon Management

-   List all Pokémon (with pagination if available).
-   View Pokémon details.
-   Create, edit, and delete Pokémon (forms and actions).
-   Add client-side validation.

### 3.4. User Management

-   Profile page for user details.
-   Search user by username.

### 3.5. API Integration

-   Abstract API calls in a service module.
-   Handle errors and show user-friendly messages.
-   Show loading indicators during requests.

### 3.6. UI/UX

-   Ensure responsive design for desktop and mobile.
-   Use a consistent theme and accessible components.

---

## 4. Testing

-   Backend: Unit and integration tests for all endpoints.
-   Frontend: Component and end-to-end tests (e.g., Cypress).

---

## 5. Documentation

-   Keep `rest-api.md` and other docs up to date.
-   Write a user guide for the frontend app.
-   Document deployment steps (Docker, environment variables, etc.).

---

## 6. Deployment

-   Dockerize all services and the frontend.
-   Use Docker Compose for local development and testing.
-   Set up CI/CD for automated testing and deployment.

---

## 7. Final Review

-   Ensure all assignment requirements are met.
-   Review security (no passwords in responses, JWT validation, etc.).
-   Polish UI and fix any remaining bugs.

---

**Tip:** Use the Vue.js frontend as the main tool for testing and demonstrating all API endpoints, providing a better experience than Postman or curl.
