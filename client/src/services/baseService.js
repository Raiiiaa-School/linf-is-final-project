/* eslint-disable no-unused-vars */

export class BaseService {
    baseURL = "http://localhost";

    constructor(port, path) {
        this.port = port;
        this.path = path;
        this.baseURL = `${this.baseURL}:${port}`;
        if (path) {
            this.baseURL += `/${path}`;
        }
    }

    async listPokemons() {
        throw new Error("Method 'listPokemons' not implemented.");
    }
    async getPokemons(_id) {
        throw new Error("Method 'getPokemons' not implemented.");
    }
    async createPokemons(_data) {
        throw new Error("Method 'createPokemons' not implemented.");
    }
    async updatePokemons(_id, _data) {
        throw new Error("Method 'updatePokemons' not implemented.");
    }
    async deletePokemons(_id) {
        throw new Error("Method 'deletePokemons' not implemented.");
    }

    async listUsers() {
        throw new Error("Method 'listUsers' not implemented.");
    }
    async getUsers(_id) {
        throw new Error("Method 'getUsers' not implemented.");
    }
    async createUsers(_data) {
        throw new Error("Method 'createUsers' not implemented.");
    }
    async updateUsers(_id, _data) {
        throw new Error("Method 'updateUsers' not implemented.");
    }
    async deleteUsers(_id) {
        throw new Error("Method 'deleteUsers' not implemented.");
    }
}
