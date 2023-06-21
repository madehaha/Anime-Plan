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
  },
  // {
  //   path: '/header',
  //   name: 'header',
  //   component: () => import('../personCenter/Header.vue')
  // },
  {
    path: '/rightSide',
    name: 'rightSide',
    component: () => import('../personCenter/rightSide.vue')
  },
  {
    path: '/content',
    name: 'content',
    component: () => import('../personCenter/left-content.vue')
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes
})

export default router
