<template>
  <div id="app">
    <Navbar />
    <b-container>
      <router-view />
    </b-container>
  </div>
</template>

<script>
import Navbar from './views/loyout/Navbar.vue';

export default {
  components: {
    Navbar,
  },
  data() {
    return {

    };
  },
};
</script>

<style lang="scss">
</style>
