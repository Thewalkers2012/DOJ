import request from '../utils/request';

// 用户注册
const register = ({ studentID, username, password }) => request.post('register', { studentID, username, password });

// 用户登录
const login = ({ studentID, password }) => request.post('login', { studentID, password });

// 获取用户信息
const info = () => request.get('info');

export default {
  register,
  info,
  login,
};
