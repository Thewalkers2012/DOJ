<template>
  <div class="mt-3">
    <b-card
      bg-variant="light"
      text-variant="dark"
      :header="context.title"
      class="text-center"
    >
      <b-card-text>
        <div>比赛时间：{{ context.start_time }} ~ {{ context.end_time }}</div>
        <div class="mt-3">发起人：{{ context.author }}</div>
        <div class="mt-3">
          <b-button
            variant="primary"
            v-if="context.defunct === '1' || context.defunct === '2'"
          >
            点击报名
          </b-button>
          <b-button variant="primary" v-else disabled>比赛已经结束</b-button>
        </div>
      </b-card-text>
      <hr />
      <b-tabs content-class="mt-5" fill text-variant="dark">
        <b-tab title="题目" @click="changeToProblem"></b-tab>
        <b-tab title="排行" @click="changeToRank"></b-tab>
        <b-tab title="提交" @click="changeToSubmit"></b-tab>
      </b-tabs>

      <div class="mt-3">
        <ContextProblemList v-if="isProblems"></ContextProblemList>
      </div>
    </b-card>
  </div>
</template>

<script>
import contextService from '../../service/context';
import ContextProblemList from '../contextProblemList/ContextProblemList.vue';

export default {
  data() {
    return {
      context: {},
      id: '',
      isProblems: true,
      isRank: false,
      isSubmit: false,
    };
  },

  methods: {
    async getContext() {
      this.id = sessionStorage.getItem('context_id');
      const { data: res } = await contextService.getContext(this.id);
      this.context = res.data.context;
    },

    changeToProblem() {
      this.isProblems = true;
      this.isRank = false;
      this.isSubmit = false;
    },

    changeToRank() {
      this.isRank = true;
      this.isProblems = false;
      this.isSubmit = false;
    },

    changeToSubmit() {
      this.isSubmit = true;
      this.isProblems = false;
      this.isRank = false;
    },
  },

  created() {
    this.getContext();
  },

  components: {
    ContextProblemList,
  },
};
</script>
