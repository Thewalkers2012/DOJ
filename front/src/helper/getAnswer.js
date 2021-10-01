function getSubmitMsg(code) {
  if (code === 0) {
    return 'Accept';
  }
  if (code === -1) {
    return 'Wrong Answer';
  }
  if (code === 1) {
    return 'Time Limit Exceeded';
  }
  if (code === 4) {
    return 'RunTime Error';
  }
  return 'Compile Error';
}

export default {
  getSubmitMsg,
};
