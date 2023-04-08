(function () {
  const $arrow = document.getElementById("arrow_scroll_up");
  $arrow.classList.add("fa", "arrow_scroll_up");

  window.addEventListener("scroll", () => {
    if (window.scrollY) {
      $arrow.style.display = "flex";

      if (!$arrow.classList.contains("arrow_scroll_up__showed")) {
        $arrow.classList.add("arrow_scroll_up__showed");
        $arrow.classList.remove("arrow_scroll_up__hidden");
      }
    } else {
      if (!$arrow.classList.contains("arrow_scroll_up__hidden")) {
        $arrow.classList.add("arrow_scroll_up__hidden");
        $arrow.classList.remove("arrow_scroll_up__showed");
      }
      setTimeout(() => {
        $arrow.style.display = null;
      }, 300);
    }
  });

  $arrow.addEventListener("click", () => {
    window.scrollTo({
      top: 0,
      left: 0,
      behavior: "smooth",
    });
  });
})();
