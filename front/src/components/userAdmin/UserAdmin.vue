<template>
  <div class="mt-3">
    <div>
      <!-- 修改角色表单 -->
      <b-modal
        id="updateUser"
        ref="modal"
        title="修改表单"
        @show="resetModal"
        @hidden="resetModal"
        @ok="handleOK"
      >
        <form ref="form" @submit.stop.prevent="updateSubmit">
          <b-form-group
            label="用户名"
            label-for="name-input"
            invalid-feedback="用户名为必填项"
            :state="updateUsernameState"
          >
            <b-form-input
              id="name-input"
              v-model="updateParams.username"
              :state="updateUsernameState"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="学号"
            label-for="studentID-input"
            invalid-feedback="学号为必填项"
            :state="updateStudentIDState"
          >
            <b-form-input
              id="studentID-input"
              v-model="updateParams.studentID"
              :state="updateStudentIDState"
              required
            ></b-form-input>
          </b-form-group>
        </form>
      </b-modal>
      <!-- 修改角色结束 -->
      <b-table
        striped
        hover
        :items="users"
        :fields="fields"
        class="text-secondary"
      >
        <template #cell(user_id)="data">
          <div class="flex justify-content-between align-items-center">
            <!-- 显示详情 -->
            <b-button
              variant="primary"
              @click="detailsUser(data.value)"
              size="sm"
            >
              详情
            </b-button>
            <!-- 显示修改 -->
            <b-button
              variant="warning"
              @click="changeUser(data.value)"
              size="sm"
              class="ml-3"
              v-b-modal.updateUser
            >
              修改
            </b-button>
            <!-- 显示删除 -->
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
      // 列表相关
      fields: [
        { key: 'username', label: '用户名' },
        { key: 'student_id', label: '学号' },
        { key: 'created_at', label: '创建时间' },
        { key: 'user_id', label: '操作' },
      ],
      users: [],
      // 更改用户相关
      updateParams: {
        username: '',
        userID: 1,
        studentID: '',
      },
      updateUsernameState: null,
      updateStudentIDState: null,
      // 页码相关
      params: {
        pageSize: 10,
        pageNum: 1,
      },
      total: 0,
      currentPage: 1,
      user: {},
    };
  },
  methods: {
    async getUserList() {
      const { data: res } = await userService.userList(this.params);
      this.users = res.data.users;
      this.total = res.data.total;
    },
    // 更新角色相关
    validateUpdateUserName() {
      const valid = this.$refs.form.checkValidity();
      this.updateUsernameState = valid;
      return valid;
    },
    validateUpdateStudentID() {
      const valid = this.$refs.form.checkValidity();
      this.updateStudentIDState = valid;
      return valid;
    },
    changeUser(id) {
      sessionStorage.setItem('updateUserID', id);
    },
    async updateSubmit() {
      // Exit when the form isn't valid
      if (!this.validateUpdateUserName()) {
        return;
      }

      if (!this.validateUpdateStudentID()) {
        return;
      }

      this.updateParams.userID = parseInt(sessionStorage.getItem('updateUserID'), 10);
      const { data: res } = await userService.updateUser(this.updateParams);
      this.user = res.data.user;
      this.getUserList();
      this.$nextTick(() => {
        this.$bvModal.hide('updateUser');
      });
    },
    resetModal() {
      this.updateParams.username = '';
      this.updateUsernameState = null;
      this.updateParams.studentID = '';
      this.updateStudentIDState = null;
    },
    handleOK(bvModalEvt) {
      bvModalEvt.preventDefault();
      this.updateSubmit();
    },
    // 删除角色相关
    deleteUser(id) {
      sessionStorage.setItem('deleteUserID', id);
    },
    // 角色细节
    detailsUser(id) {
      sessionStorage.setItem('detailsUserID', id);
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
