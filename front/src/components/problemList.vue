<template>
  <div class="ProblemList">
    <div class="mt-3">
      <h1>这里是 problemList 组件</h1>
      <hr />
      <b-list-group>
        <b-list-group-item
          class="d-flex justify-content-between align-items-center"
        >
          <div>题目名称</div>
          <div class="d-flex">
            <div>通过率</div>
            <div class="ml-5">难度</div>
          </div>
        </b-list-group-item>
        <b-list-group-item
          v-for="item in problemList"
          v-bind:key="item.id"
          button
          class="d-flex justify-content-between align-items-center"
          @click="
            getProblemByID(item.problem_id);
            $router.push(`/problem/${item.problem_id}`);
          "
        >
          <div>
            {{ item.problem_name }}
          </div>
          <div class="d-flex">
            <div>100%</div>
            <b-badge
              variant="success"
              pill
              v-if="item.difficulty_level === '简单'"
              class="ml-5"
              >简单
            </b-badge>
            <b-badge
              variant="warning"
              pill
              v-else-if="item.difficulty_level === '中等'"
              class="ml-5"
              >中等</b-badge
            >
            <b-badge variant="danger" pill v-else class="ml-5">困难</b-badge>
          </div>
        </b-list-group-item>
      </b-list-group>
    </div>
  </div>
</template>

<script>
import problemService from '../service/problemService';

export default {
  data() {
    return {
      params: {
        pageNum: 1,
        pageSize: 5,
      },
      problemList: [],
      problem: {},
    };
  },
  methods: {
    async getProblemList() {
      const { data: res } = await problemService.getProblemList(this.params);
      this.problemList = res.data.problems;
    },
    async getProblemByID(id) {
      sessionStorage.setItem('problem_id', id);
    },
  },
  created() {
    this.getProblemList();
  },
};
</script>
