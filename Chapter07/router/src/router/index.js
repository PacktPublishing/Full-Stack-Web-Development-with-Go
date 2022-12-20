import Vue from 'vue';
// import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Login from "../views/Login.vue";
import { createRouter, createWebHashHistory } from 'vue-router'

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    base: process.env.BASE_URL,
    routes
})

export default router
