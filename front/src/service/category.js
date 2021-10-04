import request from '../utils/request';

// 创建文章
const createCategory = ({ problemID, content }) => request.post('category', { problemID, content });

// 获取该问题的所有讨论
const getCategoryByProblem = ({ problemID }) => request.get('category', { params: { problemID } });

// 获取所有的文章
const getAllCategories = ({ pageNum, pageSize }) => request.get('category_list', { params: { pageNum, pageSize } });

export default {
  createCategory,
  getCategoryByProblem,
  getAllCategories,
};
