const getters = {
  getIsAuth: (state) => state.isAuth,
  getXCSRFToken: (state) => state.X_CSRF_TOKEN,
  getQuestions: (state) => state.questions,
  getEditItems: (state) => state.editItems,
  getActualSiteContent: (state) => state.actualSiteContent,
};

export default getters;
