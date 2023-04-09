const mutations = {
  userLogin: (state) => (state.isAuth = true),
  userLogout: (state) => (state.isAuth = false),
  setCSRFToken: (state, token) => (state.X_CSRF_TOKEN = token),
  setQuestions: (state, data) => (state.questions = data),
  setEditItems: (state, data) => (state.editItems = data),
  setActualSiteContent: (state, data) => (state.actualSiteContent = { ...data }),
  deleteSiteContentFromHistory: (state, id) => (state.editItems = state.editItems.filter((e) => e._id !== id)),
  saveEditItem: (state, data) => state.editItems.push(data),
};

export default mutations;
