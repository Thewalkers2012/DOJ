// 本地缓存服务

const PREFIX = 'DOJ_';

// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;
// const PROBLEM_PREFIX = `${PREFIX}problem_`;
// const PROBLEM_ID = `${PROBLEM_PREFIX}id`;

// 从 vuex 中获取用户的信息
// 储存
const set = (key, data) => {
  localStorage.setItem(key, data);
};

// 读取
const get = (key) => localStorage.getItem(key);

// 从 vuex 中获取题目信息
// setP 函数将题目的 id 保存到 localStorage 中
// const setP = (key, data) => {
//   localStorage.setItem(key, data);
// };

// const getP = (key, data) => {
//   localStorage.getItem(key, data);
// };

export default {
  set,
  get,
  // setP,
  // getP,
  USER_TOKEN,
  USER_INFO,
  // PROBLEM_ID,
};
