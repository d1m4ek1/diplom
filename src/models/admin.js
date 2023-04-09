class Admin {
  static async GetAllQuestions() {
    return await fetch("/api/admin/all", {
      method: "GET",
    });
  }

  static async DeleteQuestionById(id) {
    return await fetch(`/api/admin/delete?id=${id}`, {
      method: "DELETE",
    });
  }
}

export default Admin;