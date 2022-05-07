<template>
  <div @mouseup="endDrag">
    <v-file-input
      accept="image/*"
      label="File input"
      filled
      @change="onChangeFile"
      prepend-icon=""
      clearable
    ></v-file-input>
    <div>
      <div style="position: relative;margin: auto" @mousemove="dragSize" :style="{width: imgWidth + 'px', height: imgHeight + 'px'}">
        <div class="cropper pa-1" ref="cropper" :style="cropperPosition" @mousedown.self="startDragSize" v-if="image_data">
          <div style="width: 100%;height: 100%;" :style="{cursor: isSizeCursor ? 'all-scroll' : 'grab'}" @mousedown="startDragPosition" @mousemove="dragPosition"></div>
        </div>
        <canvas ref="canvas"></canvas>
      </div>
      <div class="pa-2">
        <v-btn block color="indigo" dark v-if="image_data" @click="onClickCropImage">
          <v-icon small dark>mdi-content-cut</v-icon>
          切り取り
        </v-btn>
      </div>
      <div class="pa-2 d-flex justify-space-between">
        <img style="max-width: 200px;width: 100%" :src="currentImage" alt="">
        <img style="max-width: 200px;width: 100%" :src="value" alt="" v-if="value">
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
    width: {
      type: Number
    },
    currentImage: {
      type: String
    },
    value: {
      type: String
    }
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
    startDragSize(event: MouseEvent) {
      this.isSizeCursor = true;
      const {x, y, clientRect} = this.getPosition(event);
      this.sizeX = x;
      this.sizeY = y;
      this.$root.$el.addEventListener('mouseup',this.endDrag, {once:true})
    },
    dragSize(event: MouseEvent) {
      const minWidth = 50;
      const maxWidth = this.imgWidth < this.imgHeight ? this.imgWidth : this.imgHeight;;
      if(!this.isSizeCursor) return;
      const {x, y, clientRect} = this.getPosition(event);
      this.cropper.maxWidth -= this.sizeX - x;
      if(minWidth >= this.cropper.maxWidth) {
        this.cropper.maxWidth = minWidth;
      }
      if(maxWidth <= this.cropper.maxWidth) {
        this.cropper.maxWidth = maxWidth;
      }
      this.sizeX = x;
      this.cropper.height = this.cropper.maxWidth;
    },
    startDragPosition(event: MouseEvent) {
      this.isPositionCursor = true;
      // 要素内におけるクリック位置を計算
      const {x, y} = this.getPosition(event);
      this.positionX = x;
      this.positionY = y;
      this.$root.$el.addEventListener('mouseup',this.endDrag, {once:true})
    },
    dragPosition(event: MouseEvent) {
      if(!this.isPositionCursor) return;
      const {x, y, clientRect} = this.getPosition(event);
      const maxX = this.imgWidth - clientRect.width;
      const maxY = this.imgHeight - clientRect.height;
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
      this.isPositionCursor = false;
      this.isSizeCursor = false;
    },
    async onChangeFile(e: File) {
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
      this.imgWidth = this.width <= image.width ? this.width : image.width;
      this.imgHeight = image.height * (this.imgWidth / image.width);
      const canvas = this.$refs.canvas as HTMLCanvasElement;
      canvas.width = this.imgWidth;
      canvas.height = this.imgHeight;
      this.cropper.maxWidth = this.imgWidth < this.imgHeight ? this.imgWidth : this.imgHeight;
      this.cropper.height = this.cropper.maxWidth;
      this.image_data = this.toCanvas(image, canvas, this.imgWidth, this.imgHeight)
    },
    async onClickCropImage() {
      const image  = await this.resize(this.image_data) as HTMLImageElement;
      const canvas = document.createElement('canvas');
      const ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
      canvas.width = this.cropper.height;
      canvas.height = this.cropper.height;
      ctx.drawImage(image, -this.cropper.left, -this.cropper.top);
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
  }
})
</script>

<style lang="scss">
  .cropper {
    position: absolute;
    width: 100%;
    border: solid black 2px;
    border-style: dashed;
    cursor: all-scroll;
  }
</style>