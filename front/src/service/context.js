import request from '../utils/request';

// 获取比赛列表
const getContextList = ({ pageNum, pageSize }) => request.get('context_list', { params: { pageNum, pageSize } });

// 获取比赛
const getContext = (id) => request.get(`context/${id}`);

export default {
  getContextList,
  getContext,
};
