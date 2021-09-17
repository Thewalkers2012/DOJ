import Vue from 'vue';
import axios from 'axios';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import VueAxios from 'vue-axios';
import Vuelidate from 'vuelidate';
import VueHighlightJS from 'vue-highlightjs';
// hybrid github monokai-sublime idea
import 'highlight.js/styles/hybrid.css';
import App from './App.vue';
import router from './router';
import store from './store';
import '../app.scss';

Vue.config.productionTip = false;

Vue.use(Vuelidate);
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);
// axios
Vue.use(VueAxios, axios);

// 代码模块主题
Vue.use(VueHighlightJS);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
