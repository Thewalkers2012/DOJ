<template>
  <div class="mt-3">
    <!-- 显示文章的详细信息的对话框 -->
    <b-modal id="category_details" size="lg" title="文章详情" hide-footer>
      <h5>作者：{{ dCategory.author }}</h5>
      <hr />
      <pre>{{ dCategory.content }}</pre>
    </b-modal>
    <!-- 显示文章详细信息对话框结束 -->
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
            v-b-modal.category_details
          >
            详情
          </b-button>
          <!-- 显示删除 -->
          <b-button
            variant="danger"
            @click="deleteCategory1(data.value)"
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
      // 删除文章相关
      deleteCategory: {},
      deleteParams: {
        categoryID: 1,
      },
      // 细节相关
      detailsCategoryParams: {
        categoryID: 1,
      },
      dCategory: {},
    };
  },
  methods: {
    async getCategoryList() {
      this.getAllCategoriesParams.pageNum = this.currentPage;
      const { data: res } = await categoryService.getAllCategories(this.getAllCategoriesParams);
      this.total = res.data.total;
      this.categories = res.data.categories;
    },
    async deleteCategory1(id) {
      this.deleteParams.categoryID = id;
      const { data: res } = await categoryService.deleteCategory(this.deleteParams);
      if (res.code === 200) {
        this.$bvToast.toast('删除文章成功', {
          title: '删除成功',
          variant: 'success',
          solid: true,
        });
      }
      this.getCategoryList();
    },
    async detailsCategory(id) {
      this.detailsCategoryParams.categoryID = id;
      const { data: res } = await categoryService.getCategoryDetails(this.detailsCategoryParams);
      this.dCategory = res.data.category;
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
