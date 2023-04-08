class Admin {
  static async GetAllQuestions() {
    return await fetch("/api/admin/all", {
      method: "GET",
    });
  }
}

export default Admin;
