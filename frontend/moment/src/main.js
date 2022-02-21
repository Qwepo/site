import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";





import HTTpPlugin from './plugins/http'
createApp(App).use(store).use(router).use(HTTpPlugin).mount("#app");

