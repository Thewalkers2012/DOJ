<template>
  <div class="mt-3">
    <b-card>
      <h2 class="text-center">比赛</h2>
      <b-table hover :items="contexts" :fields="fields" class="mt-3">
        <template #cell(context_id)="data">
          <b-button
            variant="info"
            @click="getContext(data.value)"
            size="sm"
            pill
          >
            点击进入
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
    };
  },

  methods: {
    async getContextList() {
      const { data: res } = await contextService.getContextList(this.params);
      this.contexts = res.data.contexts;
      this.title = res.data.title;
    },

    getContext(id) {
      sessionStorage.setItem('context_id', id);
      this.$router.push(`/context/${id}`);
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
