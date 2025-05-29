<template>
  <div>
    <h2>REST API</h2>
    <RestPokemonForm @pokemonCreated="refresh" />
    <RestPokemonList :key="reloadKey" />
  </div>
</template>

<script>
import RestPokemonForm from '../components/rest/RestPokemonForm.vue'
import RestPokemonList from '../components/rest/RestPokemonList.vue'

export default {
  components: { RestPokemonForm, RestPokemonList },
  data: () => ({ reloadKey: 0 }),
  methods: {
    refresh() {
      this.reloadKey++
    }
  }
}
</script>