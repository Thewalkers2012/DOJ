import storeService from '../../service/storeService';
import problemService from '../../service/problemService';

const problemModule = {
  namespaced: true,
  state: {
    token: storeService.get(storeService.USER_TOKEN),
    userInfo: storeService.get(storeService.USER_INFO) ? JSON.parse(storeService.get(storeService.USER_INFO)) : null, // eslint-disable-line
  },

  actions: {
    getProblemList(context, { pageNum, pageSize }) {
      return new Promise((reject) => {
        problemService.getProblemList({ pageNum, pageSize }).then((res) => {
          console.log(1);
          return res.data.data.problems;
        }).catch((err) => {
          reject(err);
        });
      });
    },
  },
};

export default problemModule;
