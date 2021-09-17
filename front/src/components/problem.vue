<template>
  <div class="Problem">
    <h1>题目名称</h1>
    <h2>{{ problem.problem_name }}</h2>
  </div>
</template>

<script>
import problemService from '../service/problemService';

export default {
  data() {
    return {
      problem: {},
    };
  },
  created() {
    const id = sessionStorage.getItem('problem_id');
    this.getProblemByID(id);
  },
  methods: {
    async getProblemByID(id) {
      const { data: res } = await problemService.getProblemByID(id);
      this.problem = res.data.problem;
    },
  },
};
</script>
