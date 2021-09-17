<template>
  <div id="container" style="height: 600px"></div>
  <!--宽高自行设定 -->
</template>

<script>
import * as monaco from 'monaco-editor';

export default {
  data() {
    return {
      editor: null,
    };
  },
  mounted() {
    this.initEditor();
  },
  methods: {
    initEditor() {
      this.editor = monaco.editor.create(document.getElementById('container'), {
        language: 'cpp',
        value: '',
        automaticLayout: true,
        theme: 'vs-light',
        selectOnLineNumbers: true,
        roundedSelection: false,
        readOnly: false,
        cursorStyle: 'line',
        glyphMargin: true,
        useTabStops: false,
        fontSize: 16,
        autoIndent: true,
        quickSuggestionsDelay: 500,
      });
    },
    getValue() {
      this.editor.getValue();
    },
  },
};
</script>
