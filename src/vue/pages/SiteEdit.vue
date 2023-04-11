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
        <Editor
          api-key="0obopbpk2msn7rupwvdo0t2mk83kwla7kd0znsddmdd209yc"
          v-model="dataActualContent.content"
        />
      </section>
      <section>
        <hr />
        <button @click="saveEdit()">Сохранить изменения</button>
      </section>
      <section class="new_block">
        <hr />
        <h2>История изменений</h2>

        <Filter :name="'history'" />

        <div class="history_edit_items">
          <table>
            <thead>
              <th>ID</th>
              <th>Дата изменения</th>
              <th>Просмотр изменения</th>
              <th>Установить изменение</th>
              <th>Удалить изменение</th>
            </thead>
            <tbody>
              <template
                v-for="(item, idx) in historyEditItems"
                :key="`column${idx}`"
              >
                <tr
                  :class="{
                    'active-content-table': item._id === dataActualContent._id,
                  }"
                >
                  <td>{{ item._id }}</td>
                  <td><Dateui :date="item.addDate" /></td>
                  <td>
                    <button @click="openModal(item._id)">Просмотр</button>
                  </td>
                  <td>
                    <button
                      v-if="item._id !== dataActualContent._id"
                      @click="setNewActualSiteContentFromHistory(item._id)"
                    >
                      Установить
                    </button>
                    <template v-else> Установлено </template>
                  </td>
                  <td>
                    <button
                      v-if="historyEditItems.length !== 1"
                      @click="deleteSiteContentFromHistory(item._id)"
                    >
                      Удалить
                    </button>
                  </td>
                </tr>
              </template>
            </tbody>
          </table>
        </div>
      </section>
    </div>
    <ModalIframeVue
      :idhtml="forModal.idhtml"
      :show="forModal.opened"
      @on-close-window="onCloseWindow"
    />
  </div>
</template>

<script>
import Editor from "@tinymce/tinymce-vue";
import EditorAdmin from "../../models/editor";
import Dateui from "../components/Date.vue";
import Filter from "../components/Filter.vue";
import { mapActions, mapGetters } from "vuex";

import ModalIframeVue from "../components/ModalIframe.vue";

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
      historyEditItems: [],
      forModal: {
        opened: false,
        idhtml: 0,
      },
    };
  },
  watch: {
    getActualSiteContent(newValue) {
      this.dataActualContent = { ...newValue };
    },
    getEditItems(newValue) {
      this.historyEditItems = newValue;
    },
  },
  methods: {
    ...mapActions([
      "saveEditItem",
      "setEditItems",
      "setNewActualSiteContentFromHistory",
      "deleteSiteContentFromHistory",
    ]),
    async saveEdit() {
      this.dataActualContent.addDate = new Date().toISOString();

      await this.saveEditItem(this.dataActualContent);
    },
    openModal(id) {
      this.forModal.idhtml = id;
      this.forModal.opened = true;
    },
    onCloseWindow() {
      this.forModal.opened = false;
      this.forModal.content = "";
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

    await this.setEditItems();

    this.historyEditItems = this.getEditItems;
  },
  computed: {
    ...mapGetters(["getEditItems", "getActualSiteContent"]),
  },
  components: {
    Editor,
    Dateui,
    ModalIframeVue,
    Filter,
  },
};
</script>

<style scoped>
.history_edit_items {
  max-height: 500px;
  overflow-y: auto;
}

section > input {
  width: 100%;
}

table {
  width: 100%;
}

thead {
  background-color: rgb(230, 230, 230);
}

table,
th,
td {
  border: 1px solid black;
  border-collapse: collapse;
}

th,
td {
  padding: 10px;
  text-align: center;
}
td > button {
  width: 100%;
}

table tr:nth-child(even) {
  background-color: rgb(250, 250, 250);
}

.active-content-table {
  background-color: rgb(210, 255, 210) !important;
}
</style>