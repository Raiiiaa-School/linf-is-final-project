<template>
  <div>
    <h3>Lista de Pokémon</h3>
    <PokemonCard v-for="poke in pokemons" :key="poke.id" :pokemon="poke" />
  </div>
</template>

<script>
import PokemonCard from '../common/PokemonCard.vue';

export default {
  components: { PokemonCard },
  data() {
    return { pokemons: [] }
  },
  async mounted() {
    const res = await fetch('http://localhost:8081/pokemons')
    this.pokemons = await res.json()
  }
}
</script>