import axios from "axios";
import { BaseService } from "./baseService";

export class RestService extends BaseService {
    async listPokemons() {
        const response = await axios.get(`${this.baseURL}/pokemons`);
        return response.data;
    }

    async getPokemons(id) {
        const response = await axios.get(`${this.baseURL}/pokemons/${id}`);
        return response.data;
    }

    async createPokemon(data) {
        const response = await axios.post(`${this.baseURL}/pokemons`, data);
        return response.data;
    }

    async updatePokemons(id, data) {
        const response = await axios.put(
            `${this.baseURL}/pokemons/${id}`,
            data
        );
        return response.data;
    }

    async deletePokemons(id) {
        const response = await axios.delete(`${this.baseURL}/pokemons/${id}`);
        return response.data;
    }

    async listUsers() {
        const response = await axios.get(`${this.baseURL}/users`);
        return response.data;
    }

    async getUsers(id) {
        const response = await axios.get(`${this.baseURL}/users/${id}`);
        return response.data;
    }

    async createUsers(data) {
        const response = await axios.post(`${this.baseURL}/users`, data);
        return response.data;
    }
}
