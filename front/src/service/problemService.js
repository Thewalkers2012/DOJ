import request from '../utils/request';

// 用户注册
const getProblemList = ({ pageNum, pageSize }) => request.get('problem', { params: { pageNum, pageSize } });

// 用户登录
const getProblemByID = (id) => request.get(`problem/${id}`);

// 提交代码
const submitProblem = ({ language, code, questionID }) => request.post('submit', { language, code, questionID });

export default {
  getProblemByID,
  getProblemList,
  submitProblem,
};
