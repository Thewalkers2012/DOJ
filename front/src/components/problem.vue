<template>
  <div class="Problem">
    <b-card class="mt-3">
      <h2 class="mt-3" style="font-weight: bold">
        {{ problem.problem_id }}. {{ problem.problem_name }}
      </h2>
      <div class="mt-3">
        <b-nav tabs fill style="color: black">
          <b-nav-item>
            <div style="color: black">题目</div>
          </b-nav-item>
          <b-nav-item>
            <div style="color: black">提交记录</div>
          </b-nav-item>
          <b-nav-item>
            <div style="color: black">讨论</div>
          </b-nav-item>
        </b-nav>
      </div>
      <div class="mt-3">
        <h3 style="font-weight: bold">题目描述：</h3>
        <div style="font-size: 20px" class="mt-3">
          {{ problem.description }}
        </div>
      </div>
      <div class="mt-3">
        <h3 style="font-weight: bold">输入格式：</h3>
        <div style="font-size: 20px" class="mt-3">
          {{ problem.test_case }}
        </div>
      </div>
      <div style="font-size: 20px" class="mt-3">
        <h3 style="font-weight: bold">输出格式：</h3>
        {{ problem.test_case }}
      </div>
      <!-- 代码编辑区域 -->
      <b-card class="mt-5" title="编辑器">
        <Edit> </Edit>
        <!-- <pre v-highlightjs><code class="cpp">{{ code }}</code></pre> -->
      </b-card>
    </b-card>
  </div>
</template>

<script>
import problemService from '../service/problemService';
import Edit from './edit/Edit.vue';

export default {
  data() {
    return {
      problem: {},
      code: `#include <bits/stdc++.h>
using namespace std;

int main() {
  int a, b;
  cin >> a >> b;
  cout << a + b << endl;
  return 0;
}`,
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
  components: {
    Edit,
  },
};
</script>
