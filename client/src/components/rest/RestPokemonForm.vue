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
          :rules="[v => !!v || 'Nome é obrigatório']"
          required
          :disabled="mode === 'delete'"
        ></v-text-field>
        <v-text-field
          v-model.number="pokemonData.pokedex_id"
          label="ID da Pokedex"
          type="number"
          :rules="[v => v >= 0 || 'ID deve ser positivo']"
          :disabled="mode === 'delete'"
        ></v-text-field>
        <v-text-field
          v-model.number="pokemonData.height"
          label="Altura (m)"
          type="number"
          step="0.01"
          :rules="[v => v >= 0 || 'Altura deve ser positiva']"
          :disabled="mode === 'delete'"
        ></v-text-field>
        <v-text-field
          v-model.number="pokemonData.weight"
          label="Peso (kg)"
          type="number"
          step="0.01"
          :rules="[v => v >= 0 || 'Peso deve ser positivo']"
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
          :rules="[v => !!v || 'ID é obrigatório para esta ação']"
          required
        ></v-text-field>

        <v-btn type="submit" color="primary" class="mt-4">
          <span v-if="mode === 'create'">Criar Pokémon</span>
          <span v-else-if="mode === 'update'">Atualizar Pokémon</span>
          <span v-else-if="mode === 'delete'">Eliminar Pokémon</span>
        </v-btn>
      </v-form>
    </v-card-text>
    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">{{ snackbarText }}</v-snackbar>
  </v-card>
</template>

<script>
export default {
  // O `mode` prop dirá ao componente o que ele deve fazer (criar, atualizar, eliminar)
  props: {
    mode: {
      type: String,
      default: 'create', // 'create', 'update', 'delete'
      validator: (value) => ['create', 'update', 'delete'].includes(value)
    },
    // Para o modo de atualização/eliminação, pode-se passar um Pokémon inicial para preencher o formulário
    initialPokemon: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      pokemonData: {
        name: '',
        pokedex_id: null,
        height: null,
        weight: null,
        type: [],
        abilities: [],
        ...this.initialPokemon // Preenche com dados iniciais se fornecidos
      },
      pokemonId: this.initialPokemon._id || '', // Para operações de update/delete, assumindo que o ID é _id do MongoDB
      snackbar: false,
      snackbarText: '',
      snackbarColor: ''
    };
  },
  watch: {
    // Observa mudanças em initialPokemon para atualizar o formulário se a prop mudar
    initialPokemon: {
      handler(newVal) {
        if (newVal) {
          this.pokemonData = { ...newVal };
          this.pokemonId = newVal._id || '';
        }
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    async handleSubmit() {
      const baseUrl = import.meta.env.VITE_REST_API_URL;
      let url = '';
      let method = '';
      let body = null;

      if (this.mode === 'create') {
        url = `${baseUrl}/pokemons`;
        method = 'POST';
        body = JSON.stringify({
          ...this.pokemonData,
          pokedex_id: parseInt(this.pokemonData.pokedex_id),
          height: parseFloat(this.pokemonData.height),
          weight: parseFloat(this.pokemonData.weight)
        });
      } else if (this.mode === 'update') {
        if (!this.pokemonId) {
          this.showSnackbar('Por favor, insere o ID do Pokémon.', 'error');
          return;
        }
        url = `${baseUrl}/pokemons/${this.pokemonId}`;
        method = 'PUT';
        body = JSON.stringify({
          ...this.pokemonData,
          pokedex_id: parseInt(this.pokemonData.pokedex_id),
          height: parseFloat(this.pokemonData.height),
          weight: parseFloat(this.pokemonData.weight)
        });
      } else if (this.mode === 'delete') {
        if (!this.pokemonId) {
          this.showSnackbar('Por favor, insere o ID do Pokémon para eliminar.', 'error');
          return;
        }
        if (!confirm(`Tens certeza que queres eliminar o Pokémon com ID: ${this.pokemonId}?`)) {
          return;
        }
        url = `${baseUrl}/pokemons/${this.pokemonId}`;
        method = 'DELETE';
      }

      try {
        const response = await fetch(url, {
          method: method,
          headers: {
            'Content-Type': 'application/json'
            // 'Authorization': 'Bearer <your-jwt-token>' // Adicionar se a API exigir
          },
          body: body // O corpo é null para GET e DELETE
        });

        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(`Erro: ${response.statusText} - ${errorData.message || 'Erro desconhecido'}`);
        }

        if (this.mode === 'create') {
          this.showSnackbar("Pokémon criado com sucesso!", 'success');
          // Limpa o formulário após a criação
          this.pokemonData = { name: '', pokedex_id: null, height: null, weight: null, type: [], abilities: [] };
        } else if (this.mode === 'update') {
          this.showSnackbar("Pokémon atualizado com sucesso!", 'success');
        } else if (this.mode === 'delete') {
          this.showSnackbar("Pokémon eliminado com sucesso!", 'success');
          this.pokemonId = ''; // Limpa o ID após a eliminação
          this.pokemonData = { name: '', pokedex_id: null, height: null, weight: null, type: [], abilities: [] }; // Opcional: limpar campos também
        }

        // Emite um evento para o componente pai para que ele saiba que a ação foi concluída
        this.$emit("actionCompleted");
      } catch (error) {
        console.error(`Erro na operação de ${this.mode} Pokémon:`, error);
        this.showSnackbar(`Erro ao ${this.mode} Pokémon: ${error.message || error}`, 'error');
      }
    },
    showSnackbar(message, color) {
      this.snackbarText = message;
      this.snackbarColor = color;
      this.snackbar = true;
    }
  }
};
</script>

<style scoped>
/* Estilos específicos para o formulário se necessário */
</style>