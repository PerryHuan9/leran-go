import { createRouter, createWebHistory } from 'vue-router'

import HomePage from './pages/Home.vue'
import LoginPage from './pages/Login.vue'

const routes = [
  { path: '/', component: HomePage },
  { path: '/login', component: LoginPage },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})