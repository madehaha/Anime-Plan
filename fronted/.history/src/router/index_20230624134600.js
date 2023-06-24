import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  // {
    // path: '/',
    // name: 'home',
    // component: HomeView,
    // redirect:"/contentView",
    // children:[
      {
        path:"/",
        redirect:"/contentView",
      },
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
      },
      {
        path: '/detailview',
        name: 'detailview',
        
        component: () => import( '../components/HomeViewComponents/DetailView.vue')
      },
      
    // ]
  // },
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
  {
    //注册界面的路由
    path:'/register',
    name:'register',
    component:()=>import('../views/RegisterView.vue')

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
