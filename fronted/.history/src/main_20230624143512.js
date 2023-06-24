import { createApp,ref } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import './assets/font/font.css'

const app = createApp(App)
const isadd = ref(false)
const radio = ref(1)

app.config.globalProperties.$isadd = isadd
app.config.globalProperties.$radio = radio

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
  app.use(store)
  app.use(router)
  app.use(ElementPlus)
app.mount('#app')

