import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    redirect:"/contentView",
    children:[
      {
        //主页面的路由
        path:'/contentView',
        name:'Main',
        component:()=>import('../components/HomeViewComponents/MainContentView.vue'),
        
      },
      {
        //主页面的路由
        path:'/RankingsView',
        name:'rank',
        component:()=>import('../components/HomeViewComponents/RankingsView.vue'),

      },
      {
        path:'/DayPushView',
        name:'dayPush',
        component:()=>import('../components/HomeViewComponents/DayPushView.vue'),
      }
    ]
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    //注册界面的路由
    path:'/register',
    name:'register',
    component:()=>import('../views/RegisterView.vue')

  },
  {
    //登录界面的路由
    path:'/login',
    name:'login',
    component:()=>import('../views/LoginView.vue')

  },
  
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
