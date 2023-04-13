<template>
  <div class="container">
    <div class="content__header"><h1>Редактировать страницу</h1></div>
    <div class="content">
      <section>
        <h2>
          Добавлено изменение: <Dateui :date="dataActualContent.addDate" />
        </h2>
      </section>
      <section>
        <h2>Редактирование заголовка страницы</h2>
        <input v-model="dataActualContent.header" type="text" />
      </section>
      <section>
        <h2>Редактирование контента страницы</h2>
        <Loader v-if="dataActualContent.content === ''" />
        <Editor
          v-else
          api-key="0obopbpk2msn7rupwvdo0t2mk83kwla7kd0znsddmdd209yc"
          v-model="dataActualContent.content"
        />
      </section>
      <section>
        <hr />
        <button @click="saveEdit()">Сохранить изменения</button>
      </section>
      <HistoryModule :actual-content-id="dataActualContent._id" />
    </div>
  </div>
</template>

<script>
import Editor from "@tinymce/tinymce-vue";
import EditorAdmin from "../../models/editor";
import Dateui from "../components/Date.vue";
import HistoryModule from "./SiteEdit/History.module.vue";
import Loader from "../components/Loader.vue";

import { mapActions, mapGetters } from "vuex";

export default {
  name: "SiteEditVue",
  data() {
    return {
      dataActualContent: {
        _id: 0,
        content: "",
        header: "",
        addDate: "",
      },
    };
  },
  watch: {
    getActualSiteContent(newValue) {
      this.dataActualContent = { ...newValue };
    },
  },
  methods: {
    ...mapActions(["saveEditItem"]),
    async saveEdit() {
      this.dataActualContent.addDate = new Date().toISOString();

      await this.saveEditItem(this.dataActualContent);
    },
  },
  async created() {
    await EditorAdmin.GetActualContent()
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          this.dataActualContent = { ...response.data };
        }
      });
  },
  computed: {
    ...mapGetters(["getActualSiteContent"]),
  },
  components: {
    Editor,
    Dateui,
    HistoryModule,
    Loader,
  },
};
</script>

<style>
section > input {
  width: 100%;
}
</style>