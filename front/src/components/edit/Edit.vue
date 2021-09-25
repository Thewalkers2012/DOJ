<template>
  <div id="container" style="height: 800px"></div>
  <!--宽高自行设定 -->
</template>

<script>
import * as monaco from 'monaco-editor';

export default {
  props: {
    propLanguage: String,
    propTheme: String,
  },
  data() {
    return {
      editor: null,
      code: '',
      problemID: '',
    };
  },
  created() {
    this.problemID = sessionStorage.getItem('problem_id');
  },
  mounted() {
    this.initEditor();
  },
  methods: {
    initEditor() {
      this.editor = monaco.editor.create(document.getElementById('container'), {
        language: this.propLanguage,
        value: localStorage.getItem(`code_${this.problemID}`),
        automaticLayout: true,
        theme: this.propTheme,
        selectOnLineNumbers: true,
        roundedSelection: false,
        readOnly: false,
        cursorStyle: 'line',
        glyphMargin: true,
        useTabStops: false,
        fontSize: 17,
        autoIndent: true,
        quickSuggestionsDelay: 500,
      });
      this.editor.onKeyUp(() => {
        // 当按下键盘，判断当前编辑器的文本与已保存的编辑器文本是否一致
        this.code = this.editor.getValue();
        localStorage.setItem(`code_${this.problemID}`, this.code);
      });
    },
  },
  watch: {
    propLanguage() {
      monaco.editor.setModelLanguage(this.editor.getModel(), this.propLanguage);
    },
    propTheme() {
      this.editor.updateOptions({
        theme: this.propTheme,
      });
    },
  },
};
</script>
