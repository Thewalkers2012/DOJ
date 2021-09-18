<template>
  <div id="container" style="height: 600px"></div>
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
    };
  },
  mounted() {
    this.initEditor();
  },
  methods: {
    initEditor() {
      this.editor = monaco.editor.create(document.getElementById('container'), {
        language: this.propLanguage,
        value: '',
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
    },
    getValue() {
      this.editor.getValue();
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
