// src/main.js
import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // <- IMPORTA O ROUTER CORRETAMENTE

const app = createApp(App);

app.use(router); // <- USA O ROUTER ANTES DE MONTAR
app.mount('#app');
