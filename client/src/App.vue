<template>
    <div id="app">
        <HomeView />
        <v-snackbar-queue v-model="messages.queue"></v-snackbar-queue>
    </div>
</template>

<script setup>
import { useMessagesStore } from "./stores/messages";
import HomeView from "./views/HomeView.vue";

const messages = useMessagesStore();
</script>
