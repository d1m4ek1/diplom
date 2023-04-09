class EditorAdmin {
  static async GetActualContent() {
    return await fetch("/api/admin/editor/actual", {
      method: "GET",
    });
  }

  static async SaveContent(data) {
    return await fetch("/api/admin/editor/save", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  static async GetAllEditor() {
    return await fetch("/api/admin/editor/all", {
      method: "GET",
    });
  }
}

export default EditorAdmin;
