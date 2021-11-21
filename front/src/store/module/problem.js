const problemModule = {
  namespace: true,
  state: {
    problemID: 0,
  },
  mutations: {
    SET_PROBLEM_ID(state, id) {
      // 更新 state
      state.problemID = id;
    },
    GET_PROBLEM_ID(state) {
      return state.problemID;
    },
  },
  actions: {
    setProblemID(context, id) {
      context.commit('SET_PROBLEM_ID', id);
    },
    getProblemID(context) {
      return context.commit('GET_PROBLEM_ID');
    },
  },
};

export default problemModule;
