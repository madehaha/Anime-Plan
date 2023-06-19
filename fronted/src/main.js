import { createApp } from 'vue'

import App from './App.vue'

import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'


// 使用封装的登录验证vue组件


createApp(App).use(store).use(router).use(ElementPlus).mount('#app')
