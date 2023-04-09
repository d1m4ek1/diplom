import { createStore } from "vuex";
import Admin from "../../models/admin";
import EditorAdmin from "../../models/editor";

const store = createStore({
  state() {
    return {
      isAuth: false,
      X_CSRF_TOKEN: "",
      questions: [],
      editItems: [],
      actualSiteContent: {
        id: 0,
        header: "",
        content: "",
        addDate: "",
      },
    };
  },
  getters: {
    getIsAuth: (state) => state.isAuth,
    getXCSRFToken: (state) => state.X_CSRF_TOKEN,
    getQuestions: (state) => state.questions,
    getEditItems: (state) => state.editItems,
    getActualSiteContent: (state) => state.actualSiteContent,
  },
  mutations: {
    userLogin: (state) => (state.isAuth = true),
    userLogout: (state) => (state.isAuth = false),
    setCSRFToken: (state, token) => (state.X_CSRF_TOKEN = token),
    setQuestions: (state, data) => (state.questions = data),
    setEditItems: (state, data) => (state.editItems = data),
    setActualSiteContent: (state, data) => (state.actualSiteContent = { ...data }),
    saveEditItem: (state, data) => state.editItems.push(data),
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
    async deleteQuestionById({ state, commit }, id) {
      await Admin.DeleteQuestionById(id)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            commit(
              "setQuestions",
              state.questions.filter((q) => q._id !== id)
            );
          }
        });
    },
    async saveEditItem({ commit }, data) {
      await EditorAdmin.SaveContent(data)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            commit("saveEditItem", data);
          }
        });
    },
    async setEditItems({ commit }) {
      await EditorAdmin.GetAllEditor()
        .then((response) => response.json())
        .then((response) => {
          commit("setEditItems", response.data);
        });
    },
    async setNewActualSiteContentFromHistory({ commit }, id) {
      await EditorAdmin.setNewActualSiteContentFromHistory(id)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            commit("setActualSiteContent", response.data);
          }
        });
    },
  },
});

export default store;
