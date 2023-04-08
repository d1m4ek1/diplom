import setFontSize from "./fontSizes";

const changedTheme = {
  white: {
    colorFont: "#000",
    background: "#fff",
  },
  black: {
    colorFont: "#fff",
    background: "#000",
  },
  blue: {
    colorFont: "#063462",
    background: "#9dd1ff",
  },
  default: {
    colorFont: "#000",
    background: "#fff",
  },
  mainBackgroundImage: document.body.style.backgroundImage,
};

const changeBackground = [
  document.querySelectorAll(".dark_blue"),
  document.querySelectorAll(".dark_light_blue"),
  document.querySelectorAll(".header"),
  document.querySelectorAll(".content"),
  document.querySelectorAll(".aside_content__section"),
  document.querySelectorAll(".navigation__list_items > li"),
];

const changeColorText = [
  document.querySelectorAll(".navigation__list_item a"),
  document.querySelectorAll(".dark_blue"),
  document.querySelectorAll(".wrapper h1"),
  document.querySelectorAll(".wrapper h2"),
  document.querySelectorAll(".wrapper p"),
  document.querySelectorAll(".wrapper ul"),
];

const enumerationOfArrays = (array, color, keyChange) => {
  for (let i = 0; i < array.length; i++) {
    const element = array[i];

    element.forEach((item) => {
      item.style[keyChange] = color;
    });
  }
};

const addOrRemoveClassToBlock = (className, isAdd) => {
  const addClassNotHover = [document.querySelectorAll(".sub_navigation__list_items > li")];

  for (let i = 0; i < addClassNotHover.length; i++) {
    const element = addClassNotHover[i];

    element.forEach((item) => {
      if (isAdd) item.classList.add(className);
      else item.classList.remove(className);
    });
  }
};

export const changeTheme = (color) => {
  const $buttonVisuallyImpaired = document.getElementById("version_for_visually-impaired");

  const background = color === "default" ? null : changedTheme[color].background;
  const colorFont = color === "default" ? null : changedTheme[color].colorFont;

  enumerationOfArrays(changeBackground, background, "backgroundColor");
  enumerationOfArrays(changeColorText, colorFont, "color");

  addOrRemoveClassToBlock("not_hover", true);

  $buttonVisuallyImpaired.innerText = color === "default" ? "Версия для слабовидящих" : "Обычная версия";

  document.body.style.color = changedTheme[color].colorFont;
  document.body.style.backgroundColor = background;
  document.body.style.backgroundImage = color === "default" ? changedTheme.mainBackgroundImage : "unset";

  if (color === "default") {
    setFontSize("default");
  }

  return color;
};
