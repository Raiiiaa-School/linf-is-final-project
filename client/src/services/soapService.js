import { BaseService } from "./baseService";
import axios from 'axios'; // Certifique-se de importar o axios

export class SoapService extends BaseService {
    async listPokemons(data) {
        // Ensure data has the necessary properties as arrays
        const processedData = {
            name: data?.name || '',
            types: Array.isArray(data?.types) ? data.types : [],
            abilities: Array.isArray(data?.abilities) ? data.abilities : [],
            height: data?.height || '',
            weight: data?.weight || '',
            pokedex_id: data?.pokedex_id || ''
        };

        const xmlStringList = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <PokemonRequest>
                <action>list</action>
                <pokemon>
                    <name>${processedData.name}</name>
                    ${processedData.types.map((type) => {
                        return `<type>${type}</type>`
                    }).join('')}
                    ${processedData.abilities.map((ability) => {
                        return `<abilities>${ability}</abilities>`
                    }).join('')}
                    <height>${processedData.height}</height>
                    <weight>${processedData.weight}</weight>
                    <pokedex_id>${processedData.pokedex_id}</pokedex_id>
                </pokemon>
                </PokemonRequest>
            </soap:Body>
            </soap:Envelope>
        `;

        const response = await axios.post(`${this.baseURL}/soap/pokemon`, xmlStringList);
        return response.data;
    }

    async getPokemon(id) {
        const pokemonId = id?.id || '';
        const xmlStringGetPokemon = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <PokemonRequest>
                <action>get</action>
                <pokemon>
                    <id>${pokemonId}</id>
                </pokemon>
                </PokemonRequest>
            </soap:Body>
            </soap:Envelope>
        `;

        const response = await axios.post(`${this.baseURL}/soap/pokemon`, xmlStringGetPokemon);
        return response.data;
    }

    async createPokemon(data) {
        const processedData = {
            name: data?.name || '',
            types: Array.isArray(data?.types) ? data.types : [],
            abilities: Array.isArray(data?.abilities) ? data.abilities : [],
            height: data?.height || '',
            weight: data?.weight || '',
            pokedex_id: data?.pokedex_id || ''
        };

        const xmlStringCreate = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <PokemonRequest>
                <action>create</action>
                <pokemon>
                    <name>${processedData.name}</name>
                    ${processedData.types.map((type) => {
                        return `<type>${type}</type>`
                    }).join('')}
                    ${processedData.abilities.map((ability) => {
                        return `<abilities>${ability}</abilities>`
                    }).join('')}
                    <height>${processedData.height}</height>
                    <weight>${processedData.weight}</weight>
                    <pokedex_id>${processedData.pokedex_id}</pokedex_id>
                </pokemon>
                </PokemonRequest>
            </soap:Body>
            </soap:Envelope>
        `;

        const response = await axios.post(`${this.baseURL}/soap/pokemon`, xmlStringCreate);
        return response.data;
    }

    async updatePokemon(id, data) {
        const processedData = {
            id: id?.id || '',
            name: data?.name || '',
            types: Array.isArray(data?.types) ? data.types : [],
            abilities: Array.isArray(data?.abilities) ? data.abilities : [],
            height: data?.height || '',
            weight: data?.weight || '',
            pokedex_id: data?.pokedex_id || ''
        };

        const xmlStringUpdate = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <PokemonRequest>
                <action>update</action>
                <pokemon>
                    <id>${processedData.id}</id>
                    <name>${processedData.name}</name>
                    ${processedData.types.map((type) => {
                        return `<type>${type}</type>`
                    }).join('')}
                    ${processedData.abilities.map((ability) => {
                        return `<abilities>${ability}</abilities>`
                    }).join('')}
                    <height>${processedData.height}</height>
                    <weight>${processedData.weight}</weight>
                    <pokedex_id>${processedData.pokedex_id}</pokedex_id>
                </pokemon>
                </PokemonRequest>
            </soap:Body>
            </soap:Envelope>
        `;

        const response = await axios.post(`${this.baseURL}/soap/pokemon`, xmlStringUpdate);
        return response.data;
    }

    async deletePokemon(id) {
        const pokemonId = id?.id || '';
        const xmlStringDelete = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <PokemonRequest>
                <action>delete</action>
                <pokemon>
                    <id>${pokemonId}</id>
                </pokemon>
                </PokemonRequest>
            </soap:Body>
            </soap:Envelope>
        `;

        const response = await axios.post(`${this.baseURL}/soap/pokemon`, xmlStringDelete);
        return response.data;
    }
    async listUsers(data) {
        const processedData = {
            name: data?.name || '',
            email: data?.email || ''
            // Adicione outras propriedades de filtro se necessário
        };

        const xmlStringListUsers = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <UserRequest>
                <action>list</action>
                <user>
                    <name>${processedData.name}</name>
                    <email>${processedData.email}</email>
                </user>
                </UserRequest>
            </soap:Body>
            </soap:Envelope>
        `;
        const response = await axios.post(`${this.baseURL}/soap/user`, xmlStringListUsers);
        return response.data;
    }

    async getUser(id) {
        const userId = id?.id || '';
        const xmlStringGetUser = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <UserRequest>
                <action>get</action>
                <user>
                    <id>${userId}</id>
                </user>
                </UserRequest>
            </soap:Body>
            </soap:Envelope>
        `;
        const response = await axios.post(`${this.baseURL}/soap/user`, xmlStringGetUser);
        return response.data;
    }

    async createUser(data) {
        const processedData = {
            name: data?.name || '',
            email: data?.email || '',
            password: data?.password || '' // Assumindo que a criação inclui senha
        };

        const xmlStringCreateUser = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <UserRequest>
                <action>create</action>
                <user>
                    <name>${processedData.name}</name>
                    <email>${processedData.email}</email>
                    <password>${processedData.password}</password>
                </user>
                </UserRequest>
            </soap:Body>
            </soap:Envelope>
        `;
        const response = await axios.post(`${this.baseURL}/soap/user`, xmlStringCreateUser);
        return response.data;
    }

    async updateUser(id, data) {
        const processedData = {
            id: id?.id || '',
            name: data?.name || '',
            email: data?.email || '',
            password: data?.password || '' // Incluir senha se puder ser atualizada
        };

        const xmlStringUpdateUser = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <UserRequest>
                <action>update</action>
                <user>
                    <id>${processedData.id}</id>
                    <name>${processedData.name}</name>
                    <email>${processedData.email}</email>
                    <password>${processedData.password}</password>
                </user>
                </UserRequest>
            </soap:Body>
            </soap:Envelope>
        `;
        const response = await axios.post(`${this.baseURL}/soap/user`, xmlStringUpdateUser);
        return response.data;
    }

    async deleteUser(id) {
        const userId = id?.id || '';
        const xmlStringDeleteUser = `
            <?xml version="1.0" encoding="UTF-8"?>
            <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
            <soap:Body>
                <UserRequest>
                <action>delete</action>
                <user>
                    <id>${userId}</id>
                </user>
                </UserRequest>
            </soap:Body>
            </soap:Envelope>
        `;
        const response = await axios.post(`${this.baseURL}/soap/user`, xmlStringDeleteUser);
        return response.data;
    }
}