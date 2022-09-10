import { createWebHistory, createRouter } from "vue-router";
import LoggedInPage from "@/views/LoggedInPage.vue";
import LoginPage from "@/views/LoginPage.vue";

const routes = [
  {
    path: "/",
    name: "Login",
    component: LoginPage,
  },
  {
    path: "/home",
    name: "Home",
    component: LoggedInPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;