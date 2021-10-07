// this is markdown need
import Prism from 'prismjs';
import 'prismjs/components/prism-json';

import Vue from 'vue';
// this is router need
import axios from 'axios';
import VueAxios from 'vue-axios';
// this is bootstrapVue need
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import Vuelidate from 'vuelidate';
// this is highlight need
import VueHighlightJS from 'vue-highlightjs';
// hybrid github monokai-sublime idea
import 'highlight.js/styles/hybrid.css';
// this is mark down neeed
import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
// eslint-disable-next-line import/extensions
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
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
// highlight code
Vue.use(VueHighlightJS);

// markdown 编辑器
VueMarkdownEditor.use(vuepressTheme, {
  Prism,
});

Vue.use(VueMarkdownEditor);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
