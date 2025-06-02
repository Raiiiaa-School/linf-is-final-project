<template>
    <v-card class="pa-4 mt-4" elevation="2">
        <v-card-title>
            <span v-if="mode === 'create'">Criar Novo Pokémon</span>
            <span v-else-if="mode === 'update'">Atualizar Pokémon</span>
            <span v-else-if="mode === 'delete'">Eliminar Pokémon</span>
        </v-card-title>
        <v-card-text>
            <v-form @submit.prevent="handleSubmit">
                <v-text-field
                    v-model="pokemonData.name"
                    label="Nome do Pokémon"
                    :rules="[(v) => !!v || 'Nome é obrigatório']"
                    required
                    :disabled="mode === 'delete'"
                ></v-text-field>
                <v-text-field
                    v-model.number="pokemonData.pokedex_id"
                    label="ID da Pokedex"
                    type="number"
                    :rules="[(v) => v >= 0 || 'ID deve ser positivo']"
                    :disabled="mode === 'delete'"
                ></v-text-field>
                <v-text-field
                    v-model.number="pokemonData.height"
                    label="Altura (m)"
                    type="number"
                    step="0.01"
                    :rules="[(v) => v >= 0 || 'Altura deve ser positiva']"
                    :disabled="mode === 'delete'"
                ></v-text-field>
                <v-text-field
                    v-model.number="pokemonData.weight"
                    label="Peso (kg)"
                    type="number"
                    step="0.01"
                    :rules="[(v) => v >= 0 || 'Peso deve ser positivo']"
                    :disabled="mode === 'delete'"
                ></v-text-field>
                <v-combobox
                    v-model="pokemonData.type"
                    label="Tipos (separados por vírgula)"
                    multiple
                    chips
                    clearable
                    :disabled="mode === 'delete'"
                ></v-combobox>
                <v-combobox
                    v-model="pokemonData.abilities"
                    label="Habilidades (separadas por vírgula)"
                    multiple
                    chips
                    clearable
                    :disabled="mode === 'delete'"
                ></v-combobox>

                <v-text-field
                    v-if="mode === 'update' || mode === 'delete'"
                    v-model="pokemonId"
                    label="ID do Pokémon para Atualizar/Eliminar"
                    :rules="[(v) => !!v || 'ID é obrigatório para esta ação']"
                    required
                ></v-text-field>

                <v-btn type="submit" color="primary" class="mt-4">
                    <span v-if="mode === 'create'">Criar Pokémon</span>
                    <span v-else-if="mode === 'update'">Atualizar Pokémon</span>
                    <span v-else-if="mode === 'delete'">Eliminar Pokémon</span>
                </v-btn>
            </v-form>
        </v-card-text>
        <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">{{
            snackbarText
        }}</v-snackbar>
    </v-card>
</template>

<script setup>
import axios from "axios";
import { onMounted, defineProps, reactive, ref } from "vue";

const props = defineProps({
    mode: {
        type: String,
        default: "create", // 'create', 'update', 'delete'
        validator: (value) => ["create", "update", "delete"].includes(value),
    },
    initialPokemon: {
        type: Object,
        default: () => ({}),
    },
});

const pokemonData = reactive({
    name: "",
    pokedex_id: null,
    height: null,
    weight: null,
    type: [],
    abilities: [],
});

const pokemonId = props.initialPokemon._id || ""; // Para operações de update/delete, assumindo que o ID é _id do MongoDB

const snackbar = ref(false);
const snackbarText = ref("");
const snackbarColor = ref("");

onMounted(() => {
    if (
        (props.mode === "update" || props.mode === "delete") &&
        props.initialPokemon !== JSON.stringify({})
    ) {
        Object.assign(pokemonData, props.initialPokemon);
    }
});

function handleSubmit() {
    if (props.mode === "create") {
        createPokemon();
    } else if (props.mode === "update") {
        updatePokemon();
    } else if (props.mode === "delete") {
        deletePokemon();
    }
}

async function createPokemon() {
    try {
        const response = await axios.post(
            process.env.VUE_APP_REST_SERVICE_URL + "/pokemons",
            pokemonData
        );
        console.log("Pokémon criado:", response);
    } catch (error) {
        showSnackbar(
            `Erro ao criar Pokémon: ${error.message || error}`,
            "error"
        );
        return;
    }

    showSnackbar("Pokémon criado com sucesso!", "success");
}

function updatePokemon() {
    if (!pokemonId) {
        showSnackbar("ID do Pokémon é obrigatório para atualização.", "error");
        return;
    }
    showSnackbar("Pokémon atualizado com sucesso!", "success");
}

function deletePokemon() {
    if (!pokemonId) {
        showSnackbar("ID do Pokémon é obrigatório para eliminação.", "error");
        return;
    }
    showSnackbar("Pokémon eliminado com sucesso!", "success");
}

function showSnackbar(message, color) {
    snackbarText.value = message;
    snackbarColor.value = color;
    snackbar.value = true;
}
</script>

<style scoped>
/* Estilos específicos para o formulário se necessário */
</style>
