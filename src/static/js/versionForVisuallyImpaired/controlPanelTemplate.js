const getTemplate = () => {
  return `
  <div>
    <p>Размер шрифта:</p>
    <a href="#" id="set-fontsize-v1">А</a>
    <a href="#" id="set-fontsize-v2">А</a>
    <a href="#" id="set-fontsize-v3">А</a>
    <a href="#" id="set-fontsize-v4">А</a>
  </div>
  <div>
    <p>Цветовая схема:</p>
    <a href="#" id="set-white-theme">А</a>
    <a href="#" id="set-black-theme">А</a>
    <a href="#" id="set-blue-theme">А</a>
  </div>
  <a href="#" id="default-theme">Обычная версия</a>
  `;
};

export default getTemplate;
