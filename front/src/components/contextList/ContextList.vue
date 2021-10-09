<template>
  <div class="mt-3">
    <h1>这里是比赛组件</h1>
  </div>
</template>

<script>
import contextService from '../../service/context';

export default {
  data() {
    return {
      params: {
        pageNum: 1,
        pageSize: 10,
      },
      contexts: [],
    };
  },

  methods: {
    async getContextList() {
      const { data: res } = await contextService.getContextList(this.params);
      console.log(res);
    },
  },

  created() {
    this.getContextList();
  },
};
</script>
