class EditorAdmin {
  constructor(token) {
    this.token = token;
  }

  static async GetActualContent() {
    return await fetch("/api/admin/editor/actual", {
      method: "GET",
    });
  }

  async SaveContent(data) {
    return await fetch("/api/admin/editor/save", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": this.token,
      },
    });
  }

  static async GetAllEditor() {
    return await fetch("/api/admin/editor/all", {
      method: "GET",
    });
  }

  async SetNewActualSiteContentFromHistory(_id) {
    return await fetch(`/api/admin/editor/set_new`, {
      method: "PUT",
      body: JSON.stringify({
        _id,
      }),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": this.token,
      },
    });
  }

  async DeleteSiteContentFromHistory(_id) {
    return await fetch(`/api/admin/editor/delete`, {
      method: "DELETE",
      body: JSON.stringify({
        _id,
      }),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": this.token,
      },
    });
  }

  static async FilterSortList(data) {
    const urlSearchParams = new URLSearchParams({
      sort_by: data.sortList,
      sort_by_date: data.sortListByDate,
      name_list: data.nameList,
    });

    return await fetch(`/api/admin/editor/sort_list?` + urlSearchParams, {
      method: "GET",
    });
  }
}

export default EditorAdmin;
