import { createApp } from 'vue'
import App from './App.vue'
import './assets/bootstrap.min.css'
import './assets/bootstrap.min.js'
import router from './router/index'
import axios from 'axios'
import VueAxios from 'vue-axios'

const app = createApp(App)
// vue router
app.use(router)
// axios 
app.use(VueAxios, axios)
// 挂载 app 节点
app.mount('#app')
