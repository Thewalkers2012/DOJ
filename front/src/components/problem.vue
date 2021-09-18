/* eslint-disable max-len */
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
      <b-card class="mt-5">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h4>编辑器</h4>
          </div>
          <div class="d-flex">
            <b-form-select
              v-model="selected2"
              :options="options2"
              class="ml-3"
            ></b-form-select>
            <b-form-select
              class="ml-3"
              v-model="selected1"
              :options="options1"
            ></b-form-select>
          </div>
        </div>
        <Edit class="mt-3" :propLanguage="selected1" :propTheme="selected2">
        </Edit>
        <!-- <pre v-highlightjs><code class="cpp">{{ code }}</code></pre> -->
      </b-card>
      <div class="d-flex mt-3 justify-content-between align-items-center">
        <div></div>
        <div>
          <b-button variant="light" pill>调试代码</b-button>
          <b-button variant="success" pill class="ml-3"> 提交代码 </b-button>
        </div>
      </div>
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
      selected1: 'cpp',
      options1: [
        { value: 'cpp', text: 'c++' },
        { value: 'java', text: 'java' },
        { value: 'python', text: 'python' },
      ],
      selected2: 'vs-dark',
      options2: [
        { value: 'vs-dark', text: 'vs-dark' },
        { value: 'vs-light', text: 'vs-light' },
      ],
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
