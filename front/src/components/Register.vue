<template>
  <div>
    <div class="row mt-5">
      <div class="col-md-8 offset-md-2 col-lg-6 offset-lg-3">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">登录</h5>
            <form @submit.prevent="validate">
              <div class="mt-3 form-group">
                <label for="studentID" class="form-label">学号</label>
                <input
                  class="form-control"
                  placeholder="请输入学号"
                  id="studentID"
                  v-bind:class="{ 'is-invalid': studentIDError }"
                  v-model="studentID"
                />
                <div
                  class="invalid-feedback"
                  id="feedback-1"
                  v-if="studentIDError"
                >
                  {{ studentIDErrorMsg }}
                </div>
              </div>
              <div class="mt-3 form-group">
                <label for="username" class="form-label">用户名</label>
                <input
                  class="form-control"
                  placeholder="请输入用户名"
                  id="username"
                  v-bind:class="{ 'is-invalid': nameError }"
                  v-model="name"
                />
                <div
                  class="invalid-feedback"
                  id="feedback-2"
                  v-if="nameErrorMsg"
                >
                  {{ nameErrorMsg }}
                </div>
              </div>
              <div class="mt-3 form-group">
                <label for="password" class="form-label">密码</label>
                <input
                  type="password"
                  class="form-control"
                  placeholder="请输入密码"
                  id="password"
                  v-bind:class="{ 'is-invalid': passwordError }"
                  v-model="password"
                />
                <div
                  class="invalid-feedback"
                  id="feedback-3"
                  v-if="passwordErrorMsg"
                >
                  {{ passwordErrorMsg }}
                </div>
              </div>
              <div class="mt-3 d-grid gap-2">
                <button type="submit" class="btn btn-outline-primary">
                  登录
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DojRegister',

  data() {
    return {
      studentID: '',
      name: '',
      password: '',
      nameError: false,
      studentIDError: false,
      passwordError: false,
      nameErrorMsg: '',
      passwordErrorMsg: '',
      studentIDErrorMsg: '',
      preName: true,
      prePassword: true,
      preStudentID: true,
    }
  },

  methods: {
    // 数据校验
    validate() {
      this.validStudentID()
      this.validName()
      this.validPassword()
    },
    validStudentID() {
      var len = this.studentID.length
      if (len != 10) {
        this.preStudentID = false
        this.studentIDError = true
        this.studentIDErrorMsg = '学号必须为 10 位'
      } else {
        if (this.preStudentID) {
          document.getElementById('studentID').className = 'form-control is-valid'
          return
        } else {
          this.preStudentID = true 
        }
        this.studentIDErrorMsg = ''
        document.getElementById('studentID').className = 'form-control is-valid'
        document.getElementById('feedback-1').className = 'valid-feedback'
      }
    },
    validName() {
      var len = this.name.length
      if (len == 0) {
        this.preName = false
        this.nameError = true
        this.nameErrorMsg = '用户不能为空'
      } else {
        if (this.preName) {
          document.getElementById('username').className = 'form-control is-valid'
          return
        } else {
          this.preName = true
        }
        this.nameErrorMsg = ''
        document.getElementById('username').className = 'form-control is-valid'
        document.getElementById('feedback-2').className = 'valid-feedback'
      }
    },
    validPassword() {
      var len = this.password.length
      if (len < 6 || len > 12) {
        this.passwordError = true
        this.passwordErrorMsg = '密码必须大于等于 6 位并小于等于 12 位'
        this.prePassword = false
      } else {
        if (this.prePassword) {
          document.getElementById('password').className = 'form-control is-valid'
          return
        } else {
          this.prePassword = true
        }
        this.passwordErrorMsg = ''
        document.getElementById('password').className = 'form-control is-valid'
        document.getElementById('feedback-3').className = 'valid-feedback'
      }
    }
  }
}
</script>
