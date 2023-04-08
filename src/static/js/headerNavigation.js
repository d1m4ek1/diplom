(function () {
  const $button = document.querySelector(".burger-btn");
  const $navigation = document.getElementById("navigation");

  $button.addEventListener("click", (event) => {
    isShowHidenBlock($navigation);
  });

  $navigation.addEventListener("click", clickHandler);

  function clickHandler(event) {
    const { id } = event.target.dataset;
    if (id) {
      const $findedList = $navigation.querySelector(`[data-for="${id}"]`);

      event.preventDefault();

      isShowHidenBlock($findedList);
    }
  }

  function isShowHidenBlock($block) {
    if (!$block.style.display) {
      $block.style.display = "flex";
    } else {
      $block.style.display = null;
    }
  }
})();
