<template>
  <header class="header">
    <nav class="header_navigation">
      <ul class="header_navigation__items">
        <router-link to="/admin" v-slot="{ href, navigate, isActive }" custom>
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Вопросы</a>
          </li>
        </router-link>
        <router-link
          to="/admin/site-edit"
          v-slot="{ href, navigate, isActive }"
          custom
        >
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Редактирование контента</a>
          </li>
        </router-link>
        <li>
          <a @click="logout()">Выйти</a>
        </li>
      </ul>
    </nav>
  </header>
  <main class="main">
    <router-view></router-view>
  </main>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import Admin from "../../models/admin";

export default {
  name: "AppVue",
  methods: {
    ...mapActions(["setQuestions", "setActionXCSRFToken"]),
    async logout() {
      await Admin.AdminLogout(this.getXCSRFToken)
        .then((response) => response.json())
        .then((response) =>
          response.successfully
            ? (location.href = "/admin/login")
            : console.log(response)
        );
    },
  },
  computed: {
    ...mapGetters(["getXCSRFToken"]),
  },
  async created() {
    await this.setQuestions();
    await this.setActionXCSRFToken();
  },
};
</script>

<style scoped>
.header_navigation__items li:last-child {
  margin-top: 30px;
}
</style>