<template>
  <div class="card">
    <h4>{{ pokemon.name }}</h4>
    <p>ID: {{ pokemon.pokedex_id }}</p>
    <p>Altura: {{ pokemon.height }}m | Peso: {{ pokemon.weight }}kg</p>
    <p>Tipo: {{ pokemon.type.join(', ') }}</p>
    <p>Habilidades: {{ pokemon.abilities.join(', ') }}</p>
  </div>
</template>

<script>
export default {
  props: ['pokemon']
}
</script>