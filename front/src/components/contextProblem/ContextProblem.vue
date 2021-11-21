<template>
  <div>
    <!-- 题目描述区域 -->
    <div class="mt-3">
      <v-md-editor :value="problem.description" mode="preview"></v-md-editor>
    </div>
    <!-- 代码编辑区域 -->
    <b-card class="mt-5">
      <div class="d-flex justify-content-between align-items-center">
        <div>
          <h4>编辑器</h4>
        </div>
        <div class="d-flex">
          <b-form-select
            v-model="selected2"
            :options="options2"
            class="ml-3"
          ></b-form-select>
          <b-form-select
            class="ml-3"
            v-model="selected1"
            :options="options1"
          ></b-form-select>
        </div>
      </div>
      <!-- 编辑器组件 -->
      <Edit class="mt-3" :propLanguage="selected1" :propTheme="selected2">
      </Edit>
      <!-- <pre v-highlightjs><code class="cpp">{{ code }}</code></pre> -->
    </b-card>
    <!-- 提交代码部分 -->
  </div>
</template>

<script>
export default {
  name: 'ContextProblem',
  data() {

  },
};
</script>
