<template>
  <div class="mt-3">
    <div>
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
        <!-- 显示角色详细信息的开始 -->
        <b-modal id="user_details" size="lg" title="用户详情">
          <h5>学号：{{ user.studentID }}</h5>
          <hr />
          <h5>用户名：{{ user.username }}</h5>
          <hr />
          <h5>提交总数：{{ user.acceptCount }}</h5>
          <hr />
          <h5>解决问题总数：{{ user.submissionCount }}</h5>
          <!-- 显示角色具体信息的结束 -->
        </b-modal>
        <!-- 添加角色显示位置 -->
        <b-modal
          id="addUser"
          ref="modal"
          title="增加表单"
          @show="addResetModal"
          @hidden="addResetModal"
          @ok="addHandleOK"
        >
          <form ref="form" @submit.stop.prevent="addSubmit">
            <b-form-group
              label="用户名"
              label-for="name-input"
              invalid-feedback="用户名为必填项"
              :state="addUsernameState"
            >
              <b-form-input
                id="name-input"
                v-model="addParams.username"
                :state="addUsernameState"
                required
              ></b-form-input>
            </b-form-group>
            <b-form-group
              label="学号"
              label-for="studentID-input"
              invalid-feedback="学号为必填项"
              :state="addStudentIDState"
            >
              <b-form-input
                id="studentID-input"
                v-model="addParams.studentID"
                :state="addStudentIDState"
                required
              ></b-form-input>
            </b-form-group>
            <b-form-group
              label="密码"
              label-for="password-input"
              invalid-feedback="密码为必填项"
              :state="addPasswordState"
            >
              <b-form-input
                id="password-input"
                v-model="addParams.password"
                :state="addPasswordState"
                required
              ></b-form-input>
            </b-form-group>
          </form>
        </b-modal>
        <!-- 添加角色显示位置结束 -->
      </div>
      <div class="d-flex justify-content-between align-items-center mt-3 mb-3">
        <div></div>
        <b-button variant="outline-primary" size="lg" v-b-modal.addUser
          >添加学生</b-button
        >
      </div>
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
              v-b-modal.user_details
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
      // 添加角色相关
      addParams: {
        username: '',
        studentID: '',
        password: '',
      },
      addUsernameState: null,
      addStudentIDState: null,
      addPasswordState: null,
      // 获取用户详细信息相关
      detailsParams: {
        userID: '',
      },
      // 删除角色相关
      deleteParams: {
        userID: 0,
      },
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
    // 添加角色相关
    validateAddUsername() {
      const valid = this.$refs.form.checkValidity();
      this.addUsernameState = valid;
      return valid;
    },
    validateAddStudentID() {
      const valid = this.$refs.form.checkValidity();
      this.addStudentIDState = valid;
      return valid;
    },
    validateAddPassword() {
      const valid = this.$refs.form.checkValidity();
      this.addPasswordState = valid;
      return valid;
    },
    addResetModal() {
      this.addParams.username = '';
      this.addUsernameState = null;
      this.addParams.studentID = '';
      this.addStudentIDState = null;
      this.addParams.password = '';
      this.addPasswordState = null;
    },
    addHandleOK(bvModalEvt) {
      bvModalEvt.preventDefault();
      this.addSubmit();
    },
    async addSubmit() {
      if (!this.validateAddUsername()) {
        return;
      }
      if (!this.validateAddStudentID()) {
        return;
      }
      if (!this.validateAddPassword()) {
        return;
      }

      await userService.register(this.addParams).then(() => {
        this.$bvToast.toast('添加角色成功', {
          title: '添加成功',
          variant: 'success',
          solid: true,
        });
        this.getUserList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '登录失败',
          variant: 'danger',
          solid: true,
        });
      });

      this.$nextTick(() => {
        this.$bvModal.hide('addUser');
      });
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
    async deleteUser(id) {
      console.log(id);
      this.deleteParams.userID = id;
      const { data: res } = await userService.deleteUser(this.deleteParams);
      if (res.code === 200) {
        this.$bvToast.toast('删除角色成功', {
          title: '删除成功',
          variant: 'success',
          solid: true,
        });
      } else {
        this.$bvToast.toast(res.data.msg, {
          title: '删除失败',
          variant: 'danger',
          solid: true,
        });
      }
      // 重新获取列表
      this.getUserList();
    },
    // 角色细节
    async detailsUser(id) {
      this.detailsParams.userID = id;
      const { data: res } = await userService.getUserDetails(this.detailsParams);
      this.user = res.data.userDetail;
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
