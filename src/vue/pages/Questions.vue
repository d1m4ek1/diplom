<template>
  <div class="container">
    <div class="content__header"><h1>Все вопросы</h1></div>
    <div v-if="questions === null || questions.length === 0" class="content">
      <p>Вопрос пока нет</p>
    </div>
    <div v-else class="content">
      <Filter :name="'history'" />

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
import Filter from "../components/Filter.vue";
import Date from "../components/Date.vue";

export default {
  components: { ModalReadQuestion, Date, Filter },
  name: "QuesitionsVue",
  data() {
    return {
      questions: [],
      forModal: {
        opened: false,
        content: "",
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
    ...mapActions(["deleteQuestionById"]),
  },
  created() {
    this.questions = this.getQuestions;
  },
  computed: {
    ...mapGetters(["getQuestions"]),
  },
};
</script>