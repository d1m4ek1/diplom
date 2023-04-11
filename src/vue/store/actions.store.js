import Admin from "../../models/admin";
import CSRF from "../../models/csrf";
import EditorAdmin from "../../models/editor";

const actions = {
  async setQuestions({ commit }) {
    await Admin.GetAllQuestions()
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          commit("setQuestions", response.items);
        }
      });
  },
  async deleteQuestionById({ state, getters, commit }, id) {
    await new Admin(getters.getXCSRFToken)
      .DeleteQuestionById(id)
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
  async saveEditItem({ commit, dispatch, getters }, data) {
    await new EditorAdmin(getters.getXCSRFToken)
      .SaveContent(data)
      .then((response) => response.json())
      .then(async (response) => {
        if (response.successfully) {
          commit("saveEditItem", data);
          await dispatch("setEditItems");
          commit("setActualSiteContent", [...getters.getEditItems].pop());
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
  async setNewActualSiteContentFromHistory({ getters, commit }, id) {
    await new EditorAdmin(getters.getXCSRFToken)
      .SetNewActualSiteContentFromHistory(id)
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          commit("setActualSiteContent", response.data);
        }
      });
  },
  async deleteSiteContentFromHistory({ getters, commit }, id) {
    await new EditorAdmin(getters.getXCSRFToken)
      .DeleteSiteContentFromHistory(id)
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          commit("deleteSiteContentFromHistory", id);
          commit("setActualSiteContent", [...getters.getEditItems].pop());
        }
      });
  },
  async sortList({ commit }, data) {
    await EditorAdmin.FilterSortList(data)
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          switch (data.nameList) {
            case "history":
              commit("setEditItems", response.items);
              break;

            case "questions":
              commit("setQuestions", response.items);
              break;
          }
        }
      });
  },
  async setActionXCSRFToken({ commit }) {
    await CSRF.GetToken().then((response) => commit("setCSRFToken", response.headers.get("X-CSRF-TOKEN")));
  },
};

export default actions;
