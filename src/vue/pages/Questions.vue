<template>
  <div class="container">
    <div class="content__header"><h1>Все вопросы</h1></div>
    <div class="content">
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
      <div class="question_items">
        <div
          class="question_item"
          v-for="(item, idx) in questions"
          :key="`question${idx}`"
        >
          <p>ID: {{ item._id }}</p>
          <p>
            ФИО: {{ `${item.secondName} ${item.firstName} ${item.thirdName}` }}
          </p>
          <p>{{ item.typeUser }}</p>
          <p><Date :date="item.addDate" /></p>
          <p>Email: {{ item.email }}</p>
          <p>Номер телефона: {{ item.tel }}</p>
          <button @click="openModal(item.question)">Читать вопрос</button>
          <button @click="deleteQuestionById(item._id)">Удалить вопрос</button>
        </div>
      </div>
    </div>
    <ModalReadQuestion
      :show="forModal.opened"
      :content="forModal.content"
      @on-close-window="onCloseWindow"
    />
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import ModalReadQuestion from "../components/ModalReadQuestion.vue";
import Date from "../components/Date.vue";

export default {
  components: { ModalReadQuestion, Date },
  name: "QuesitionsVue",
  data() {
    return {
      questions: [],
      forModal: {
        opened: false,
        content: "",
      },
      filterData: {
        sortList: "ascending",
        sortListByDate: "ascending",
        nameList: "questions",
      },
    };
  },
  watch: {
    getQuestions(newValue) {
      this.questions = newValue;
    },
  },
  methods: {
    openModal(content) {
      this.forModal.content = content;
      this.forModal.opened = true;
    },
    onCloseWindow() {
      this.forModal.opened = false;
      this.forModal.content = "";
    },
    async filterSortList(key) {
      Object.keys(this.filterData).forEach((k) => {
        if (k === "nameList") return;
        if (key !== k) {
          this.filterData[k] = "ascending";
        }
        return;
      });

      await this.sortList(this.filterData);
    },
    ...mapActions(["deleteQuestionById", "sortList"]),
  },
  created() {
    this.questions = this.getQuestions;
  },
  computed: {
    ...mapGetters(["getQuestions"]),
  },
};
</script>

<style scoped>
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