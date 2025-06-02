<template>
  <v-card class="pa-4 mt-4" elevation="2">
    <v-card-title>Buscar Utilizador</v-card-title>
    <v-card-text>
      <v-select
        v-model="searchMode"
        :items="['Buscar por ID', 'Buscar por Username']"
        label="Modo de Busca"
        variant="outlined"
        class="mb-4"
      ></v-select>

      <v-text-field
        v-if="searchMode === 'Buscar por ID'"
        v-model="userId"
        label="ID do Utilizador"
        :rules="[v => !!v || 'ID é obrigatório']"
        required
      ></v-text-field>

      <v-text-field
        v-if="searchMode === 'Buscar por Username'"
        v-model="username"
        label="Nome de Utilizador"
        :rules="[v => !!v || 'Nome de Utilizador é obrigatório']"
        required
      ></v-text-field>

      <v-btn @click="handleSearch" color="primary" class="mt-4">Buscar</v-btn>

      <v-divider class="my-4"></v-divider>

      <div v-if="loading">A buscar utilizador...</div>
      <div v-else-if="error">{{ error }}</div>
      <div v-else-if="user">
        <h4>Detalhes do Utilizador:</h4>
        <p><strong>ID:</strong> {{ user.id }}</p>
        <p><strong>Nome:</strong> {{ user.name }}</p>
        <p><strong>Password (para fins de teste):</strong> {{ user.password }}</p> 
      </div>
      <div v-else-if="searched">Nenhum utilizador encontrado com os critérios fornecidos.</div>
    </v-card-text>
    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">{{ snackbarText }}</v-snackbar>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      searchMode: 'Buscar por ID', // 'Buscar por ID' ou 'Buscar por Username'
      userId: '',
      username: '',
      user: null, // Para armazenar o utilizador encontrado
      loading: false,
      error: null,
      searched: false, // Flag para saber se já foi feita uma busca
      snackbar: false,
      snackbarText: '',
      snackbarColor: ''
    };
  },
  methods: {
    async handleSearch() {
      this.loading = true;
      this.error = null;
      this.user = null;
      this.searched = true;

      const baseUrl = import.meta.env.VITE_REST_API_URL;
      let url = '';

      if (this.searchMode === 'Buscar por ID') {
        if (!this.userId) {
          this.showSnackbar('Por favor, insere o ID do Utilizador.', 'error');
          this.loading = false;
          return;
        }
        url = `${baseUrl}/users/${this.userId}`;
      } else if (this.searchMode === 'Buscar por Username') {
        if (!this.username) {
          this.showSnackbar('Por favor, insere o Nome de Utilizador.', 'error');
          this.loading = false;
          return;
        }
        url = `${baseUrl}/users/username/${this.username}`;
      }

      try {
        const response = await fetch(url);
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(`Erro: ${response.statusText} - ${errorData.message || 'Utilizador não encontrado.'}`);
        }
        this.user = await response.json();
        this.showSnackbar("Utilizador encontrado com sucesso!", 'success');
      } catch (error) {
        console.error("Erro ao buscar utilizador:", error);
        this.error = `Erro ao buscar utilizador: ${error.message || error}`;
        this.showSnackbar(`Erro ao buscar utilizador: ${error.message || error}`, 'error');
      } finally {
        this.loading = false;
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