<template>
  <div>
    <!-- 题目描述区域 -->
    <div>
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
    <!-- 提交代码部分 -->
    <div class="d-flex mt-3 justify-content-between align-items-center">
      <div></div>
      <div>
        <b-button variant="success" pill class="ml-3" @click="submitProblem">
          提交代码
        </b-button>
      </div>
    </div>
    <!-- 描述部分结束 -->
  </div>
</template>

<script>
import problemService from '../../service/problemService';
import Edit from '../edit/Edit.vue';
// import getSubmitMsg from '../../helper/getAnswer';

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
      submitParams: {
        code: '',
        language: '',
        questionID: 0,
      },
      id: '',
      submitMsg: '',
    };
  },
  created() {
    const id = sessionStorage.getItem('problem_id');
    this.id = id;
    this.getProblemByID(id);
  },
  methods: {
    async getProblemByID(id) {
      const { data: res } = await problemService.getProblemByID(id);
      this.problem = res.data.problem;
    },
    async submitProblem() {
      this.submitParams.questionID = parseInt(sessionStorage.getItem('problem_id'), 10);
      this.submitParams.language = this.selected1;
      this.submitParams.code = localStorage.getItem(`code_${this.submitParams.questionID}`);
      const { data: res } = await problemService.submitProblem(this.submitParams);

      if (res.data.data.answer_code === 0) {
        this.submitMsg = 'Accept';
      } else if (res.data.data.answer_code === -1) {
        this.submitMsg = 'Wrong Answer';
      } else if (res.data.data.answer_code === 1) {
        this.submitMsg = 'Time Limit Exceeded';
      } else if (res.data.data.answer_code === 4) {
        this.submitMsg = 'RunTime Error';
      } else {
        this.submitMsg = 'Compile Error';
      }

      console.log(this.submitMsg);
    },
  },
  components: {
    Edit,
  },
};
</script>
