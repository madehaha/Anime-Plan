import { createRouter, createWebHistory } from 'vue-router'
import login from '../views/login.vue'

const routes = [
  {
    path: '/',
    name: 'login',
    component: () => import('../views/login.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
