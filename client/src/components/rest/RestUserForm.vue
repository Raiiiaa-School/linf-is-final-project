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

<script>
export default {
  data() {
    return {
      userData: {
        name: '',
        password: '',
      },
      snackbar: false,
      snackbarText: '',
      snackbarColor: ''
    };
  },
  methods: {
    async handleSubmit() {
      const baseUrl = import.meta.env.VITE_REST_API_URL;
      try {
        const response = await fetch(`${baseUrl}/users`, { // Endpoint para criar utilizador
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.userData)
        });

        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(`Erro: ${response.statusText} - ${errorData.message || 'Erro desconhecido'}`);
        }

        const createdUser = await response.json();
        this.showSnackbar(`Utilizador '${createdUser.name}' criado com sucesso!`, 'success');
        
        // Limpar formulário
        this.userData.name = '';
        this.userData.password = '';

        this.$emit("actionCompleted"); // Informa o pai que a ação foi concluída
      } catch (error) {
        console.error("Erro ao criar utilizador:", error);
        this.showSnackbar(`Erro ao criar utilizador: ${error.message || error}`, 'error');
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