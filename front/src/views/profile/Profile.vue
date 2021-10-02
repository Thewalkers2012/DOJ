<template>
  <div>
    <b-card class="mt-3">
      <div class="text-center">
        <b-avatar src="https://placekitten.com/300/300" size="16rem"></b-avatar>
        <h4 class="mt-3">{{ userInfo.username }}</h4>
      </div>
      <hr />
      <b-row>
        <b-col>
          <b-card title="Solved">
            <h5>已经通过：{{ solved }}</h5>
          </b-card>
        </b-col>
        <b-col>
          <b-card title="Submissions">
            <h5>提交总数： {{ submitCount }}</h5>
          </b-card>
        </b-col>
      </b-row>
    </b-card>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import userService from '../../service/userService';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  data() {
    return {
      items: [
        {
          text: '首页',
          to: { name: 'Home' },
        },
        {
          text: '个人主页',
          href: '#',
        },
      ],
      solved: '',
      submitCount: '',
    };
  },
  methods: {
    async getSolved() {
      const { data: res } = await userService.solved();
      this.solved = res.data.total;
    },
    async getSubmitCount() {
      const { data: res } = await userService.submissionCount();
      this.submitCount = res.data.total;
    },
  },
  created() {
    this.getSolved();
    this.getSubmitCount();
  },
};
</script>

<style lang="scss" scoped>
</style>
