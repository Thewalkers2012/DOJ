<template>
  <div>
    <div class="d-flex justify-content-between align-items-center">
      <div></div>
      <b-button variant="outline-primary" size="lg"> 添加比赛 </b-button>
    </div>
    <b-card class="mt-3">
      <b-table hover :items="contexts" :fields="fields" class="mt-3">
        <template #cell(context_id)="data">
          <b-button
            variant="primary"
            @click="updateContext(data.value)"
            size="sm"
            pill
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
    };
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

    updateContext(id) {
      console.log(id);
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
