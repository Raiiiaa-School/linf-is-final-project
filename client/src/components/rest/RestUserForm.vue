<template>
  <v-card class="pa-4 mt-4" elevation="2">
    <v-card-title>Criar Novo Utilizador</v-card-title>
    <v-card-text>
      <v-form @submit.prevent="handleSubmit">
        <v-text-field
          v-model="userData.name"
          label="Nome de Utilizador"
          :rules="[v => !!v || 'Nome é obrigatório']"
          required
        ></v-text-field>
        <v-text-field
          v-model="userData.password"
          label="Password"
          type="password"
          :rules="[v => !!v || 'Password é obrigatória']"
          required
        ></v-text-field>
        <v-btn type="submit" color="primary" class="mt-4">Criar Utilizador</v-btn>
      </v-form>
    </v-card-text>
    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">{{ snackbarText }}</v-snackbar>
  </v-card>
</template>

<script setup>
import axios from "axios";
import { reactive, ref } from "vue";

const userData = reactive({
  name: "",
  password: "",
});

const snackbar = ref(false);
const snackbarText = ref("");
const snackbarColor = ref("");

function handleSubmit() {
  createUser();
}

async function createUser() {
  try {
    // CORRIGIDO: Usando process.env.VUE_APP_REST_SERVICE_URL para a base
    const baseUrl = process.env.VUE_APP_REST_SERVICE_URL;
    if (!baseUrl) {
      showSnackbar("Variavel de ambiente VUE_APP_REST_SERVICE_URL não está definida.", "error");
      return;
    }
    // CORRIGIDO: Adiciona '/users' à baseUrl
    const response = await axios.post(`${baseUrl}/users`, userData);
    console.log("User criado:", response);
    showSnackbar("Utilizador criado com sucesso!", "success");
    // Limpar o formulário após a criação bem-sucedida
    Object.assign(userData, { name: "", password: "" });
  } catch (error) {
    showSnackbar(
      `Erro ao criar User: ${error.response?.data?.message || error.message || error}`,
      "error"
    );
    console.error("Erro ao criar utilizador:", error);
  }
}

function showSnackbar(message, color) {
  snackbarText.value = message;
  snackbarColor.value = color;
  snackbar.value = true;
}
</script>