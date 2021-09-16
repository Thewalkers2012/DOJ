<template>
  <div class="Problem">
    <div class="mt-3">
      <h1>这里 是 problem 组件</h1>
      <hr />
      <b-list-group> </b-list-group>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      params: {
        pageNum: 1,
        pageSize: 5,
      },
      problemList: [],
    };
  },
  methods: {
    ...mapActions('problemModule', { getProblemList: 'getProblemList' }),
    getList() {
      this.problemList = this.getProblemList(this.params);
    },
  },
  created() {
    this.getList();
  },
};
</script>
