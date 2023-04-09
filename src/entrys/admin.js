import { createApp } from "vue";
import router from "../vue/routes/index.routes";
import store from "../vue/store/store";
import App from "../vue/pages/App.vue";

createApp(App).use(store).use(router).mount("#app");
