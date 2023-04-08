const texts = [
  document.querySelectorAll("header a"),
  document.querySelectorAll(".sub_navigation__list_items > li"),
  document.querySelectorAll(".wrapper h1"),
  document.querySelectorAll(".wrapper h2"),
  document.querySelectorAll(".wrapper p"),
  document.querySelectorAll(".wrapper ul"),
];

const setFontSize = (fontSize) => {
  for (let i = 0; i < texts.length; i++) {
    const element = texts[i];

    element.forEach((item) => {
      item.style.fontSize = fontSize === "default" ? null : fontSize;
    });
  }
};

export default setFontSize;
