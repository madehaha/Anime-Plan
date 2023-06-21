import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/',
    name: 'home',
    component:() => import('../views/home.vue'),
    redirect: '/'
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
