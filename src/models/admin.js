class Admin {
  constructor(token) {
    this.token = token;
  }

  static async GetAllQuestions() {
    return await fetch("/api/admin/all", {
      method: "GET",
    });
  }

  async DeleteQuestionById(id) {
    return await fetch(`/api/admin/delete?id=${id}`, {
      method: "DELETE",
      headers: {
        "X-CSRF-TOKEN": this.token,
      },
    });
  }

  async AdminLogin(data) {
    return await fetch("/api/admin/auth/login", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "X-CSRF-TOKEN": this.token,
        "Content-Type": "application/json",
      },
    });
  }

  async AdminLogout() {
    return await fetch("/api/admin/auth/logout", {
      method: "PATCH",
      headers: {
        "X-CSRF-TOKEN": this.token,
        "Content-Type": "application/json",
      },
    });
  }
}

export default Admin;
