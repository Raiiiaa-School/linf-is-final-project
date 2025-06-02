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
import axios from "axios";
import { onMounted, ref } from "vue";

const loading = ref(true);
const error = ref(null);
const pokemons = ref([]);

// const headers = [
//     { title: "ID", key: "id" },
//     { title: "Nome", key: "name" },
//     { title: "Pokedex ID", key: "pokedex_id" },
//     { title: "Altura (m)", key: "height" },
//     { title: "Peso (kg)", key: "weight" },
//     { title: "Tipos", key: "type" },
//     { title: "Habilidades", key: "abilities" },
// ];

onMounted(async () => {
    loading.value = true;
    error.value = null;
    try {
        const response = await axios.get(
            process.env.VUE_APP_REST_SERVICE_URL + "/pokemons"
        );
        if (response.status === 200) {
            pokemons.value = response.data;
        }
        loading.value = false;
    } catch (error) {
        error.value = `Erro ao carregar Pokémon: ${error.message || error}`;
        console.error("Erro ao buscar pokémons:", error);
    }
});

// export default {
//   data() {
//     return {
//       pokemons: [],
//       loading: false,
//       error: null,
//       headers: [
//         { title: 'ID', key: 'id' },
//         { title: 'Nome', key: 'name' },
//         { title: 'Pokedex ID', key: 'pokedex_id' },
//         { title: 'Altura (m)', key: 'height' },
//         { title: 'Peso (kg)', key: 'weight' },
//         { title: 'Tipos', key: 'type' },
//         { title: 'Habilidades', key: 'abilities' },
//       ],
//     };
//   },
//   computed: {
//     itemTypeSlotName() {
//       return 'item.type';
//     },
//     itemAbilitiesSlotName() {
//       return 'item.abilities';
//     }
//   },
//   async mounted() {
//     await this.fetchPokemons();
//   },
//   methods: {
//     async fetchPokemons() {
//       this.loading = true;
//       this.error = null;
//       try {
//         const baseUrl = import.meta.env.VITE_REST_API_URL;
//         const response = await fetch(`${baseUrl}/pokemons`);
//         if (!response.ok) {
//           const errorData = await response.json();
//           throw new Error(`Erro: ${response.statusText} - ${errorData.message || 'Erro desconhecido'}`);
//         }
//         this.pokemons = await response.json();
//       } catch (error) {
//         console.error("Erro ao buscar pokémons:", error);
//         this.error = `Erro ao carregar Pokémon: ${error.message || error}`;
//       } finally {
//         this.loading = false;
//       }
//     }
//   }
// };
</script>

<style>
/* Estilos específicos para a lista se necessário */
.v-pagination__list button {
    color: black;
}
</style>
