import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'home',
    component:() => import('../views/home.vue'),
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
  {
    path: '/header',
    name: 'header',
    component: () => import('../views/personCenter/header.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes
})

export default router
