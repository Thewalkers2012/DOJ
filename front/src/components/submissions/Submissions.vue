<template>
  <div>
    <b-modal id="code-modal" size="xl">
      <pre v-highlightjs><code :class="language">{{ code }}</code></pre>
    </b-modal>
    <b-table
      striped
      hover
      :items="submissions"
      :fields="fields"
      class="text-secondary"
    >
      <template #cell(submission_id)="data">
        <b-button
          variant="info"
          @click="getCode(data.value)"
          size="sm"
          v-b-modal="'code-modal'"
        >
          Code
        </b-button>
      </template>
      <template #cell(result)="data">
        <span :class="getColor(data.value)">
          {{ data.value }}
        </span>
      </template>
    </b-table>
  </div>
</template>

<script>
import submitService from '../../service/problemService';

export default {
  data() {
    return {
      fields: [
        { key: 'created_at', label: '提交时间' },
        { key: 'result', label: '运行结果' },
        { key: 'time', label: '运行时间' },
        { key: 'memory', label: '运行内存' },
        { key: 'language', label: '语言' },
        { key: 'submission_id', label: '查看代码' },
      ],
      params: {
        problemID: 0,
      },
      submissions: [],
      code: '',
      language: '',
    };
  },
  methods: {
    async getAllSubmissions() {
      const { data: res } = await submitService.getSubmissions(this.params);
      this.submissions = res.data.submission;

      for (let i = 0; i < this.submissions.length; i += 1) {
        this.submissions[i].time += ' ms';
        this.submissions[i].memory += ' kb';
      }
    },
    getColor(msg) {
      if (msg === 'Accept') return 'text-success';
      return 'text-danger';
    },
    getCode(id) {
      for (let i = 0; i < this.submissions.length; i += 1) {
        if (this.submissions[i].submission_id === id) {
          this.code = this.submissions[i].code;
          this.language = this.submissions[i].language;
        }
      }
    },
  },
  created() {
    this.params.problemID = parseInt(sessionStorage.getItem('problem_id'), 10);
    this.getAllSubmissions();
  },
};
</script>
