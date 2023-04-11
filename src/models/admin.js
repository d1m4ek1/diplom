class Admin {
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

  static async AdminLogin(token, data) {
    return await fetch("/api/admin/auth/login", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "X-CSRF-TOKEN": token,
        "Content-Type": "application/json",
      },
    });
  }

  static async AdminLogout(token) {
    return await fetch("/api/admin/auth/logout", {
      method: "PATCH",
      headers: {
        "X-CSRF-TOKEN": token,
        "Content-Type": "application/json",
      },
    });
  }
}

export default Admin;
