<template>
  <form @submit.prevent="submit">
    <input v-model="name" placeholder="Nome" required />
    <input v-model="pokedex_id" placeholder="ID Pokedex" />
    <input v-model="height" placeholder="Altura" />
    <input v-model="weight" placeholder="Peso" />
    <input v-model="type" placeholder="Tipos (separados por vírgula)" />
    <input v-model="abilities" placeholder="Habilidades (separadas por vírgula)" />
    <button type="submit">Criar</button>
  </form>
</template>

<script>
export default {
  data() {
    return { name: '', pokedex_id: '', height: '', weight: '', type: '', abilities: '' }
  },
  methods: {
    async submit() {
      const res = await fetch('http://localhost:8081/pokemons', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name: this.name,
          pokedex_id: parseInt(this.pokedex_id),
          height: parseFloat(this.height),
          weight: parseFloat(this.weight),
          type: this.type.split(',').map(t => t.trim()),
          abilities: this.abilities.split(',').map(a => a.trim()),
        })
      })
      if (res.ok) {
        this.$emit('pokemonCreated')
        alert('Criado com sucesso!')
      }
    }
  }
}
</script>