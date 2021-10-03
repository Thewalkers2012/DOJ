<template>
  <div class="mt-3">
    <div>
      <b-table
        striped
        hover
        :items="users"
        :fields="fields"
        class="text-secondary"
      >
        <template #cell(user_id)="data">
          <div class="flex justify-content-between align-items-center">
            <b-button
              variant="warning"
              @click="changeUser(data.value)"
              size="sm"
            >
              修改
            </b-button>
            <b-button
              variant="danger"
              @click="deleteUser(data.value)"
              size="sm"
              class="ml-3"
            >
              删除
            </b-button>
          </div>
        </template>
      </b-table>
    </div>
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
import userService from '../../service/userService';

export default {
  data() {
    return {
      fields: [
        { key: 'username', label: '用户名' },
        { key: 'student_id', label: '学号' },
        { key: 'created_at', label: '创建时间' },
        { key: 'user_id', label: '操作' },
      ],
      users: [],
      params: {
        pageSize: 10,
        pageNum: 1,
      },
      total: 0,
      currentPage: 1,
    };
  },
  methods: {
    async getUserList() {
      const { data: res } = await userService.userList(this.params);
      this.users = res.data.users;
      this.total = res.data.total;
    },
    changeUser(id) {
      console.log(id);
    },
    deleteUser(id) {
      console.log(id);
    },
  },
  created() {
    this.getUserList();
  },
  watch: {
    currentPage() {
      this.params.pageNum = this.currentPage;
      this.getUserList();
    },
  },
};
</script>
