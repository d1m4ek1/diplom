import Admin from "../../models/admin";
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
  async saveEditItem({ commit, dispatch, getters }, data) {
    await EditorAdmin.SaveContent(data)
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
  async setNewActualSiteContentFromHistory({ commit }, id) {
    await EditorAdmin.SetNewActualSiteContentFromHistory(id)
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          commit("setActualSiteContent", response.data);
        }
      });
  },
  async deleteSiteContentFromHistory({ getters, commit }, id) {
    await EditorAdmin.DeleteSiteContentFromHistory(id)
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
          commit("setEditItems", response.items);
        }
      });
  },
};

export default actions;
