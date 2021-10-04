import request from '../utils/request';

// 用户列表
const userList = ({ pageSize, pageNum }) => request.get('user', { params: { pageNum, pageSize } });

// 用户注册
const register = ({ studentID, username, password }) => request.post('register', { studentID, username, password });

// 用户登录
const login = ({ studentID, password }) => request.post('login', { studentID, password });

// 获取用户信息
const info = () => request.get('info');

// 获取 AC 题目数量
const solved = () => request.get('solved');

// 获取提交总数
const submissionCount = () => request.get('submit_count');

// 修改角色信息
const updateUser = ({ userID, username, studentID }) => request.put('user', { userID, username, studentID });

export default {
  register,
  info,
  login,
  solved,
  submissionCount,
  userList,
  updateUser,
};
