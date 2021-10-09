<template>
  <div class="mt-3">
    <h2>{{ context.title }}</h2>
  </div>
</template>

<script>
import contextService from '../../service/context';

export default {
  data() {
    return {
      context: {},
      id: '',
    };
  },

  methods: {
    async getContext() {
      this.id = sessionStorage.getItem('context_id');
      const { data: res } = await contextService.getContext(this.id);
      this.context = res.data.context;
    },
  },

  created() {
    this.getContext();
  },
};
</script>
