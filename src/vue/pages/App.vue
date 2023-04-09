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
      </ul>
    </nav>
  </header>
  <main class="main">
    <router-view></router-view>
  </main>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "AppVue",
  methods: {
    ...mapActions(["setQuestions"]),
  },
  async created() {
    await this.setQuestions();
  },
};
</script>