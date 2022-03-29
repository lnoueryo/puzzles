<template>
  <div>
    <v-row class="mb-6" no-gutters>
      <v-col sm="5" md="6" class="pr-1">
        <div style="overflow: scroll;height: 312px;">
          <v-textarea
            ref="detail"
            filled
            label="タスクの詳細"
            auto-grow
            rows="12"
            hide-details
            @keydown.89.90="preventUndoRedo"
            v-model="currentText"
          ></v-textarea>
        </div>
        <template v-if="$vuetify.breakpoint.mdAndUp">
          <v-btn-toggle class="my-1" color="primary" dense>
            <v-btn v-for="(btn, i) in buttons" :key="i" :value="btn.value" text @click="btn.func">
              <v-icon>mdi-{{ btn.icon }}</v-icon>
            </v-btn>
          </v-btn-toggle>
        </template>
      </v-col>
      <v-col sm="5" md="6" class="pl-1">
        <div class="px-4" style="border-radius: 3px 3px 0 0;overflow: scroll;height: 312px;background-color: rgba(255, 255, 255, 0.08)">
          <div v-html="html"></div>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import Vue from 'vue'
// import { Markdown } from '~/modules/markdown';
export default Vue.extend({
  props: ['text'],
  data:() => ({
    markdown: [''],
    currentIndex: 0,
    textarea: '',
    buttons: [],
  }),
  computed: {
    html() {
      const html = this.$md.render(this.markdown[this.currentIndex] ?? '');
      return html;
    },
    currentText: {
      get() {
        return this.markdown[this.currentIndex];
      },
      set(v) {
        if(this.currentIndex == 200) {
          this.currentIndex -= 1;
          this.markdown.shift();
        }
        if(this.markdown.length - 1 != this.currentIndex) {
          this.markdown = this.markdown.filter((_, index) => {
            return this.currentIndex >= index;
          })
        }
        this.markdown.push(v)
        this.currentIndex += 1;
        this.$emit('currentText', this.markdown[this.currentIndex])
      }
    }
  },
  watch: {
    text: {
      handler(v) {
        console.log(v);
      }
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.textarea = this.$refs.detail.$refs.input;
      // this.markdown = new Markdown(textarea);
      this.buttons = [
        {value: 1, icon: 'format-bold', func: this.bold},
        {value: 2, icon: 'format-italic', func: this.italic},
        {value: 3, icon: 'format-underline', func: this.underline},
        {value: 4, icon: 'format-color-fill', func: this.italic},
        {value: 5, icon: 'border-color', func: this.italic},
        {value: 6, icon: 'format-align-left', func: this.italic},
        {value: 7, icon: 'format-align-center', func: this.italic},
        {value: 8, icon: 'format-align-right', func: this.italic},
        {value: 9, icon: 'format-align-justify', func: this.italic},
      ]
      this.markdown[1] = this.text;
      this.currentIndex += 1
    })
  },
  methods: {
    inputText(v) {
      if(this.currentIndex == 200) {
        this.currentIndex -= 1;
        this.markdown.shift();
      }
      if(this.markdown.length - 1 != this.currentIndex) {
        this.markdown = this.markdown.filter((_, index) => {
          return this.currentIndex >= index;
        })
      }
      this.markdown.push(v)
      this.currentIndex += 1;
      this.$emit('currentText', this.currentText);
    },
    insertText(v) {
      if(this.currentIndex == 200) {
        this.currentIndex -= 1;
        this.markdown.shift();
      }
      if(this.markdown.length - 1 != this.currentIndex) {
        this.markdown = this.markdown.filter((_, index) => {
          return this.currentIndex >= index;
        })
      }
      this.markdown.push(v)
      this.currentIndex += 1;
    },
    preventUndoRedo(e){
      if(e.key == 'y') {
        if(e.ctrlKey) {
          // win Redo
          e.preventDefault();
          if(this.currentIndex == this.markdown.length - 1) return;
          this.currentIndex += 1;
          return;
        }
      }
      if(e.key == 'z') {
        if(e.shiftKey && e.metaKey) {
          // mac Redo
          e.preventDefault();
          if(this.currentIndex == this.markdown.length - 1) return;
          this.currentIndex += 1;
          return;
        }
        if(e.metaKey || e.ctrlKey) {
          // mac win Undo
          e.preventDefault();
          if(this.currentIndex == 0) return;
          this.currentIndex -= 1;
          return;
        }
      }
    },
    changeToHalf(str) {
      return str.replace(/[Ａ-Ｚａ-ｚ０-９]/g, (s) => {
        return String.fromCharCode(s.charCodeAt(0) - 0xFEE0);
      });
    },
    bold() {
      const insertText = '**';
      const newText = this.betweenSelectedText(insertText, insertText);
      this.insertText(newText);
    },
    italic() {
      const insertText = '*';
      const newText = this.betweenSelectedText(insertText, insertText);
      this.insertText(newText);
    },
    underline() {
      const insertTextStart = '<u>';
      const insertTextEnd = '</u>';
      const newText = this.betweenSelectedText(insertTextStart, insertTextEnd);
      this.insertText(newText);
    },
    centering() {
      const insertTextStart = `<div align="center">`
      const insertTextEnd = `</div>`
      const newText = this.betweenSelectedText(insertTextStart, insertTextEnd);
      this.insertText(newText);
    },
    betweenSelectedText(insertTextstart, insertTextEnd) {
      const {text, selectionStart, selectionEnd, textEnd} = this.getSelectedArea();
      const insertTextLength = insertTextstart.length;
      const newSelectionStart = selectionStart + insertTextLength;
      this.textarea.focus();
      this.$nextTick(() => {
        this.textarea.setSelectionRange(newSelectionStart, newSelectionStart);
      })
      const before = text.substring(0, selectionStart);
      const selectedText = text.substring(selectionStart, selectionEnd);
      const after = text.substring(selectionEnd, textEnd);
      const addedText = before + insertTextstart + selectedText + insertTextEnd + after
      return addedText;
    },
    getSelectedArea() {
      const textarea = this.$refs.detail.$refs.input;
      const text = textarea.value;
      let selectionStart = textarea.selectionStart;
      const selectionEnd = textarea.selectionEnd;
      const textEnd = text.length;
      if (selectionStart === undefined) {
        selectionStart = 0
      }
      return {text, selectionStart, selectionEnd, textEnd};
    },
  }
})
</script>