<template>
  <div class="container">
    <div class="content__header"><h1>Все вопросы</h1></div>
    <div class="content">
      <div class="filter"></div>
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