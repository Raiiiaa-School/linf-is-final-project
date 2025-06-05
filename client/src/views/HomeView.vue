<template>
    <div class="flex">
        <v-select
            v-model="selectedService"
            label="Select"
            :items="SERVICES_OPTIONS"
        ></v-select>
        <v-select
            v-model="selectedObject"
            label="Select"
            :items="OBJECTS_OPTIONS"
        ></v-select>
    </div>
    <v-data-table
        v-if="selectedObject === OBJECTS_OPTIONS[0]"
        ref="data-table"
        v-model="selectedItems"
        show-select
        :headers="POKEMONS_HEADERS"
        :items="items"
        @update:current-items="loadPokemons"
    >
        <template v-slot:top>
            <v-toolbar flat>
                <v-toolbar-title> Pokemons </v-toolbar-title>

                <v-btn
                    rounded="lg"
                    text="Create"
                    border
                    @click="createPokemon"
                ></v-btn>
            </v-toolbar>
        </template>

        <template v-slot:item.actions="{ item }">
            <div class="d-flex ga-2 justify-end">
                <v-btn
                    color="primary"
                    text="Edit"
                    @click="editPokemon(item.id)"
                >
                </v-btn>
                <v-btn
                    color="primary"
                    text="Delete"
                    @click="deletePokemon(item.id)"
                >
                </v-btn>
            </div>
        </template>
    </v-data-table>

    <v-data-table
        v-if="selectedObject === OBJECTS_OPTIONS[1]"
        v-model="selectedItems"
        show-select
        :headers="USERS_HEADERS"
        :items="items"
    >
        <template v-slot:top>
            <v-toolbar flat>
                <v-toolbar-title> Pokemons </v-toolbar-title>

                <v-btn
                    rounded="lg"
                    text="Create"
                    border
                    @click="createUser"
                ></v-btn>
            </v-toolbar>
        </template>
    </v-data-table>

    <v-dialog v-model="pokemonDialog" max-width="1000">
        <v-card
            :subtitle="`${isEditing ? 'Update' : 'Create'} your favorite book`"
            :title="`${isEditing ? 'Edit' : 'Add'} a Book`"
        >
            <template v-slot:text>
                <v-row>
                    <v-col cols="12">
                        <v-text-field
                            v-model="record.name"
                            label="Name"
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-combobox
                            multiple
                            chips
                            clearable
                            :items="[
                                'Fire',
                                'Water',
                                'Grass',
                                'Electric',
                                'Psychic',
                                'Rock',
                                'Ground',
                            ]"
                            v-model="record.type"
                            label="Types"
                        ></v-combobox>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-combobox
                            multiple
                            chips
                            v-model="record.abilities"
                            :items="[
                                'Overgrow',
                                'Blaze',
                                'Torrent',
                                'Shield Dust',
                                'Levitate',
                                'Static',
                                'Flame Body',
                            ]"
                            label="Abilities"
                        ></v-combobox>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-number-input
                            v-model="record.height"
                            :min="1"
                            label="Height (cm)"
                        ></v-number-input>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-number-input
                            v-model="record.weight"
                            :min="1"
                            label="Weight (kg)"
                        ></v-number-input>
                    </v-col>
                </v-row>
            </template>

            <v-divider></v-divider>

            <v-card-actions class="bg-surface-light">
                <v-btn
                    text="Cancel"
                    variant="plain"
                    @click="dialog = false"
                ></v-btn>

                <v-spacer></v-spacer>

                <v-btn text="Save" @click="savePokemon"></v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>

    <v-dialog v-model="userDialog" max-width="500">
        <v-card
            :subtitle="`${isEditing ? 'Update' : 'Create'} your favorite book`"
            :title="`${isEditing ? 'Edit' : 'Add'} a Book`"
        >
            <template v-slot:text>
                <v-row>
                    <v-col cols="12">
                        <v-text-field
                            v-model="record.title"
                            label="Title"
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-text-field
                            v-model="record.author"
                            label="Author"
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-select
                            v-model="record.genre"
                            :items="[
                                'Fiction',
                                'Dystopian',
                                'Non-Fiction',
                                'Sci-Fi',
                            ]"
                            label="Genre"
                        ></v-select>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-number-input
                            v-model="record.year"
                            :max="adapter.getYear(adapter.date())"
                            :min="1"
                            label="Year"
                        ></v-number-input>
                    </v-col>

                    <v-col cols="12" md="6">
                        <v-number-input
                            v-model="record.pages"
                            :min="1"
                            label="Pages"
                        ></v-number-input>
                    </v-col>
                </v-row>
            </template>

            <v-divider></v-divider>

            <v-card-actions class="bg-surface-light">
                <v-btn
                    text="Cancel"
                    variant="plain"
                    @click="dialog = false"
                ></v-btn>

                <v-spacer></v-spacer>

                <v-btn text="Save" @click="save"></v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup>
import { useQueryParams } from "@/components/useQueryParams";
import { RestService } from "@/services/restService";
import { useMessagesStore } from "@/stores/messages";
import { onMounted, ref, shallowRef, watch } from "vue";
import { templateRef } from "vuetify/lib/util";

const { getQueryParam, setQueryParams } = useQueryParams();
const messages = useMessagesStore();

const options = ref({
    page: 1,
    itemsPerPage: 10,
});

const SERVICES_OPTIONS = ["Rest", "GraphQL", "gRPC", "Soap"];
const OBJECTS_OPTIONS = ["Pokemons", "Users"];

const POKEMONS_HEADERS = [
    { title: "ID", value: "id" },
    { title: "Name", value: "name" },
    { title: "Type", value: "type" },
    { title: "Abilities", value: "abilities" },
    { title: "Height", value: "height" },
    { title: "Weight", value: "weight" },
    { title: "Pokedex ID", value: "pokedex_id" },
    { title: "Actions", key: "actions", align: "end", sortable: false },
];

const USERS_HEADERS = [
    { title: "ID", value: "id" },
    { title: "Username", value: "name" },
];

const DEFAULT_RECORD = {
    id: null,
    name: "",
    type: [],
    abilities: [],
    height: 0,
    weight: 0,
    pokedex_id: 0,
};

const record = ref(DEFAULT_RECORD);

const dataTableRef = templateRef("data-table");

const pokemonDialog = shallowRef(false);
const userDialog = shallowRef(false);
const isEditing = ref(false);

const selectedService = ref("");
const selectedObject = ref(OBJECTS_OPTIONS[0]);

const items = ref([]);
const selectedItems = ref([]);

const currentService = ref(null);

onMounted(() => {
    const service = getQueryParam("service");
    if (!service) {
        selectedService.value = SERVICES_OPTIONS[0];
    } else {
        selectedService.value = service;
    }
});

watch(selectedService, (newService) => {
    setQueryParams({ service: newService }, true);

    switch (selectedService.value) {
        case "Rest":
            currentService.value = new RestService(8081);
            break;
        case "GraphQL":
            selectedObject.value = OBJECTS_OPTIONS[1];
            break;
        case "gRPC":
            selectedObject.value = OBJECTS_OPTIONS[0];
            break;
        case "Soap":
            selectedObject.value = OBJECTS_OPTIONS[1];
            break;
        default:
            selectedObject.value = OBJECTS_OPTIONS[0];
    }
});

watch(currentService, async (newService) => {
    if (!newService) return;
    loadPokemons();
});

async function loadPokemons() {
    if (!currentService.value) {
        return;
    }
    const resposnse = await currentService.value.listPokemons();
    items.value = resposnse;
}

function createPokemon() {
    isEditing.value = false;
    record.value = DEFAULT_RECORD;
    pokemonDialog.value = true;
}

function savePokemon() {
    if (isEditing.value) {
        currentService.value.updatePokemon(record.value.id, record.value);
        messages.add("Pokemon updated successfully!");
    } else {
        currentService.value.createPokemon(record.value);
        messages.add("Pokemon created successfully!");
    }
    pokemonDialog.value = false;
}

function editPokemon(id) {
    isEditing.value = true;
    pokemonDialog.value = true;

    record.value = {
        ...DEFAULT_RECORD,
        ...items.value.find((pokemon) => pokemon.id === id),
    };
}

function deletePokemon(id) {
    currentService.value.deletePokemon(id);
    items.value = items.value.filter((pokemon) => pokemon.id !== id);
    messages.add("Pokemon deleted successfully!");
}

// function edit(id) {
//     isEditing.value = true;

//     const found = items.value.find((book) => book.id === id);

//     record.value = {
//         id: found.id,
//         name: found.name,
//         type: found.type,
//         abilities: found.abilities,
//         height: found.height,
//         weight: found.weight,
//         pokedex_id: found.pokedex_id,
//     };
//     pokemonDialog.value = true;
// }
</script>

<style scoped>
.flex {
    display: flex;
    flex-direction: row;
    gap: 10px;
}
</style>
