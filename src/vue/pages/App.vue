<template>
  <header class="header">
    <NavBar />
  </header>
  <main class="main">
    <router-view></router-view>
  </main>
</template>

<script>
import { mapActions } from "vuex";
import NavBar from "../components/NavBar.vue";

export default {
  name: "AppVue",
  methods: {
    ...mapActions(["setQuestions", "setActionXCSRFToken"]),
  },
  async created() {
    await this.setQuestions();
    await this.setActionXCSRFToken();
  },
  components: {
    NavBar,
  },
};
</script>

<style scoped>
.header_navigation__items li:last-child {
  margin-top: 30px;
}
</style>