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
        <p><strong>ID:</strong> {{ user._id || user.id }}</p> <p><strong>Nome:</strong> {{ user.name }}</p>
        <p><strong>Password (para fins de teste):</strong> {{ user.password }}</p>
      </div>
      <div v-else-if="searched && !user">Nenhum utilizador encontrado com os critérios fornecidos.</div>
    </v-card-text>
    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">{{ snackbarText }}</v-snackbar>
  </v-card>
</template>

<script setup>
import axios from "axios";
import { ref } from "vue";

const searchMode = ref('Buscar por ID');
const userId = ref('');
const username = ref('');

const loading = ref(false);
const error = ref(null);
const user = ref(null);
const searched = ref(false);

const snackbar = ref(false);
const snackbarText = ref("");
const snackbarColor = ref("");

async function handleSearch() {
  loading.value = true;
  error.value = null;
  user.value = null;
  searched.value = true;

  try {
    // CORRIGIDO: Agora baseUrl contém APENAS o endereço base do serviço
    const baseUrl = process.env.VUE_APP_REST_SERVICE_URL;

    if (!baseUrl) {
      showSnackbar("Variável de ambiente VUE_APP_REST_SERVICE_URL não está definida.", "error");
      loading.value = false;
      return;
    }

    let response;
    if (searchMode.value === 'Buscar por ID') {
      if (!userId.value) {
        showSnackbar("Por favor, insira o ID do utilizador.", "warning");
        loading.value = false;
        return;
      }
      // CORRIGIDO: Adiciona '/users/${userId.value}' à baseUrl
      response = await axios.get(`${baseUrl}/users/${userId.value}`);
    } else { // Buscar por Username
      if (!username.value) {
        showSnackbar("Por favor, insira o nome de utilizador.", "warning");
        loading.value = false;
        return;
      }
      // CORRIGIDO: Adiciona '/users?name=' à baseUrl
      response = await axios.get(`${baseUrl}/users?name=${username.value}`);
    }

    if (response.status === 200) {
      const responseData = response.data;
      user.value = (Array.isArray(responseData) && responseData.length > 0)
        ? responseData[0]
        : (responseData && !Array.isArray(responseData) ? responseData : null);

      if (!user.value) {
        showSnackbar("Nenhum utilizador encontrado.", "info");
      } else {
        showSnackbar("Utilizador encontrado com sucesso!", "success");
      }
    } else {
      showSnackbar(`Erro ao buscar utilizador: Status ${response.status}`, "error");
    }
  } catch (err) {
    error.value = `Erro ao buscar utilizador: ${err.message || err}`;
    console.error("Erro ao buscar utilizador:", err);
    showSnackbar(`Erro: ${err.response?.data?.message || err.message || 'Erro desconhecido'}`, "error");
    user.value = null;
  } finally {
    loading.value = false;
  }
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