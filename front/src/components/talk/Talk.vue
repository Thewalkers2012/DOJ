<template>
  <div class="mt-3">
    <h4>评论</h4>
    <div>
      <div class="mt-3">
        <b-form-textarea
          id="textarea-large"
          size="lg"
          placeholder="发起一个评论吧"
          v-model="createParams.content"
        ></b-form-textarea>
      </div>
      <div class="d-flex mt-3 justify-content-between align-items-center">
        <div></div>
        <div>
          <b-button
            variant="outline-info"
            size="md"
            class="ml-3"
            @click="createCategoryByProblem"
          >
            评论
          </b-button>
        </div>
      </div>
    </div>
    <hr />
    <h4 class="text-center mt-3">精选评论</h4>
    <b-overlay
      rounded="sm"
      v-for="item in categories"
      v-bind:key="item.ID"
      class="mt-3"
    >
      <b-card>
        <div class="d-flex mt-3 justify-content-between align-items-center">
          <h5>{{ item.author }}</h5>
          <h5 class="mt-3">{{ item.create_at }}</h5>
        </div>
        <b-card-text> {{ item.content }} </b-card-text>
      </b-card>
    </b-overlay>
  </div>
</template>

<script>
import categoryService from '../../service/category';

export default {
  data() {
    return {
      createParams: {
        problemID: 0,
        content: '',
      },
      getParams: {
        problemID: 0,
      },
      categories: [],
      category: {},
    };
  },
  created() {
    this.getParams.problemID = parseInt(sessionStorage.getItem('problem_id'), 10);
    this.getCategoryByProblem();
  },
  methods: {
    async getCategoryByProblem() {
      const { data: res } = await categoryService.getCategoryByProblem(this.getParams);
      this.categories = res.data.categories;
    },
    async createCategoryByProblem() {
      if (this.createParams.content.length === 0) {
        this.$bvToast.toast('内容不能为空', {
          title: '评论失败',
          variant: 'danger',
          solid: true,
        });
        return;
      }

      this.createParams.problemID = parseInt(sessionStorage.getItem('problem_id'), 10);
      const { data: res } = await categoryService.createCategory(this.createParams);
      this.createParams.title = '';
      this.createParams.content = '';
      this.category = res.data.category;
      this.getCategoryByProblem();
    },
  },
};
</script>
