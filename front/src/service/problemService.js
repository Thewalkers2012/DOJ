import request from '../utils/request';

// 获取题目列表
const getProblemList = ({ pageNum, pageSize }) => request.get('problem', { params: { pageNum, pageSize } });

// 通过编号获取题目
const getProblemByID = (id) => request.get(`problem/${id}`);

// 提交代码
const submitProblem = ({ language, code, problemID }) => request.post('submit', { language, code, problemID });

// 获取提交记录
const getSubmissions = ({ problemID }) => request.get('submission', { params: { problemID } });

// 创建题目
const createProblem = ({
  problemName, description, testCase, author, timeLimit, memoryLimit, difficultyLevel,
}) => request.post(
  'problem',
  {
    problemName,
    description,
    testCase,
    author,
    timeLimit,
    memoryLimit,
    difficultyLevel,
  },
);

export default {
  getProblemByID,
  getProblemList,
  submitProblem,
  getSubmissions,
  createProblem,
};
