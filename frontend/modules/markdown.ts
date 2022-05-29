export class Markdown {
  public markdown = ['']
  public text: string = ''
  public currentIndex = 0
  public textarea
  constructor(textarea: HTMLTextAreaElement) {
    this.textarea = textarea;
  }
  insertText(v: string) {
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
  }
  preventUndoRedo(e: KeyboardEvent) {
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
  }
  changeToHalf(str: string) {
    return str.replace(/[Ａ-Ｚａ-ｚ０-９]/g, (s) => {
      return String.fromCharCode(s.charCodeAt(0) - 0xFEE0);
    });
  }
  bold() {
    const insertText = '**';
    const newText = this.betweenSelectedText(insertText, insertText);
    this.insertText(newText);
  }
  italic() {
    const insertText = '*';
    const newText = this.betweenSelectedText(insertText, insertText);
    this.insertText(newText);
  }
  underline() {
    const insertTextStart = '<u>';
    const insertTextEnd = '</u>';
    const newText = this.betweenSelectedText(insertTextStart, insertTextEnd);
    this.insertText(newText);
  }
  centering() {
    const insertTextStart = `<div align="center">`
    const insertTextEnd = `</div>`
    const newText = this.betweenSelectedText(insertTextStart, insertTextEnd);
    this.insertText(newText);
  }
  betweenSelectedText(insertTextStart: string, insertTextEnd: string) {
    const {text, selectionStart, selectionEnd, textEnd} = this.getSelectedArea();
    const insertTextLength = insertTextStart.length;
    const newSelectionStart = selectionStart + insertTextLength;
    this.textarea.focus();
    window.$nuxt.$nextTick(() => {
      this.textarea.setSelectionRange(newSelectionStart, newSelectionStart);
    })
    const before = text.substring(0, selectionStart);
    const selectedText = text.substring(selectionStart, selectionEnd);
    const after = text.substring(selectionEnd, textEnd);
    const addedText = before + insertTextStart + selectedText + insertTextEnd + after
    return addedText;
  }
  getSelectedArea() {
    const textarea = this.textarea;
    const text = textarea.value;
    let selectionStart = textarea.selectionStart;
    const selectionEnd = textarea.selectionEnd;
    const textEnd = text.length;
    if (selectionStart === undefined) {
      selectionStart = 0
    }
    return {text, selectionStart, selectionEnd, textEnd};
  }
}
