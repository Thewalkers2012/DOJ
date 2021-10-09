import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '../store';
import Home from '../views/Home.vue';
import adminRoutes from './module/admin';
import problemRoutes from './module/problem';
import userRoutes from './module/user';
import ContextRoutes from './module/context';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue'),
  },
  ...userRoutes,
  ...problemRoutes,
  ...adminRoutes,
  ...ContextRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.auth) { // 判断是否需要登录
    // 用户是否登录
    if (store.state.userModule.token) {
      // 判读 token 的有效性，比如有没有过期，需要后台发放 token 的时候，带上 token 的有效期
      // 如果 token 无效需要请求 token
      next();
    } else {
      // 跳转登录
      router.push({ name: 'login' });
    }
  } else {
    next();
  }
});

export default router;
