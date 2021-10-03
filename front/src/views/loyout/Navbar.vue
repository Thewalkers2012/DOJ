<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-container>
        <b-navbar-brand @click="$router.push({ name: 'Home' })"
          >DOJ</b-navbar-brand
        >
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav>
            <b-nav-item @click="$router.push({ name: 'problemList' })"
              >题库</b-nav-item
            >
            <b-nav-item href="#">关于我们</b-nav-item>
            <b-nav-item @click="$router.push({ name: 'Admin' })"
              >后台管理</b-nav-item
            >
          </b-navbar-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item-dropdown right v-if="userInfo">
              <template v-slot:button-content>
                <em>{{ userInfo.username }}</em>
              </template>
              <b-dropdown-item @click="$router.push({ name: 'profile' })">
                <b-icon icon="person-fill"></b-icon
                ><span class="ml-3">个人主页</span></b-dropdown-item
              >
              <b-dropdown-item @click="logout"
                ><b-icon icon="power" aria-hidden="true"></b-icon
                ><span class="ml-3">登出</span></b-dropdown-item
              >
            </b-nav-item-dropdown>
            <!-- <div v-if="!userInfo"> -->
            <b-nav-item
              v-if="!userInfo && $route.name != 'login'"
              @click="$router.replace({ name: 'login' })"
              >登录</b-nav-item
            >
            <b-nav-item
              v-if="!userInfo && $route.name != 'register'"
              @click="$router.replace({ name: 'register' })"
              >注册</b-nav-item
            >
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),

  methods: mapActions('userModule', ['logout']),
};
</script>

<style lang="scss">
</style>
