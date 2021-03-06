<template>
  <div>
    <!-- 题目描述区域 -->
    <div class="mt-3">
      <v-md-editor :value="problem.description" mode="preview"></v-md-editor>
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
      <!-- 编辑器组件 -->
      <Edit class="mt-3" :propLanguage="selected1" :propTheme="selected2">
      </Edit>
      <!-- <pre v-highlightjs><code class="cpp">{{ code }}</code></pre> -->
    </b-card>
    <!-- 提交代码部分 -->
    <div class="d-flex mt-3 justify-content-between align-items-center">
      <div></div>
      <div>
        <b-button
          variant="success"
          pill
          class="ml-3"
          @click="submitProblem"
          v-bind:disabled="isSubmit"
          size="lg"
        >
          提交代码
        </b-button>
      </div>
    </div>
    <!-- 描述部分结束 -->
    <!-- 点击 Submit 按钮，然后会显示提交代码显示的信息 -->
    <b-card class="mt-3" v-if="showMsg" border-variant="Secondary">
      <div class="d-flex align-items-center">
        <!-- 代码提交的状态 -->
        <h3>代码提交状态：</h3>
        <h3 class="ml-1" :style="{ color: color }">{{ submitMsg }}</h3>
        <b-spinner
          class="float-right ml-3"
          variant="primary"
          v-if="submitMsg === 'Judging'"
        ></b-spinner>
      </div>
      <!-- 编译错误的提示信息 -->
      <b-alert
        show
        variant="danger"
        class="mt-3"
        v-if="submitMsg === 'Compile Error'"
      >
        <h5>错误信息</h5>
        <pre>{{ compileMsg }}</pre>
      </b-alert>
    </b-card>
  </div>
</template>

<script>
import problemService from '../../service/problemService';
import Edit from '../edit/Edit.vue';

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
        problemID: 0,
      },
      id: '',
      submitMsg: '',
      isSubmit: false, // 控制提交按钮的 disable 属性
      color: '', // 显示 Message 的显示
      finish: false,
      showMsg: false, // 控制显示信息
      compileMsg: '', // 显示编译错误信息
    };
  },
  created() {
    const id = sessionStorage.getItem('problem_id');
    this.id = id;
    this.getProblemByID(id);
  },
  methods: {
    startSubmit() {
      this.color = 'darkblue';
      this.isSubmit = true;
      this.submitMsg = 'Judging';
      this.finish = false;
      this.showMsg = true;
      this.submitParams.problemID = parseInt(sessionStorage.getItem('problem_id'), 10);
      this.submitParams.language = this.selected1;
      this.submitParams.code = localStorage.getItem(`code_${this.submitParams.problemID}`);
    },
    endSubmit(codeValue, msg) {
      // 提交的信息的颜色
      if (codeValue === 0) {
        this.color = 'darkgreen';
      } else {
        this.color = 'darkred';
      }

      // 显示提交的信息
      if (codeValue === 0) {
        this.submitMsg = 'Accept';
      } else if (codeValue === -1) {
        this.submitMsg = 'Wrong Answer';
      } else if (codeValue === 1) {
        this.submitMsg = 'Time Limit Exceeded';
      } else if (codeValue === 4) {
        this.submitMsg = 'RunTime Error';
      } else if (codeValue === 3) {
        this.submitMsg = 'Memory Limit Exceeded';
      } else {
        this.submitMsg = 'Compile Error';
        this.compileMsg = msg;
      }

      this.finish = true;
      this.isSubmit = false;
    },
    async getProblemByID(id) {
      const { data: res } = await problemService.getProblemByID(id);
      this.problem = res.data.problem;
    },
    async submitProblem() {
      this.startSubmit();
      const { data: res } = await problemService.submitProblem(this.submitParams);
      this.endSubmit(res.data.data.answer_code, res.msg);
    },
  },
  components: {
    Edit,
  },
};
</script>
