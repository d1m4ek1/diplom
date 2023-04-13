<template>
  <section class="new_block">
    <hr />
    <h2>История изменений</h2>

    <Filter :name="'history'" />

    <div class="history_edit_items">
      <p v-if="historyEditItems.length === 0">История пуста</p>
      <table v-else>
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
                'active-content-table': item._id === actualContentId,
              }"
            >
              <td>{{ item._id }}</td>
              <td><Dateui :date="item.addDate" /></td>
              <td>
                <button @click="openModal(item._id)">Просмотр</button>
              </td>
              <td>
                <button
                  v-if="item._id !== actualContentId"
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
    <ModalIframeVue
      :idhtml="forModal.idhtml"
      :show="forModal.opened"
      @on-close-window="onCloseWindow"
    />
  </section>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

import ModalIframeVue from "../../components/ModalIframe.vue";
import Filter from "../../components/Filter.vue";
import Dateui from "../../components/Date.vue";
import Loader from "../../components/Loader.vue";

export default {
  name: "HistoryModuleVue",
  props: {
    actualContentId: Number,
  },
  data() {
    return {
      forModal: {
        opened: false,
        idhtml: 0,
      },
      historyEditItems: [],
    };
  },
  watch: {
    getEditItems(newValue) {
      this.historyEditItems = newValue;
    },
  },
  methods: {
    openModal(id) {
      this.forModal.idhtml = id;
      this.forModal.opened = true;
    },
    onCloseWindow() {
      this.forModal.opened = false;
      this.forModal.content = "";
    },
    ...mapActions([
      "setNewActualSiteContentFromHistory",
      "deleteSiteContentFromHistory",
      "setEditItems",
    ]),
  },
  async created() {
    await this.setEditItems();
    this.historyEditItems = this.getEditItems;
  },
  computed: {
    ...mapGetters(["getEditItems"]),
  },
  components: {
    ModalIframeVue,
    Filter,
    Dateui,
    Loader,
  },
};
</script>

<style scoped>
.table-load {
  display: flex;
  justify-content: center;
}
.history_edit_items {
  display: flex;
  justify-content: center;
  max-height: 500px;
  overflow-y: auto;
}
</style>