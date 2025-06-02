<template>
    <div>
        <GraphQLPokemonForm @pokemonCreated="refreshList" />
        <GraphQLPokemonList :key="reloadKey" />
    </div>
</template>

<script setup>
import GraphQLPokemonForm from "@/components/graphql/GraphQLPokemonForm.vue";
import GraphQLPokemonList from "@/components/graphql/GraphQLPokemonList.vue";
import { ref } from "vue";

const reloadKey = ref(0);

function refreshList() {
    reloadKey.value++;
}
</script>
