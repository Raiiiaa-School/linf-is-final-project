import { graphqlClient } from "@/main";
import { BaseService } from "./baseService";
import { gql } from "urql";

export class GraphQLService extends BaseService {
    constructor(port = 8082, path = "/query") {
        super(port, path);
        this.client = graphqlClient;
    }

    async listPokemons() {
        const query = gql`
            query GetPokemons {
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
        `;

        const result = await this.client.query(query).toPromise();

        if (result.error) {
            console.error("GraphQL Error (listPokemons):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to fetch pokemons"
            );
        }
        return result.data.pokemons;
    }

    async getPokemon(id) {
        const query = gql`
            query GetPokemonById($id: ID!) {
                pokemon(id: $id) {
                    id
                    name
                    type
                    abilities
                    height
                    weight
                    pokedexId
                }
            }
        `;

        const result = await this.client.query(query, { id }).toPromise();

        if (result.error) {
            console.error("GraphQL Error (getPokemon):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to fetch pokemon"
            );
        }
        return result.data.pokemon;
    }

    async createPokemon(input) {
        const mutation = gql`
            mutation CreatePokemon($input: CreatePokemonInput!) {
                createPokemon(input: $input) {
                    id
                    name
                    type
                    abilities
                    height
                    weight
                    pokedexId
                }
            }
        `;

        const result = await this.client
            .mutation(mutation, { input })
            .toPromise();

        if (result.error) {
            console.error("GraphQL Error (createPokemon):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to create pokemon"
            );
        }
        return result.data.createPokemon;
    }

    async updatePokemon(id, input) {
        const mutation = gql`
            mutation UpdatePokemon($id: ID!, $input: UpdatePokemonInput!) {
                updatePokemon(id: $id, input: $input) {
                    id
                    name
                    type
                    abilities
                    height
                    weight
                    pokedexId
                }
            }
        `;

        const result = await this.client
            .mutation(mutation, { id, input })
            .toPromise();

        if (result.error) {
            console.error("GraphQL Error (updatePokemon):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to update pokemon"
            );
        }
        return result.data.updatePokemon;
    }

    async deletePokemon(id) {
        const mutation = gql`
            mutation DeletePokemon($id: ID!) {
                deletePokemon(id: $id)
            }
        `;

        const result = await this.client.mutation(mutation, { id }).toPromise();

        if (result.error) {
            console.error("GraphQL Error (deletePokemon):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to delete pokemon"
            );
        }
        // The deletePokemon mutation returns a boolean (or ID)
        // Check if the operation was successful based on your backend's return.
        // If it returns true on success:
        return result.data.deletePokemon;
    }

    // --- User Operations ---

    async listUsers() {
        // Your docs don't show a 'listUsers' query, only 'user(id)' and 'userByUsername)'.
        // Assuming you meant 'pokemons' earlier, but if you have a listUsers,
        // you'd define its GraphQL query here. For now, I'll return an empty array
        // or throw an error to reflect that the GraphQL API doesn't support this based on your docs.
        console.warn(
            "GraphQL API documentation does not show a 'listUsers' query."
        );
        return [];
    }

    async getUser(id) {
        const query = gql`
            query GetUserById($id: ID!) {
                user(id: $id) {
                    id
                    username
                }
            }
        `;

        const result = await this.client.query(query, { id }).toPromise();

        if (result.error) {
            console.error("GraphQL Error (getUser):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to fetch user"
            );
        }
        return result.data.user;
    }

    async getUserByUsername(username) {
        const query = gql`
            query GetUserByUsername($username: String!) {
                userByUsername(username: $username) {
                    id
                    username
                }
            }
        `;

        const result = await this.client.query(query, { username }).toPromise();

        if (result.error) {
            console.error("GraphQL Error (getUserByUsername):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to fetch user by username"
            );
        }
        return result.data.userByUsername;
    }

    async createUser(input) {
        const mutation = gql`
            mutation CreateUser($input: CreateUserInput!) {
                createUser(input: $input) {
                    id
                    username
                }
            }
        `;

        const result = await this.client
            .mutation(mutation, { input })
            .toPromise();

        if (result.error) {
            console.error("GraphQL Error (createUser):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Failed to create user"
            );
        }
        return result.data.createUser;
    }

    async login(input) {
        const mutation = gql`
            mutation Login($input: LoginInput!) {
                login(input: $input) {
                    token
                }
            }
        `;

        const result = await this.client
            .mutation(mutation, { input })
            .toPromise();

        if (result.error) {
            console.error("GraphQL Error (login):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Login failed"
            );
        }
        // Store the token and return it
        const token = result.data.login.token;
        localStorage.setItem("authToken", token); // Store for future requests
        return token;
    }

    async register(input) {
        const mutation = gql`
            mutation Register($input: RegisterInput!) {
                register(input: $input) {
                    id
                    username
                }
            }
        `;

        const result = await this.client
            .mutation(mutation, { input })
            .toPromise();

        if (result.error) {
            console.error("GraphQL Error (register):", result.error);
            throw new Error(
                result.error.graphQLErrors?.[0]?.message ||
                    result.error.networkError?.message ||
                    "Registration failed"
            );
        }
        return result.data.register;
    }
}
