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
        <div class="filter_history">
          <h3>Фильтр</h3>
          <div class="filter_history__inputs">
            <div class="select_input_block">
              <p>Сортировать список</p>
              <select v-model="filterData.sortList">
                <option disabled>Сортировать список</option>
                <option
                  value="ascending"
                  selected
                  @click="filterSortList('sortList')"
                >
                  По возрастанию
                </option>
                <option value="descending" @click="filterSortList('sortList')">
                  По убыванию
                </option>
              </select>
            </div>
            <div class="select_input_block">
              <p>Сортировать по дате</p>
              <select v-model="filterData.sortListByDate">
                <option disabled>Сортировать по дате</option>
                <option
                  value="ascending"
                  selected
                  @click="filterSortList('sortListByDate')"
                >
                  По возрастанию
                </option>
                <option
                  value="descending"
                  @click="filterSortList('sortListByDate')"
                >
                  По убыванию
                </option>
              </select>
            </div>
          </div>
        </div>
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
      filterData: {
        sortList: "ascending",
        sortListByDate: "ascending",
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
      "sortList",
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
    async filterSortList(key) {
      Object.keys(this.filterData).forEach((k) => {
        if (key !== k) {
          this.filterData[k] = "ascending";
        }
        return;
      });

      await this.sortList(this.filterData);
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

.filter_history {
  padding-bottom: 20px;
}

.filter_history__inputs {
  display: flex;
  flex-wrap: wrap;
}

.select_input_block {
  display: flex;
  flex-direction: column;
  margin-right: 10px;
}

.select_input_block p {
  margin: 0;
}

.select_input_block select {
  border: 2px solid #eee;
  padding: 5px;
  min-width: 150px;
  height: 30px;
}
</style>