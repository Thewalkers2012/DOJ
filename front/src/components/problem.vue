/* eslint-disable max-len */
<template>
  <div class="Problem">
    <b-card class="mt-3">
      <h2 class="mt-3" style="font-weight: bold">
        {{ problem.problem_id }}. {{ problem.problem_name }}
      </h2>
      <div class="mt-3">
        <b-nav tabs fill style="color: black">
          <b-nav-item @click="clickDescription">
            <div style="color: black">题目</div>
          </b-nav-item>
          <b-nav-item @click="clickSubmittion">
            <div style="color: black">提交记录</div>
          </b-nav-item>
          <b-nav-item @click="clickTalk">
            <div style="color: black">讨论</div>
          </b-nav-item>
        </b-nav>
      </div>
      <!-- description -->
      <router-view></router-view>
      <Description v-if="description"></Description>
      <Category v-else-if="talk"></Category>
      <Submissions v-else> </Submissions>
    </b-card>
  </div>
</template>

<script>
import problemService from '../service/problemService';
import Description from './description/Description.vue';
import Submissions from './submissions/Submissions.vue';
import Category from './talk/Talk.vue';

export default {
  data() {
    return {
      problem: {},
      description: true,
      submittion: false,
      talk: false,
    };
  },
  created() {
    const id = sessionStorage.getItem('problem_id');
    this.getProblemByID(id);
  },
  methods: {
    async getProblemByID(id) {
      const { data: res } = await problemService.getProblemByID(id);
      this.problem = res.data.problem;
    },
    clickDescription() {
      this.description = true;
      this.submition = false;
      this.talk = false;
    },
    clickSubmittion() {
      this.description = false;
      this.submittion = true;
      this.talk = false;
    },
    clickTalk() {
      this.description = false;
      this.submittion = false;
      this.talk = true;
    },
  },
  components: {
    Description,
    Category,
    Submissions,
  },
};
</script>
