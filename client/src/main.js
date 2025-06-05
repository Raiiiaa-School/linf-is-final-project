import { createApp, h } from "vue";
import App from "./App.vue";

import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

import {
    createClient,
    provideClient,
    debugExchange,
    cacheExchange,
    fetchExchange,
} from "@urql/vue";
import { createPinia } from "pinia";

const vuetify = createVuetify({
    components,
    directives,
});

const pinia = createPinia();

const graphqlClient = createClient({
    url: process.env.VUE_APP_GRAPHQL_SERVICE || "http://localhost:8082/query",
    exchanges: [debugExchange, cacheExchange, fetchExchange],
});

const app = createApp({
    setup() {
        provideClient(graphqlClient);
    },
    render: () => h(App),
});

app.use(pinia);
app.use(vuetify);
app.mount("#app");
