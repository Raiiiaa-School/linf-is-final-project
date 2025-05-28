import { useEffect, useState } from "react";
import "./App.css";
import axios from "axios";

const REST_API_URL =
    import.meta.env.VITE_REST_API_URL || "http://localhost:8081";

interface Pokemon {
    id: string;
    name: string;
    pokedex_id: number;
    type: string[];
    abilities: string[];
    height: number;
    weight: number;
}

function App() {
    const [pokemons, setPokemons] = useState<Pokemon[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<Error | null>(null);

    useEffect(() => {
        const fetchPokemons = async () => {
            try {
                setLoading(true);
                const response = await axios.get(`${REST_API_URL}/pokemons`);
                setPokemons(response.data);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err);
                }
            } finally {
                setLoading(false);
            }
        };

        fetchPokemons();
    }, []);

    if (loading) {
        return <div>Loading Pokemons...</div>;
    }

    if (error) {
        return <div>Error loading Pokemons: {error.message}</div>;
    }

    return (
        <div className="App">
            <h1>Pokedex (REST)</h1>
            <ul>
                {pokemons.map((pokemon) => (
                    <li key={pokemon.id}>
                        <h2>{pokemon.name}</h2>
                        <p>Pokedex ID: {pokemon.pokedex_id}</p>
                        <p>Type(s): {pokemon.type.join(", ")}</p>
                        <p>Abilities: {pokemon.abilities.join(", ")}</p>
                        <p>Height: {pokemon.height} m</p>
                        <p>Weight: {pokemon.weight} kg</p>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;
