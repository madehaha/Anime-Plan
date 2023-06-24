import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/',
    redirect: '/MainContentView'
  },
  {
    path: '/MainContentView',
    name: 'MainContentView',
    component:() => import('../views/MainContentView'),
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
  {
    path: '/rightSide',
    name: 'rightSide',
    component: () => import('../personCenter/rightSide.vue')
  },
  // {
  //   path: '/header',
  //   name: 'Header',
  //   component: () => import('../personCenter/header.vue')
  // },
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
