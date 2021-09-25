<template>
  <div>
    <b-row>
      <b-col
        class="d-flex justify-content-between align-items-center"
        md="8"
        offset-md="2"
      >
        <b-form-input
          v-model="text"
          placeholder="请输入题目的编号"
          type="number"
          class="mr-3"
        ></b-form-input>
        <b-button variant="outline-primary">Search</b-button>
      </b-col>
    </b-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: '',
    };
  },
};
</script>
