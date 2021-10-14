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

// 创建比赛
// eslint-disable-next-line object-curly-newline
const createContext = ({ title, startTime, endTime, author }) => request.post('context', { title, startTime, endTime, author });

// 该题目是否在比赛里面
const problemInContext = ({ problemID, contextID }) => request.get('context_problem', { params: { problemID, contextID } });

// 将题目添加到比赛当中
const addProblemToContext = ({
  problemID,
  contextID,
  title,
}) => request.post('context_problem', { problemID, contextID, title });

// 将比赛中的题目删除
const deleteProblemInContext = ({ problemID, contextID }) => request.delete('context_problem', { params: { problemID, contextID } });

// 获取比赛中的题目列表
const getContextProblemList = ({ contextID, pageNum, pageSize }) => request.get('context_problem_list', { params: { contextID, pageNum, pageSize } });

export default {
  getContextList,
  getContext,
  deleteContext,
  updateContext,
  createContext,
  problemInContext,
  addProblemToContext,
  deleteProblemInContext,
  getContextProblemList,
};
