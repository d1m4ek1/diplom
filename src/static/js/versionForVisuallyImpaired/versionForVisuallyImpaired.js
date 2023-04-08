import getTemplate from "./controlPanelTemplate";
import setFontSize from "./fontSizes";
import { changeTheme } from "./themes";

const data = {
  fontSize: {
    v1: "16px",
    v2: "20px",
    v3: "24px",
    v4: "28px",
  },
  settedFontSize: "default",
  settedTheme: "default",
  mainBackgroundImage: document.body.style.backgroundImage,
};

class VersionForVisuallyImpaired {
  constructor(selector) {
    this.$renderBlock = document.querySelector(selector);

    this.$links = null;
  }

  render() {
    this.$renderBlock.classList.add("theme_control_panel");
    this.$renderBlock.innerHTML = getTemplate();

    this.$links = {
      colorsTheme: {
        default: document.getElementById("default-theme"),
        white: document.getElementById("set-white-theme"),
        black: document.getElementById("set-black-theme"),
        blue: document.getElementById("set-blue-theme"),
      },
      fontSize: {
        v1: document.getElementById("set-fontsize-v1"),
        v2: document.getElementById("set-fontsize-v2"),
        v3: document.getElementById("set-fontsize-v3"),
        v4: document.getElementById("set-fontsize-v4"),
      },
      main: document.getElementById("version_for_visually-impaired"),
    };

    this.#mainHandler();

    this.#fontSizeHandlers();

    this.#colorThemeHandlers();
  }

  #mainHandler() {
    this.$links.main.addEventListener("click", (e) => {
      e.preventDefault();

      if (data.settedTheme === "default") {
        data.settedTheme = changeTheme("white");
        document.querySelector(".burger-btn").click();

        this.$renderBlock.style.display = "flex";
      } else {
        data.settedTheme = changeTheme("default");
        this.$renderBlock.style.display = "none";
      }
    });
  }

  #colorThemeHandlers() {
    Object.keys(this.$links.colorsTheme).forEach((color) => {
      this.$links.colorsTheme[color].addEventListener("click", () => {
        if (color === "default") {
          this.$renderBlock.style.display = "none";
        }

        data.settedTheme = changeTheme(color);
      });
    });
  }

  #fontSizeHandlers() {
    Object.keys(this.$links.fontSize).forEach((key) => {
      this.$links.fontSize[key].addEventListener("click", () => {
        setFontSize(data.fontSize[key]);
      });
    });
  }
}

export default VersionForVisuallyImpaired;
