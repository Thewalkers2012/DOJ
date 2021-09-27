<template>
  <div class="ProblemList">
    <div class="mt-3 text-center">
      <h1>在线题库</h1>
      <Search class="mt-3"> </Search>

      <b-list-group class="mt-3">
        <b-list-group-item
          class="d-flex justify-content-between align-items-center"
        >
          <div class="d-flex justify-content-between align-items-center">
            <div class="mr-5">#</div>
            <div>题目名称</div>
          </div>
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
          <div class="d-flex justify-content-between align-items-center">
            <div class="mr-5">
              {{ item.problem_id }}
            </div>
            <div>
              {{ item.problem_name }}
            </div>
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
      <div class="overflow-auto mt-3">
        <b-pagination
          v-model="currentPage"
          :total-rows="total"
          :per-page="params.pageSize"
          align="right"
        ></b-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import problemService from '../service/problemService';
import Search from './search/Search.vue';

export default {
  data() {
    return {
      params: {
        pageNum: 1,
        pageSize: 10,
      },
      problemList: [],
      problem: {},
      total: 100,
      currentPage: 1,
    };
  },
  methods: {
    async getProblemList() {
      const { data: res } = await problemService.getProblemList(this.params);
      this.problemList = res.data.problems;
      this.total = res.data.total;
    },
    async getProblemByID(id) {
      sessionStorage.setItem('problem_id', id);
    },
  },
  created() {
    this.getProblemList();
  },
  components: {
    Search,
  },
  watch: {
    currentPage() {
      this.params.pageNum = this.currentPage;
      this.getProblemList();
    },
  },
};
</script>
