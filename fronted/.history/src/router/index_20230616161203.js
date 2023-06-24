import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/home',
    name: 'home',
    component:() => import('../views/home.vue'),
    redirect: '/home'
  },
  {
    path: '/login',
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
  routes: routes
})

export default router
