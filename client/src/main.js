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

export const graphqlClient = createClient({
    url: process.env.VUE_APP_GRAPHQL_SERVICE || "http://localhost:8082/query",
    exchanges: [debugExchange, cacheExchange, fetchExchange],
    fetchOptions: () => {
        const token = localStorage.getItem("authToken");
        return {
            headers: {
                Authorization: token ? `Bearer ${token}` : "",
            },
        };
    },
});

const app = createApp({
    setup() {
        provideClient(graphqlClient);
    },
    render: () => h(App),
});

app.config.devtools = false;
app.config.warnHandler = () => null;
app.config.errorHandler = () => null;

app.use(pinia);
app.use(vuetify);
app.mount("#app");
