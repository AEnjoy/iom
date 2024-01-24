import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import axios from './api/axiosInstance.js'
import VForm3 from 'vform3-builds'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'vform3-builds/dist/designer.style.css'

const app = createApp(App)

app.use(ElementPlus)
app.use(router)
app.use(VForm3)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.mount('#app')
app.config.globalProperties.$axios=axios;