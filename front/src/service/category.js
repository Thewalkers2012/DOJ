import request from '../utils/request';

// 创建文章
const createCategory = ({ problemID, content }) => request.post('category', { problemID, content });

// 获取该问题的所有讨论
const getCategoryByProblem = ({ problemID }) => request.get('category', { params: { problemID } });

export default {
  createCategory,
  getCategoryByProblem,
};
