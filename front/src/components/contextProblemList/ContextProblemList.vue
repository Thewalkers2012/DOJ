<template>
  <div>
    <h1>fuck</h1>
  </div>
</template>

<script>
import contextService from '../../service/context';

export default {
  data() {
    return {
      problems: [],
      problemListParams: {
        contextID: 1,
      },
    };
  },

  methods: {
    async getAllContextProblem() {
      this.problemListParams.contextID = parseInt(sessionStorage.getItem('context_id'), 10);

      const { data: res } = await contextService.getAllContextProblem(this.problemListParams);
      this.problems = res.data.problems;
    },
  },

  created() {
    this.getAllContextProblem();
  },
};
</script>
