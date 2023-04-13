import Admin from "../models/admin";
import CSRF from "../models/csrf";

const $login = document.getElementById("login");
const $password = document.getElementById("password");
const $button = document.getElementById("submit");

const authAdmin = async () => {
  let token = "";
  const login = $login.value;
  const password = $password.value;

  if (login !== "") {
    if (password !== "") {
      await CSRF.GetToken().then((response) => (token = response.headers.get("X-CSRF-TOKEN")));

      if (token !== "") {
        await new Admin(token)
          .AdminLogin({
            login,
            password,
          })
          .then((response) => response.json())
          .then((response) => {
            if (response.successfully) {
              if (response.isLogin) {
                location.href = "/admin";
              }
            }
          });
      }
    }
  }
};

$button.addEventListener("click", authAdmin);
