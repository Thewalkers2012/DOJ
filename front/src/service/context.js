import request from '../utils/request';

const getContextList = ({ pageNum, pageSize }) => request.get('context_list', { params: { pageNum, pageSize } });

export default {
  getContextList,
};
