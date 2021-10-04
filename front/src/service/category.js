import request from '../utils/request';

// 创建文章
const createCategory = ({ problemID, content }) => request.post('category', { problemID, content });

// 获取该问题的所有讨论
const getCategoryByProblem = ({ problemID }) => request.get('category', { params: { problemID } });

// 获取所有的文章
const getAllCategories = ({ pageNum, pageSize }) => request.get('category_list', { params: { pageNum, pageSize } });

// 获取单个文章的细节
const getCategoryDetails = ({ categoryID }) => request.get('category_details', { params: { categoryID } });

// 删除单个文章
const deleteCategory = ({ categoryID }) => request.delete('category', { params: { categoryID } });

export default {
  createCategory,
  getCategoryByProblem,
  getAllCategories,
  getCategoryDetails,
  deleteCategory,
};
