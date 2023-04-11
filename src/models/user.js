import CSRF from "./csrf";

class User {
  static async sendForm(data) {
    let token = "";

    await CSRF.GetToken().then((response) => (token = response.headers.get("X-CSRF-TOKEN")));

    return await fetch("/api/form/send", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": token,
      },
    });
  }
}

export default User;
