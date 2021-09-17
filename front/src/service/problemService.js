import request from '../utils/request';

// 用户注册
const getProblemList = ({ pageNum, pageSize }) => request.get('problem', { params: { pageNum, pageSize } });

// 用户登录
const getProblemByID = (id) => request.get(`problem/${id}`);

export default {
  getProblemByID,
  getProblemList,
};
