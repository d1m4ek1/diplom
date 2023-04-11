<template>
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
</template>

<script>
import { mapGetters } from "vuex";
import Admin from "../../models/admin";

export default {
  name: "NavBar",
  methods: {
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
};
</script>

<style>
</style>