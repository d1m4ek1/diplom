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
          <p>{{ toDate(item.addDate) }}</p>
          <p>Email: {{ item.email }}</p>
          <p>Номер телефона: {{ item.tel }}</p>
          <button @click="openModal(item.question)">Читать вопрос</button>
          <button>Удалить вопрос</button>
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
import { mapGetters } from "vuex";
import ModalReadQuestion from "../components/ModalReadQuestion.vue";

export default {
  components: { ModalReadQuestion },
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
    toDate(date) {
      return new Intl.DateTimeFormat("ru-RU", {
        day: "2-digit",
        month: "long",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      }).format(new Date(date));
    },
    openModal(content) {
      this.forModal.content = content;
      this.forModal.opened = true;
    },
    onCloseWindow() {
      this.forModal.opened = false;
      this.forModal.content = "";
    },
  },
  created() {
    this.questions = this.getQuestions;
  },
  computed: {
    ...mapGetters(["getQuestions"]),
  },
};
</script>