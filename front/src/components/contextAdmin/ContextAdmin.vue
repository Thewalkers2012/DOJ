<template>
  <div>
    <div class="d-flex justify-content-between align-items-center">
      <div></div>
      <b-button variant="outline-primary" size="lg"> 添加比赛 </b-button>
    </div>
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
    <b-card class="mt-3">
      <b-table hover :items="contexts" :fields="fields" class="mt-3">
        <template #cell(context_id)="data">
          <b-button
            variant="primary"
            @click="setUpdateID(data.value)"
            size="sm"
            pill
            v-b-modal.updateContext
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

    setUpdateID(id) {
      sessionStorage.setItem('context_id', id);
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

    validateState1(name) {
      // 这里是 es6 的解构赋值
      const { $dirty, $error } = this.$v.updateParams[name];
      return $dirty ? !$error : null;
    },
  },

  created() {
    this.getContextList();
  },

  watch: {
    currentPage() {
      this.params.pageNum = this.currentPage;
      this.getProblemList();
    },
  },
};
</script>
