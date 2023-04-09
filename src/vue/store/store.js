import { createStore } from "vuex";

import actions from "./actions.store";
import mutations from "./mutations.store";
import getters from "./getters.store";

const store = createStore({
  state() {
    return {
      isAuth: false,
      X_CSRF_TOKEN: "",
      questions: [],
      editItems: [],
      actualSiteContent: {
        _id: 0,
        header: "",
        content: "",
        addDate: "",
      },
    };
  },
  getters,
  mutations,
  actions,
});

export default store;
