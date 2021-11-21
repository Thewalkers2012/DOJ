<template>
  <div>
    <b-list-group class="mt-3">
      <b-list-group-item
        class="d-flex justify-content-between align-items-center"
      >
        <strong>题目</strong>
        <strong>通过率</strong>
      </b-list-group-item>
      <b-list-group-item
        button
        v-for="item in problems"
        v-bind:key="item.problem_id"
        class="d-flex justify-content-between align-items-center"
        variant="default"
        @click="ToProblem(item.problem_id)"
      >
        <div>{{ item.problem_name }}</div>
        <div>100/500</div>
      </b-list-group-item>
    </b-list-group>
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
    ToProblem(id) {
      sessionStorage.setItem('problem_id', id);
      console.log(id);
    },
  },

  created() {
    this.getAllContextProblem();
  },
};
</script>
