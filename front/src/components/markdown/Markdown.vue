<template>
  <div class="mt-3">
    <v-md-editor :value="content" height="800px" mode="preview"></v-md-editor>
    <v-md-editor v-model="text" height="800px"> </v-md-editor>
  </div>
</template>

<script>

export default {
  name: 'Markdown',
  data() {
    return {
      text: '',
      content: '# adasjdh\n## adhajsdh\n### adjsdkal\n#### dajkdsjakld',
    };
  },
};
</script>
