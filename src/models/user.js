class User {
  static async sendForm(data) {
    return await fetch("/api/form/send", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}

export default User;
