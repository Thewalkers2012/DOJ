<template>
  <div class="mt-3">
    <h2 class="text-center" v-if="isUser">用户管理</h2>
    <h2 class="text-center" v-else-if="isCategory">文章管理</h2>
    <b-card class="mt-3">
      <div class="d-flex justify-content-between align-items-center">
        <div></div>
        <b-button-group>
          <b-button variant="success" @click="changeToUser">用户管理</b-button>
          <b-button variant="info" @click="changeToProblem">题目管理</b-button>
          <b-button variant="warning" @click="changeToContext"
            >比赛管理</b-button
          >
          <b-button variant="danger" @click="changeToCategory"
            >文章管理</b-button
          >
        </b-button-group>
        <div></div>
      </div>
      <hr />
      <UserAdmin v-if="isUser"></UserAdmin>
      <CategoryAdmin v-else-if="isCategory"></CategoryAdmin>
    </b-card>
  </div>
</template>

<script>
import UserAdmin from '../../components/userAdmin/UserAdmin.vue';
import CategoryAdmin from '../../components/categoryAdmin/CategoryAdmin.vue';

export default {
  name: 'Admin',
  data() {
    return {
      isUser: true,
      isProblem: false,
      isContext: false,
      isCategory: false,
    };
  },
  methods: {
    changeToProblem() {
      this.isProblem = true;
      this.isUser = false;
      this.isContext = false;
      this.isCategory = false;
    },
    changeToUser() {
      this.isUser = true;
      this.isProblem = false;
      this.isContext = false;
      this.isCategory = false;
    },
    changeToContext() {
      this.isContext = true;
      this.isProblem = false;
      this.isUser = false;
      this.isCategory = false;
    },
    changeToCategory() {
      this.isCategory = true;
      this.isProblem = false;
      this.isContext = false;
      this.isUser = false;
    },
  },
  components: {
    UserAdmin,
    CategoryAdmin,
  },
};
</script>
