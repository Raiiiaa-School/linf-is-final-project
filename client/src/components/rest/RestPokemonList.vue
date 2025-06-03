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

</script>

<style>
/* Estilos específicos para a lista se necessário */
.v-pagination__list button {
    color: black;
}
</style>
