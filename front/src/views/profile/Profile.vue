<template>
  <h3>
    您好 <span> {{ userInfo.name }} </span>
  </h3>
</template>

<script>
import { mapState } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
};
</script>

<style lang="scss" scoped>
</style>
