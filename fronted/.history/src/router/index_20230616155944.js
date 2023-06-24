import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/',
    name: 'login',
    component: () => import('../views/login.vue')
  },
  {
    path: '/center',
    name: 'center',
    component: () => import('../views/personCenter.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
