import { createApp } from 'vue'
import App from './App.vue'
import './assets/bootstrap.min.css'
import './assets/bootstrap.min.js'
import router from './router/index'
import './index.css'

const app = createApp(App)

app.use(router)

app.mount('#app')
