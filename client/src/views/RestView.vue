<template>
  <div>
    <RestPokemonForm @pokemonCreated="refreshList" />
    <RestPokemonList :key="reloadKey" />
  </div>
</template>

<script>
import RestPokemonForm from '@/components/rest/RestPokemonForm.vue';
import RestPokemonList from '@/components/rest/RestPokemonList.vue';

export default {
  name: 'RestView',
  components: {
    RestPokemonForm,
    RestPokemonList
  },
  data() {
    return {
      reloadKey: 0
    };
  },
  methods: {
    refreshList() {
      this.reloadKey++;
    }
  }
};
</script>