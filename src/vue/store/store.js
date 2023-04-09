import { createStore } from "vuex";
import Admin from "../../models/admin";

const store = createStore({
  state() {
    return {
      isAuth: false,
      X_CSRF_TOKEN: "",
      questions: [],
    };
  },
  getters: {
    getIsAuth: (state) => state.isAuth,
    getXCSRFToken: (state) => state.X_CSRF_TOKEN,
    getQuestions: (state) => state.questions,
  },
  mutations: {
    userLogin(state) {
      state.isAuth = true;
    },
    userLogout(state) {
      state.isAuth = false;
    },
    setCSRFToken(state, token) {
      state.X_CSRF_TOKEN = token;
    },
    setQuestions(state, data) {
      state.questions = data;
    },
  },
  actions: {
    async setQuestions({ commit }) {
      await Admin.GetAllQuestions()
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            commit("setQuestions", response.items);
          }
        });
    },
  },
});

export default store;
