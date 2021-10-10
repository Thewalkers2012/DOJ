import request from '../utils/request';

// 获取比赛列表
const getContextList = ({ pageNum, pageSize }) => request.get('context_list', { params: { pageNum, pageSize } });

// 获取比赛
const getContext = (id) => request.get(`context/${id}`);

// 删除比赛
const deleteContext = ({ contextID }) => request.delete('context', { params: { contextID } });

// 更新比赛
// eslint-disable-next-line object-curly-newline
const updateContext = ({ contextID, title, startTime, endTime, author }) => request.put('context', { contextID, title, startTime, endTime, author });

export default {
  getContextList,
  getContext,
  deleteContext,
  updateContext,
};
