/* eslint-disable no-await-in-loop */
<template>
  <div>
    <div class="d-flex justify-content-between align-items-center">
      <div></div>
      <b-button variant="outline-primary" size="lg" v-b-modal.createContext>
        创建比赛
      </b-button>
    </div>

    <!-- 创建比赛开始 -->
    <b-modal
      id="createContext"
      ref="createContext"
      title="创建表单"
      hide-footer
    >
      <b-form-group label="比赛名称">
        <b-form-input
          v-model="$v.createParams.title.$model"
          required
          :state="validateState2('title')"
          placeholder="请输入比赛名称"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState2('title')">
          比赛名称为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="发起人">
        <b-form-input
          v-model="$v.createParams.author.$model"
          required
          :state="validateState2('author')"
          placeholder="请输入比赛的发起人"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState2('author')">
          比赛的发起人为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="开始时间">
        <b-form-input
          v-model="$v.createParams.startTime.$model"
          required
          :state="validateState2('startTime')"
          placeholder="时间格式：xx-xx-xx xx:xx:xx"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState2('startTime')">
          比赛的开始时间为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="结束时间">
        <b-form-input
          v-model="$v.createParams.endTime.$model"
          required
          :state="validateState2('endTime')"
          placeholder="时间格式：xx-xx-xx xx:xx:xx"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState2('endTime')">
          比赛的结束时间为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <!-- 按钮 -->
      <b-form-group>
        <b-button variant="outline-info" block @click="createContext"
          >创建</b-button
        >
      </b-form-group>
    </b-modal>
    <!-- 创建比赛结束 -->

    <!-- 修改比赛开始 -->
    <b-modal
      id="updateContext"
      ref="updateContext"
      title="修改表单"
      hide-footer
    >
      <b-form-group label="比赛名称">
        <b-form-input
          v-model="$v.updateParams.title.$model"
          required
          :state="validateState1('title')"
          placeholder="请输入比赛名称"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState1('title')">
          比赛名称为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="发起人">
        <b-form-input
          v-model="$v.updateParams.author.$model"
          required
          :state="validateState1('author')"
          placeholder="请输入比赛的发起人"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState1('author')">
          比赛的发起人为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="开始时间">
        <b-form-input
          v-model="$v.updateParams.startTime.$model"
          required
          :state="validateState1('startTime')"
          placeholder="时间格式：xx-xx-xx xx:xx:xx"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState1('startTime')">
          比赛的开始时间为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="结束时间">
        <b-form-input
          v-model="$v.updateParams.endTime.$model"
          required
          :state="validateState1('endTime')"
          placeholder="时间格式：xx-xx-xx xx:xx:xx"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState1('endTime')">
          比赛的结束时间为必填字段
        </b-form-invalid-feedback>
      </b-form-group>
      <!-- 按钮 -->
      <b-form-group>
        <b-button variant="outline-info" block @click="updateContext"
          >修改</b-button
        >
      </b-form-group>
    </b-modal>
    <!-- 修改比赛结束 -->

    <!-- 添加题目开始 -->
    <b-modal
      title="添加题目"
      hide-footer
      id="addProblem"
      ref="addProblem"
      size="xl"
    >
      <b-table
        hover
        :items="problems"
        :fields="field_problem"
        class="text-dark"
      >
        <template #cell(inContext)="data">
          <b-badge pill variant="info" v-if="data.value"> 已存在 </b-badge>
          <b-badge pill variant="warning" v-else>未存在 </b-badge>
        </template>
        <template #cell(problemID)="data">
          <div>
            <b-button
              variant="primary"
              pill
              size="sm"
              @click="addProblemInContext(data.value)"
            >
              添加
            </b-button>
            <b-button
              variant="danger"
              pill
              size="sm"
              class="ml-3"
              @click="deleteProblem(data.value)"
            >
              删除
            </b-button>
          </div>
        </template>
      </b-table>
      <div class="overflow-auto mt-3">
        <b-pagination
          v-model="currentPage_problem"
          :total-rows="total_problem"
          :per-page="params_problem.pageSize"
          align="right"
        ></b-pagination>
      </div>
    </b-modal>
    <!-- 添加题目结束 -->
    <b-card class="mt-3">
      <b-table hover :items="contexts" :fields="fields" class="mt-3">
        <template #cell(context_id)="data">
          <b-button
            variant="info"
            size="sm"
            pill
            v-b-modal.addProblem
            @click="setContextID(data.value)"
          >
            添加题目
          </b-button>
          <b-button
            variant="primary"
            @click="setContextID(data.value)"
            size="sm"
            pill
            v-b-modal.updateContext
            class="ml-3"
          >
            修改
          </b-button>
          <b-button
            variant="danger"
            class="ml-3"
            @click="deleteContext(data.value)"
            size="sm"
            pill
          >
            删除
          </b-button>
        </template>
        <template #cell(defunct)="data">
          <b-badge pill variant="info" v-if="data.value === '1'"
            >未开始</b-badge
          >
          <b-badge pill variant="primary" v-else-if="data.value === '2'"
            >进行中</b-badge
          >
          <b-badge pill variant="danger" v-else>已经结束</b-badge>
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
    </b-card>
  </div>
</template>

<script>
import { required } from 'vuelidate/lib/validators';
import contextService from '../../service/context';

export default {
  data() {
    return {
      // 获取比赛列表的相关内容
      params: {
        pageNum: 1,
        pageSize: 10,
      },
      contexts: [],
      currentPage: 1,
      total: 1,
      fields: [
        { key: 'defunct', label: '比赛状态' },
        { key: 'title', label: '比赛名称' },
        { key: 'start_time', label: '开始时间' },
        { key: 'end_time', label: '结束时间' },
        { key: 'author', label: '发起人' },
        { key: 'context_id', label: '' },
      ],
      // 删除比赛相关内容
      deleteParams: {
        contextID: 1,
      },
      // 更新表单相关
      updateParams: {
        contextID: 1,
        title: '',
        startTime: '',
        endTime: '',
        author: '',
      },
      // 创建表单相关
      createParams: {
        title: '',
        startTime: '',
        endTime: '',
        author: '',
      },
      // 给比赛添加题目相关
      problems: [],
      field_problem: [
        { key: 'inContext', label: '状态' },
        { key: 'author', label: '作者' },
        { key: 'difficulty_level', label: '难度' },
        { key: 'problemName', label: '题目名称' },
        { key: 'time_limit', label: '时间限制（ms）' },
        { key: 'memory_limit', label: '内存限制（b）' },
        { key: 'problemID', label: '' },
      ],
      params_problem: {
        contextID: 1,
        pageNum: 1,
        pageSize: 10,
      },
      currentPage_problem: 1,
      total_problem: 1,
      // 比赛题目相关
      problemInContextParams: {
        contextID: 1,
        problemID: 1,
      },
      // 向比赛里面添加题目
      addProblemInContextParams: {
        contextID: 1,
        problemID: 1,
        title: '',
      },
      // 从比赛中删除题目
      deleteProblemParams: {
        contextID: 1,
        problemID: 1,
      },
    };
  },

  // 数据校验部分
  validations: {
    updateParams: {
      title: {
        required,
      },
      startTime: {
        required,
      },
      endTime: {
        required,
      },
      author: {
        required,
      },
    },
    createParams: {
      title: {
        required,
      },
      startTime: {
        required,
      },
      endTime: {
        required,
      },
      author: {
        required,
      },
    },
  },

  methods: {
    async getContextList() {
      const { data: res } = await contextService.getContextList(this.params);
      this.contexts = res.data.contexts;
      this.title = res.data.title;
    },

    async deleteContext(id) {
      this.deleteParams.contextID = id;
      await contextService.deleteContext(this.deleteParams).then(() => {
        this.$bvToast.toast('删除比赛成功', {
          title: '删除成功',
          variant: 'success',
          solid: true,
        });
        this.getContextList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.dada.msg, {
          title: '删除失败',
          variant: 'danger',
          solid: true,
        });
      });
    },

    setContextID(id) {
      sessionStorage.setItem('context_id', id);
      this.getProblemList();
    },

    clearUpdateParams() {
      this.updateParams.title = '';
      this.updateParams.startTime = '';
      this.updateParams.endTime = '';
      this.updateParams.author = '';
      // eslint-disable-next-line dot-notation
      this.$refs['updateContext'].hide();
      this.$v.$reset();
    },

    clearCreateParams() {
      this.createParams.title = '';
      this.createParams.startTime = '';
      this.createParams.endTime = '';
      this.createParams.author = '';
      // eslint-disable-next-line dot-notation
      this.$refs['createContext'].hide();
      this.$v.$reset();
    },

    async updateContext() {
      this.$v.updateParams.$touch();
      if (this.$v.updateParams.$anyError) {
        return;
      }

      this.updateParams.contextID = parseInt(sessionStorage.getItem('context_id'), 10);

      await contextService.updateContext(this.updateParams).then(() => {
        this.$bvToast.toast('更新比赛成功', {
          title: '更新成功',
          variant: 'success',
          solid: true,
        });

        this.clearUpdateParams();
        this.getContextList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '更新失败',
          variant: 'danger',
          solid: true,
        });
      });
    },

    validateState2(name) {
      // 这里是 es6 的解构赋值
      const { $dirty, $error } = this.$v.createParams[name];
      return $dirty ? !$error : null;
    },

    async createContext() {
      this.$v.createParams.$touch();
      if (this.$v.createParams.$anyError) {
        return;
      }

      await contextService.createContext(this.createParams).then(() => {
        this.$bvToast.toast('创建比赛成功', {
          title: '创建成功',
          variant: 'success',
          solid: true,
        });

        this.clearCreateParams();
        this.getContextList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '更新失败',
          variant: 'danger',
          solid: true,
        });
      });
    },

    validateState1(name) {
      // 这里是 es6 的解构赋值
      const { $dirty, $error } = this.$v.updateParams[name];
      return $dirty ? !$error : null;
    },

    async getProblemList() {
      this.problems = [];
      this.params_problem.contextID = parseInt(sessionStorage.getItem('context_id'), 10);
      const { data: res } = await contextService.getContextProblemList(this.params_problem);
      this.problems = res.data.contextProblems;

      this.total_problem = res.data.total;
    },

    // 关于将题目放到比赛里面

    clearAddParams() {
      this.addProblemInContextParams.problemID = 1;
      this.addProblemInContextParams.contextID = 1;
      this.addProblemInContextParams.title = '';
    },

    async addProblemInContext(problemID) {
      this.problemInContextParams.problemID = parseInt(problemID, 10);
      this.problemInContextParams.contextID = parseInt(sessionStorage.getItem('context_id'), 10);

      this.addProblemInContextParams.problemID = parseInt(problemID, 10);
      this.addProblemInContextParams.contextID = parseInt(sessionStorage.getItem('context_id'), 10);
      this.addProblemInContextParams.title = `problemID_${problemID}`;

      const { data: val } = await contextService.problemInContext(this.problemInContextParams);
      if (val.data.inThere) {
        this.$bvToast.toast('该题目已经在比赛中', {
          title: '向比赛中添加题目失败',
          variant: 'danger',
          solid: true,
        });
        return;
      }

      await contextService.addProblemToContext(this.addProblemInContextParams).then(() => {
        this.$bvToast.toast('向比赛中添加题目成功', {
          title: '创建成功',
          variant: 'success',
          solid: true,
        });

        this.clearAddParams();
        this.getProblemList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '向比赛中添加题目失败',
          variant: 'danger',
          solid: true,
        });
      });
    },

    async problemInContext(problemID) {
      this.problemInContextParams.problemID = parseInt(problemID, 10);
      this.problemInContextParams.contextID = parseInt(sessionStorage.getItem('context_id'), 10);

      const { data: res } = await contextService.problemInContext(this.problemInContextParams);
      return res;
    },

    async deleteProblem(problemID) {
      this.deleteProblemParams.problemID = problemID;
      this.deleteProblemParams.contextID = parseInt(sessionStorage.getItem('context_id'), 10);

      await contextService.deleteProblemInContext(this.deleteProblemParams).then(() => {
        this.$bvToast.toast('从比赛中删除题目成功', {
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
    this.getContextList();
  },

  watch: {
    currentPage() {
      this.params.pageNum = this.currentPage;
      this.getContextList();
    },
    currentPage_problem() {
      this.params_problem.pageNum = this.currentPage_problem;
      this.getProblemList();
    },
  },
};
</script>
