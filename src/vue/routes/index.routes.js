import { createRouter, createWebHistory } from "vue-router";

import Questions from "../pages/Questions.vue";
import SiteEdit from "../pages/SiteEdit.vue";

const routes = [
  {
    path: "/admin",
    component: Questions,
    name: "Все вопросы",
  },
  {
    path: "/admin/site-edit",
    component: SiteEdit,
    name: "Редактирование контента",
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  linkActiveClass: "active",
  linkExactActiveClass: "exact-active",
});

router.beforeEach((to) => {
  document.title = to.name;
});

export default router;
