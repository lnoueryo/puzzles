<template>
  <div>
    <v-file-input
      accept="image/*"
      label="File input"
      filled
      @change="onChangeFile"
      prepend-icon=""
      clearable
    ></v-file-input>
    <div>
      <div class="relative ma" @mouseleave="endDrag" @mousemove="dragMove" :style="{width: toStyleSize(imgWidth) + 'px', height: toStyleSize(imgHeight) + 'px'}">
        <div class="cropper pa-1" ref="cropper" :style="cropperPosition" @mousedown.self="startDragSize" @mouseup="endDrag" v-if="image_data">
          <div class="w100 h100" :style="{cursor: isSizeCursor ? 'all-scroll' : 'grab'}" @mousedown="startDragPosition" @mouseup="endDrag"></div>
        </div>
        <canvas ref="canvas" @mouseup="endDrag" @mousemove="dragMove" :style="{width: '100%', maxWidth: this.width + 'px'}"></canvas>
      </div>
      <div class="pa-2">
        <v-btn block color="#295caa" dark v-if="image_data" @click="onClickCropImage">
          <v-icon small dark>mdi-content-cut</v-icon>
          切り取り
        </v-btn>
      </div>
      <div class="pa-2 d-flex justify-space-between">
        <div class="image-container" v-if="currentImage&&!value">
          <v-img class="w100" :src="currentImage">
            <template v-slot:placeholder>
              <v-row class="fill-height ma-0" align="center" justify="center">
                <v-progress-circular indeterminate color="grey lighten-5" />
              </v-row>
            </template>
          </v-img>
        </div>
        <div class="image-container" v-if="value">
          <v-img class="w100" :src="value">
            <template v-slot:placeholder>
              <v-row class="fill-height ma-0" align="center" justify="center">
                <v-progress-circular indeterminate color="grey lighten-5" />
              </v-row>
            </template>
          </v-img>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
// import { resizeFile } from '~/modules/utils'
declare module 'vue/types/vue' {
  interface Vue {
    addPixel: () => {};
  }
}
export default Vue.extend({
  props: {
    width: Number,
    pixel: Number,
    currentImage: String,
    value: String,
    ratio: String
  },
  // props: ['width', 'currentImage'],
  data: () => ({
    image: '',
    image_data: '',
    newWidth: 0,
    newHeight: 0,
    imgWidth: 0,
    imgHeight: 0,
    isPositionCursor: false,
    isSizeCursor: false,
    positionX: 0,
    positionY: 0,
    sizeX: 0,
    sizeY: 0,
    cropper: {
      originalPosition: 0,
      mainPosition: 'x' as 'x' | 'y',
      maxWidth: 0,
      height: 0,
      top: 0,
      left: 0,
    }
  }),
  computed: {
    cropperPosition() {
      return this.addPixel()
    },
  },
  methods: {
    dragMove(e: MouseEvent) {
      this.dragSize(e);
      this.dragPosition(e);
    },
    startDragSize(event: MouseEvent) {
      this.isSizeCursor = true;
      const {x, y, clientRect} = this.getPosition(event);
      this.sizeX = x;
      this.sizeY = y;
    },
    dragSize(event: MouseEvent) {
      if(!this.isSizeCursor) return;
      const minWidth = 50;
      const {x, y, clientRect} = this.getPosition(event);
      const cursorPosition = clientRect[this.cropper.mainPosition] - this.cropper.originalPosition;
      const maxWidth = this.calculateHeight(this.imgWidth) < this.imgHeight ? this.toStyleSize(this.imgWidth - cursorPosition) : this.toStyleSize(this.calculateWidth(this.imgHeight - cursorPosition));
      this.cropper.maxWidth -= this.sizeX - x;
      if(minWidth >= this.cropper.maxWidth) {
        this.cropper.maxWidth = minWidth;
      }
      if(maxWidth <= this.cropper.maxWidth) {
        this.cropper.maxWidth = maxWidth;
      }
      this.sizeX = x;
      this.cropper.height = this.calculateHeight(this.cropper.maxWidth);
    },
    startDragPosition(event: MouseEvent) {
      this.isPositionCursor = true;
      // 要素内におけるクリック位置を計算
      const {x, y} = this.getPosition(event);
      this.positionX = x;
      this.positionY = y;
      // this.$root.$el.addEventListener('mouseup',this.endDrag, {once:true})
    },
    dragPosition(event: MouseEvent) {
      if(!this.isPositionCursor) return;
      if(event.stopPropagation) event.stopPropagation();
      if(event.preventDefault) event.preventDefault();
      event.cancelBubble=true;
      const {x, y, clientRect} = this.getPosition(event);
      const maxX = this.toStyleSize(this.imgWidth) - clientRect.width;
      const maxY = this.toStyleSize(this.imgHeight) - clientRect.height;
      this.cropper.top -= this.positionY - y;
      this.cropper.left -= this.positionX - x;
      if(this.cropper.top <= 0) {
        this.cropper.top = 0;
        this.positionY = y;
      }
      if(this.cropper.top >= maxY) {
        this.cropper.top = maxY;
        this.positionY = y;
      }
      if(this.cropper.left <= 0) {
        this.cropper.left = 0;
        this.positionX = x;
      }
      if(this.cropper.left >= maxX) {
        this.cropper.left = maxX;
        this.positionX = x;
      }
    },
    endDrag() {
      window?.getSelection()?.removeAllRanges();
      this.isPositionCursor = false;
      this.isSizeCursor = false;
    },
    async onChangeFile(e: File) {
      this.cropper = {
        originalPosition: 0,
        mainPosition: 'x',
        maxWidth: 0,
        height: 0,
        top: 0,
        left: 0,
      }
      if(!e) {
        const canvas = this.$refs.canvas as HTMLCanvasElement;
        const ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        canvas.width = 0
        canvas.height = 0
        this.imgWidth = 0
        this.imgHeight = 0
        this.image_data = '';
        this.$emit('input', '')
        return;
      };
      const image = await this.resizeFile(e) as HTMLImageElement;
      this.imgWidth = this.pixel;
      // this.imgWidth = this.width <= image.width ? this.width : image.width;
      this.imgHeight = image.height * (this.imgWidth / image.width);
      const canvas = this.$refs.canvas as HTMLCanvasElement;
      canvas.width = this.imgWidth;
      canvas.height = this.imgHeight;
      this.cropper.maxWidth = this.calculateHeight(this.imgWidth) < this.imgHeight ? this.toStyleSize(this.imgWidth) : this.toStyleSize(this.calculateWidth(this.imgHeight));
      this.cropper.height = this.calculateHeight(this.cropper.maxWidth);
      this.image_data = this.toCanvas(image, canvas, this.imgWidth, this.imgHeight)
      this.$nextTick(() => {
        const cropper = this.$refs.cropper as HTMLDivElement
        const clientRect = cropper.getBoundingClientRect();
        this.cropper.originalPosition = this.calculateHeight(this.imgWidth) <= this.imgHeight ? clientRect.x : clientRect.y;
        this.cropper.mainPosition = this.calculateHeight(this.imgWidth) <= this.imgHeight ? 'x' : 'y';
      })
    },
    async onClickCropImage() {
      const image  = await this.resize(this.image_data) as HTMLImageElement;
      const canvas = document.createElement('canvas');
      const ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
      canvas.width = this.toImageSize(this.cropper.maxWidth);
      canvas.height = this.toImageSize(this.cropper.height);
      ctx.drawImage(image, -this.toImageSize(this.cropper.left), -this.toImageSize(this.cropper.top));
      const imgType = image.src.substring(5, image.src.indexOf(';'));
      this.$emit('input', canvas.toDataURL(imgType));
    },
    async resizeFile (e: File) {
      const reader = await this.fileReader(e) as FileReader;
      const resizedImage = await this.resize(reader.result as string);
      return resizedImage;
    },
    fileReader (e: File) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(e);
        reader.onload = () => {
          resolve(reader)
        }
        reader.onerror = reject
      })
    },
    resize (src: string) {
      return new Promise((resolve, reject) => {
        const image = new Image();
        image.src = src;
        image.onload = () => {
          resolve(image);
        }
        image.onerror = reject
      })
    },
    toCanvas(image: HTMLImageElement, canvas: HTMLCanvasElement, width: number, height: number, startX: number = 0, startY: number = 0) {
      const ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
      ctx.drawImage(image, startX, startY, width, height);
      const imgType = image.src.substring(5, image.src.indexOf(';'));
      const resizedImage = canvas.toDataURL(imgType)
      return resizedImage;
    },
    getPosition(event: MouseEvent) {
      const clickX = event.pageX;
      const clickY = event.pageY;
      const cropper = this.$refs.cropper as HTMLDivElement
      const clientRect = cropper.getBoundingClientRect();
      const positionX = clientRect.left + window.pageXOffset;
      const positionY = clientRect.top + window.pageYOffset;
      const x = clickX - positionX;
      const y = clickY - positionY;
      return {x, y, clientRect};
    },
    addPixel() {
      const cropper = JSON.parse(JSON.stringify(this.cropper))
      cropper.maxWidth += 'px'
      cropper.height += 'px'
      cropper.top += 'px'
      cropper.left += 'px'
      return cropper
    },
    calculateWidth(height: number): number {
      if(!this.ratio) return height;
      const colonIndex = this.ratio.indexOf(':')
      const firstNum = Number(this.ratio.substr(0, colonIndex));
      const secondNum = Number(this.ratio.substr(colonIndex + 1));
      const width = height * firstNum / secondNum;
      return width;
    },
    calculateHeight(width: number): number {
      if(!this.ratio) return width;
      const colonIndex = this.ratio.indexOf(':')
      const firstNum = Number(this.ratio.substr(0, colonIndex));
      const secondNum = Number(this.ratio.substr(colonIndex + 1));
      const height = width * secondNum / firstNum
      return height;
    },
    toStyleSize(num: number): number {
      return num * this.width / this.pixel;
    },
    toImageSize(num: number): number {
      return num * this.pixel / this.width;
    },
  }
})
</script>

<style lang="scss" scoped>
  .cropper {
    position: absolute;
    width: 100%;
    border: solid black 2px;
    border-style: dashed;
    cursor: all-scroll;
  }
  .image-container {
    max-width: 450px;
    margin: auto
  }
</style>