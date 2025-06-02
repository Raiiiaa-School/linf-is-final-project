<template>
    <v-card class="pa-4 mt-4" elevation="2">
        <v-card-title>Lista de Pokémon</v-card-title>
        <v-card-text>
            <div v-if="loading">A carregar Pokémon...</div>
            <div v-else-if="error">{{ error }}</div>
            <v-data-table v-else :items="pokemons" item-value="id">
                <template v-slot:[itemTypeSlotName]="{ item }">
                    {{ item.type ? item.type.join(", ") : "" }}
                </template>
                <template v-slot:[itemAbilitiesSlotName]="{ item }">
                    {{ item.abilities ? item.abilities.join(", ") : "" }}
                </template>
            </v-data-table>
            <div v-if="!pokemons.length && !loading && !error">
                Nenhum Pokémon encontrado.
            </div>
        </v-card-text>
    </v-card>
</template>

<script setup>
import { gql, useQuery } from "@urql/vue";

import { computed } from "vue";

const GetPokemons = gql`
    query GetAllPokemons {
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

const { data, fetching, error } = useQuery({
    query: GetPokemons,
});

const pokemons = computed(() => data.value?.pokemons || []);
const loading = computed(() => fetching.value);
</script>

<!-- const GET_POSTS = gql`
  query GetPosts {
    posts {
      id
      title
      content
    }
  }
`;

// `useQuery` returns a reactive result object
const { data, fetching, error } = useQuery({
  query: GET_POSTS,
}); -->
