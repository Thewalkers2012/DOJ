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
        <b-tab title="题目"></b-tab>
        <b-tab title="排行"></b-tab>
        <b-tab title="提交"></b-tab>
      </b-tabs>
    </b-card>
  </div>
</template>

<script>
import contextService from '../../service/context';

export default {
  data() {
    return {
      context: {},
      id: '',
    };
  },

  methods: {
    async getContext() {
      this.id = sessionStorage.getItem('context_id');
      const { data: res } = await contextService.getContext(this.id);
      this.context = res.data.context;
      console.log(res.data.context);
    },
  },

  created() {
    this.getContext();
  },
};
</script>
