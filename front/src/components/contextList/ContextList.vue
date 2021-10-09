<template>
  <div class="mt-3">
    <b-card>
      <h2 class="text-center">进行中的比赛</h2>
      <b-table
        striped
        hover
        :items="contexts"
        :fields="fields"
        class="text-secondary mt-3"
      >
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
