<template>
  <div class="mt-3">
    <div>
      <div>
        <!-- 修改角色表单 -->
        <b-modal id="updateUser" ref="updateUser" title="修改表单" hide-footer>
          <!-- 用户名 -->
          <b-form-group label="用户名">
            <b-form-input
              v-model="$v.updateParams.username.$model"
              required
              :state="validateState1('username')"
              placeholder="请输入用户名"
            ></b-form-input>
            <b-form-invalid-feedback :state="validateState1('username')">
              用户名为必填字段
            </b-form-invalid-feedback>
          </b-form-group>
          <b-form-group label="学号">
            <b-form-input
              v-model="$v.updateParams.studentID.$model"
              required
              :state="validateState1('studentID')"
              placeholder="请输入学号"
            ></b-form-input>
            <b-form-invalid-feedback :state="validateState1('studentID')">
              学号必须为 11 位
            </b-form-invalid-feedback>
          </b-form-group>
          <!-- 按钮 -->
          <b-form-group>
            <b-button variant="outline-info" block @click="updateUser"
              >修改</b-button
            >
          </b-form-group>
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
        <b-modal id="addUser" ref="addUser" title="增加表单" hide-footer>
          <!-- 用户名 -->
          <b-form-group label="用户名">
            <b-form-input
              v-model="$v.addParams.username.$model"
              required
              :state="validateState2('username')"
              placeholder="请输入用户名"
            ></b-form-input>
            <b-form-invalid-feedback :state="validateState2('username')">
              用户名为必填字段
            </b-form-invalid-feedback>
          </b-form-group>
          <b-form-group label="学号">
            <b-form-input
              v-model="$v.addParams.studentID.$model"
              required
              :state="validateState2('studentID')"
              placeholder="请输入学号"
            ></b-form-input>
            <b-form-invalid-feedback :state="validateState2('studentID')">
              学号必须为 11 位
            </b-form-invalid-feedback>
          </b-form-group>
          <b-form-group label="密码">
            <b-form-input
              type="password"
              v-model="$v.addParams.password.$model"
              required
              :state="validateState2('password')"
              placeholder="请输入密码"
            ></b-form-input>
            <b-form-invalid-feedback :state="validateState2('password')">
              密码必须大于 6 位并小于 12 位
            </b-form-invalid-feedback>
          </b-form-group>
          <!-- 按钮 -->
          <b-form-group>
            <b-button variant="outline-info" block @click="addUser"
              >添加</b-button
            >
          </b-form-group>
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
              @click="setUpdateID(data.value)"
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
import { required, minLength, maxLength } from 'vuelidate/lib/validators';
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
  // 数据校验部分
  validations: {
    addParams: {
      username: {
        required,
      },
      studentID: {
        required,
        minLength: minLength(10),
        maxLength: maxLength(10),
      },
      password: {
        required,
        minLength: minLength(6),
        maxLength: maxLength(12),
      },
    },
    updateParams: {
      username: {
        required,
      },
      studentID: {
        required,
        minLength: minLength(10),
        maxLength: maxLength(10),
      },
    },
  },
  methods: {
    async getUserList() {
      const { data: res } = await userService.userList(this.params);
      this.users = res.data.users;
      this.total = res.data.total;
    },
    // 添加角色相关
    validateState2(name) {
      const { $dirty, $error } = this.$v.addParams[name];
      return $dirty ? !$error : null;
    },

    clearAddParams() {
      this.addParams.username = '';
      this.addParams.studentID = '';
      this.addParams.password = '';
      // eslint-disable-next-line dot-notation
      this.$refs['addUser'].hide();
      this.$v.$reset();
    },
    async addUser() {
      this.$v.addParams.$touch();
      if (this.$v.addParams.$anyError) {
        return;
      }

      await userService.register(this.addParams).then(() => {
        this.$bvToast.toast('创建角色成功', {
          title: '创建成功',
          variant: 'success',
          solid: true,
        });
        this.clearAddParams();
        this.getUserList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '创建失败',
          variant: 'danger',
          solid: true,
        });
      });
    },
    // 更新角色相关
    validateState1(name) {
      // 这里是 es6 的解构赋值
      const { $dirty, $error } = this.$v.updateParams[name];
      return $dirty ? !$error : null;
    },
    clearUpdateParams() {
      this.updateParams.username = '';
      this.updateParams.studentID = '';
      this.updateParams.userID = 1;
      // eslint-disable-next-line dot-notation
      this.$refs['updateUser'].hide();
      this.$v.$reset();
    },
    setUpdateID(id) {
      this.updateParams.userID = id;
      console.log(this.updateParams.userID);
    },
    async updateUser() {
      this.$v.updateParams.$touch();
      if (this.$v.updateParams.$anyError) {
        return;
      }

      await userService.updateUser(this.updateParams).then(() => {
        this.$bvToast.toast('修改角色成功', {
          title: '修改成功',
          variant: 'success',
          solid: true,
        });
        this.clearUpdateParams();
        this.getUserList();
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '修改失败',
          variant: 'danger',
          solid: true,
        });
      });
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
