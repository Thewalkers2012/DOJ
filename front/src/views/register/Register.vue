<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.username.$model"
                type="text"
                required
                placeholder="请输入您的姓名"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('username')">
                学生的姓名不能为空
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="学号">
              <b-form-input
                v-model="$v.user.studentID.$model"
                type="number"
                required
                :state="validateState('studentID')"
                placeholder="请输入您的学号"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('studentID')">
                学号必须为 10 位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                required
                :state="validateState('password')"
                placeholder="请输入您的密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码必须大于等于 6 位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button variant="outline-primary" block @click="register"
                >注册</b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        username: '',
        studentID: '',
        password: '',
      },
    };
  },

  validations: {
    user: {
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
      },
    },
  },

  methods: {
    ...mapActions('userModule', { userRegister: 'register' }),
    validateState(name) {
      // 这里是 es6 的解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userRegister(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '注册失败',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
