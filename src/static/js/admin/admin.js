import Admin from "../../../models/admin";

const getQuestions = async () => {
  const questions = await Admin.GetAllQuestions();

  const json = await questions.json();

  setTimeout(() => {
    if (json.successfully) {
      setAllQuestions(json.items);
    }
  }, 2000);
};

function toDate(date) {
  return new Intl.DateTimeFormat("ru-RU", {
    day: "2-digit",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  }).format(new Date(date));
}

function setAllQuestions(data = []) {
  const $items = document.getElementById("question_items");

  $items.insertAdjacentHTML(
    "afterbegin",
    data
      .map((item) => {
        return `<div class="question_item">
    <p>Идентификатор: ${item._id}</p>
    <p>ФИО: ${item.secondName} ${item.firstName} ${item.thirdName}</p>
    <p>${item.typeUser}</p>
    <p>${toDate(item.addDate)}</p>
    <div class="connect">
      <p>Email: ${item.email}</p>
      <p>Номер телефона: ${item.tel}</p>
    </div>
    <button>Читать вопрос</button>
    <button>Удалить вопрос</button>
  </div>`;
      })
      .join("")
  );
}

getQuestions();
