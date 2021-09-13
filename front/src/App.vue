<template>
  <doj-navbar></doj-navbar>
  <div class="container">
    <router-view></router-view>
  </div>
</template>

<script>

import DojNavbar from './views/layout/Navbar.vue'

export default {
  name: 'App',
  components: {
    DojNavbar,
  }
}
</script>
