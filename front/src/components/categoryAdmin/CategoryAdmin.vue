<template>
  <div class="mt-3">
    <b-table
      striped
      hover
      :items="categories"
      :fields="fields"
      class="text-secondary"
    >
      <template #cell(category_id)="data">
        <div class="flex justify-content-between align-items-center">
          <!-- 显示详情 -->
          <b-button
            variant="primary"
            @click="detailsCategory(data.value)"
            size="sm"
            v-b-modal.user_details
          >
            详情
          </b-button>
          <!-- 显示删除 -->
          <b-button
            variant="danger"
            @click="deleteCategory(data.value)"
            size="sm"
            class="ml-3"
          >
            删除
          </b-button>
        </div>
      </template>
    </b-table>
    <div class="overflow-auto mt-3">
      <b-pagination
        v-model="currentPage"
        :total-rows="total"
        :per-page="getAllCategoriesParams.pageSize"
        align="right"
      ></b-pagination>
    </div>
  </div>
</template>

<script>
import categoryService from '../../service/category';

export default {
  data() {
    return {
      categories: [],
      // 获取所有文章相关
      getAllCategoriesParams: {
        pageNum: 1,
        pageSize: 10,
      },
      currentPage: 1,
      total: 1,
      fields: [
        { key: 'author', label: '作者' },
        { key: 'user_id', label: '用户编号' },
        { key: 'problem_id', label: '相关题目' },
        { key: 'create_at', label: '创建时间' },
        { key: 'category_id', label: '操作' },
      ],
    };
  },
  methods: {
    async getCategoryList() {
      this.getAllCategoriesParams.pageNum = this.currentPage;
      const { data: res } = await categoryService.getAllCategories(this.getAllCategoriesParams);
      this.total = res.data.total;
      this.categories = res.data.categories;
    },
    deleteCategory(id) {
      console.log(id);
    },
    detailsCategory(id) {
      console.log(id);
    },
  },
  created() {
    this.getCategoryList();
  },
  watch: {
    currentPage() {
      this.getAllCategoriesParams.pageNum = this.currentPage;
      this.getCategoryList();
    },
  },
};
</script>
