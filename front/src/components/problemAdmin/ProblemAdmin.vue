<template>
  <div>
    <!-- model 开始 -->
    <b-modal
      id="addProblem"
      ref="addProblem"
      title="增加表单"
      size="xl"
      hide-footer
    >
      <b-form>
        <!-- 题目名称 -->
        <b-form-group label="题目名称">
          <b-form-input
            v-model="$v.problem.problemName.$model"
            required
            :state="validateState('problemName')"
            placeholder="请输入题目的名称"
          ></b-form-input>
          <b-form-invalid-feedback :state="validateState('problemName')">
            题目名称为必填项
          </b-form-invalid-feedback>
        </b-form-group>
        <!-- 题目描述 -->
        <b-form-group label="题目描述">
          <v-md-editor
            v-model="$v.problem.description.$model"
            height="800px"
            :state="validateState('description')"
          >
          </v-md-editor>
          <b-form-invalid-feedback :state="validateState('description')">
            题目描述为必填项
          </b-form-invalid-feedback>
        </b-form-group>
        <!-- 难易程度 -->
        <b-form-group label="难易程度" v-slot="{ ariaDescribedby }">
          <b-form-radio-group
            id="difficultyLevel"
            v-model="problem.difficultyLevel"
            :aria-describedby="ariaDescribedby"
            name="radio-sub-component"
          >
            <b-form-radio value="简单">简单</b-form-radio>
            <b-form-radio value="中等">中等</b-form-radio>
            <b-form-radio value="困难">困难</b-form-radio>
          </b-form-radio-group>
        </b-form-group>
        <!-- 时间限制 -->
        <b-form-group label="时间限制">
          <b-form-input
            type="number"
            v-model="$v.problem.timeLimit.$model"
            required
            :state="validateState('timeLimit')"
            placeholder="时间限制"
          ></b-form-input>
          <b-form-invalid-feedback :state="validateState('timeLimit')">
            时间限制为必填项
          </b-form-invalid-feedback>
        </b-form-group>
        <!-- 空间限制 -->
        <b-form-group label="空间限制">
          <b-form-input
            type="number"
            v-model="$v.problem.memoryLimit.$model"
            required
            :state="validateState('memoryLimit')"
            placeholder="空间限制"
          ></b-form-input>
          <b-form-invalid-feedback :state="validateState('memoryLimit')">
            空间限制为必填项
          </b-form-invalid-feedback>
        </b-form-group>
        <!-- 测试样例数目 -->
        <b-form-group label="测试样例数目">
          <b-form-input
            type="number"
            v-model="$v.problem.testCase.$model"
            required
            :state="validateState('testCase')"
            placeholder="请填入测试样例数目"
          ></b-form-input>
          <b-form-invalid-feedback :state="validateState('testCase')">
            测试样例数目为必填项
          </b-form-invalid-feedback>
        </b-form-group>
        <!-- 循环样例 -->
        <div v-for="count in parseInt(problem.testCase, 10)" v-bind:key="count">
          <h5 class="mt-3">第 {{ count }} 组测试样例</h5>
          <div class="d-flex justify-content-between align-items-center">
            <b-form-textarea
              id="textarea"
              placeholder="输入"
              rows="3"
              max-rows="6"
              v-model="store[count].input"
            ></b-form-textarea>
            <b-form-textarea
              id="textarea"
              placeholder="输出"
              rows="3"
              max-rows="6"
              v-model="store[count].output"
            ></b-form-textarea>
          </div>
        </div>
        <!-- 按钮 -->
        <b-form-group class="mt-3">
          <b-button variant="outline-info" block @click="createProblem"
            >创建</b-button
          >
        </b-form-group>
      </b-form>
    </b-modal>
    <!-- model 结束 -->

    <!-- 将题目添加到比赛的 model 结束 -->
    <div class="d-flex justify-content-between align-items-center mt-3 mb-3">
      <div></div>
      <b-button variant="outline-primary" size="lg" v-b-modal.addProblem
        >添加题目</b-button
      >
    </div>
    <b-table
      striped
      hover
      :items="problems"
      :fields="fields"
      class="text-secondary"
    >
      <!-- 对题目进行具体的操作 -->
      <template #cell(problem_id)="data">
        <div class="flex justify-content-between align-items-center">
          <!-- 显示删除 -->
          <b-button
            variant="danger"
            @click="problemDelete(data.value)"
            size="sm"
            class="ml-3"
            pill
          >
            删除
          </b-button>
        </div>
      </template>
    </b-table>
    <div class="overflow-auto mt-3">
      <b-pagination
        v-model="currentPage"
        :total-rows="total"
        :per-page="params.pageSize"
        align="right"
      ></b-pagination>
    </div>
  </div>
</template>

<script>
import { required } from 'vuelidate/lib/validators';
import problemService from '../../service/problemService';
// import contextService from '../../service/context';

export default {
  data() {
    return {
      // 获取题目列表
      problems: [],
      params: {
        pageNum: 1,
        pageSize: 10,
      },
      currentPage: 1,
      total: 1,
      fields: [
        { key: 'author', label: '作者' },
        { key: 'difficulty_level', label: '难度' },
        { key: 'problem_name', label: '题目描述' },
        { key: 'time_limit', label: '时间限制 （ms）' },
        { key: 'memory_limit', label: '内存限制（b）' },
        { key: 'problem_id', label: '' },
      ],
      testCases: [],
      store: [],
      // 关于添加题目的位置
      problem: {
        problemName: '', // 题目的名称
        description: '', // 题目描述
        difficultyLevel: '', // 难易程度
        timeLimit: 1000, // 时间限制
        memoryLimit: 67108864, // 空间限制
        testCase: '1', // 测试样例的数目
        author: 'admin',
        cases: [], // 所有的测试样例
      },
      // 关于删除题目
      deleteProblemParams: {
        problemID: 1,
      },
    };
  },
  // 数据校验部分
  validations: {
    problem: {
      problemName: {
        required,
      },
      description: {
        required,
      },
      testCase: {
        required,
      },
      timeLimit: {
        required,
      },
      memoryLimit: {
        required,
      },
    },
  },
  methods: {
    // 数据校验函数
    validateState(name) {
      // 这里是 es6 的解构赋值
      const { $dirty, $error } = this.$v.problem[name];
      return $dirty ? !$error : null;
    },

    async getProblemList() {
      const { data: res } = await problemService.getProblemList(this.params);
      this.problems = res.data.problems;

      this.total = res.data.total;
    },

    clearParams() {
      this.problem.problemName = '';
      this.problem.description = '';
      this.problem.difficultyLevel = '';
      this.problem.timeLimit = 0;
      this.problem.memoryLimit = 0;
      this.problem.testCase = '1';
      this.problem.author = 'admin';
      // eslint-disable-next-line dot-notation
      this.$refs['addProblem'].hide();
      this.$v.$reset();
    },

    async createProblem() {
      this.problem.timeLimit = parseInt(this.problem.timeLimit, 10);
      this.problem.memoryLimit = parseInt(this.problem.memoryLimit, 10);
      // 验证数据
      this.$v.problem.$touch();
      if (this.$v.problem.$anyError) {
        return;
      }
      // 添加题目
      for (let i = 0; i < parseInt(this.problem.testCase, 10); i += 1) {
        this.problem.cases.push(this.store[i + 1]);
      }

      for (let i = 0; i < 10; i += 1) {
        this.store[i] = {};
      }

      await problemService.createProblem(this.problem).then(() => {
        this.$bvToast.toast('创建题目成功', {
          title: '创建成功',
          variant: 'success',
          solid: true,
        });
        this.clearParams();
        this.getProblemList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '创建失败',
          variant: 'danger',
          solid: true,
        });
      });
    },

    async problemDelete(id) {
      this.deleteProblemParams.problemID = id;
      await problemService.deleteProblem(this.deleteProblemParams).then(() => {
        this.$bvToast.toast('删除题目成功', {
          title: '删除成功',
          variant: 'success',
          solid: true,
        });
        this.getProblemList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '删除失败',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },

  created() {
    this.getProblemList();
    for (let i = 0; i < 10; i += 1) {
      this.store[i] = {};
    }
  },

  watch: {
    currentPage() {
      this.params.pageNum = this.currentPage;
      this.getProblemList();
    },
  },
};
</script>
