class CSRF {
  static async GetToken() {
    return await fetch("/api/token/get_token", {
      method: "GET",
    });
  }
}

export default CSRF;
