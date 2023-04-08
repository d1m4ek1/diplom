import User from "../../../models/user";
import { validateEmail } from "../utils/validEmail";

const formValues = {
  firstName: "",
  secondName: "",
  thirdName: "",
  typeUser: "абитуриент",
  email: "",
  tel: "",
  question: "",
  addDate: new Date().toISOString(),
};

const inputs = [
  document.getElementById("firstName"),
  document.getElementById("secondName"),
  document.getElementById("thirdName"),
  document.getElementById("email"),
  document.getElementById("tel"),
  document.getElementById("question"),
];

inputs.forEach(($element) => {
  $element.addEventListener("input", (event) => {
    const { value, id } = event.target;

    switch (id) {
      case "firstName":
        if (value !== "" && value.length > 1 && value.length <= 50) {
          formValues[id] = value;
          document.getElementById(id).classList.remove("error_input");
        }
        break;

      case "secondName" || "thirdName":
        if (value !== "" && value.length > 1 && value.length <= 50) {
          formValues[id] = value;
          document.getElementById(id).classList.remove("error_input");
        }
        break;

      case "email":
        if (value !== "" && value.length > 1 && value.length <= 126 && validateEmail(value)) {
          $element.classList.remove("error_input");
          formValues[id] = value;
        } else {
          $element.classList.add("error_input");
        }
        break;

      case "tel":
        if (value !== "" && value.length > 1 && value.length <= 12) {
          $element.classList.remove("error_input");
          formValues[id] = value;
        } else {
          $element.classList.add("error_input");
        }
        break;

      case "question":
        if (value !== "" && value.length > 1 && value.length <= 2500) {
          $element.classList.remove("error_input");
          formValues[id] = value;
        } else {
          $element.classList.add("error_input");
        }
        break;
    }
  });
});

document.querySelectorAll(`[data-type-user]`).forEach(($element) => {
  $element.addEventListener("click", (event) => {
    const { typeUser } = event.target.dataset;

    formValues.typeUser = typeUser;
  });
});

document.getElementById("sendForm").addEventListener("click", async () => {
  let error = false;

  Object.keys(formValues).forEach((key) => {
    if (key === "typeUser" || key === "secondName" || key === "thirdName") return;
    if (formValues[key] === "") {
      document.getElementById(key).classList.add("error_input");
      error = true;
      return;
    }
  });

  if (error) document.querySelector(".error-notif-inputs").style.display = null;

  if (!error) {
    document.querySelector(".error-notif-inputs").style.display = "none";

    await User.sendForm(formValues)
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          document.querySelector(".notif-error").style.display = "none";
          document.querySelector(".notif-successfully").style.display = null;

          inputs.forEach(($element) => {
            $element.value = "";
          });

          Object.keys(formValues).forEach((key) => {
            if (key === "typeUser") return;
            formValues[key] = "";
          });
        } else {
          document.querySelector(".notif-error").style.display = null;
          document.querySelector(".notif-successfully").style.display = "none";
        }
      });
  }
});
