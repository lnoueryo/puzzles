<template>
  <div>
    <div v-for="(comment, i) in comments" :key="i">
      <div :class="['d-flex', 'comment', 'justify-space-between', {'select' : selectedComment.id === comment.id}]" @click="selectComment(comment.id, hierarchy)">
        <div class="d-flex align-center">
          <v-btn icon @click="checkInput('tree' + comment.id)" v-if="isReadyArr(comment.replies)">
            <v-icon :class="[{'active': isOpen(comment.id)}, 'icon']">mdi-triangle-small-down</v-icon>
          </v-btn>
          <input
            :ref="'tree' + comment.id"
            :id="'tree' + comment.id"
            type="checkbox"
            :value="comment.id"
            v-model="selectedComments"
            style="display: none"
          >
          <user-cell :class="{'ml-9' : isEmptyArr(comment.replies)}" :user="comment.user"></user-cell>
          <div>{{ comment.content }}</div>
        </div>
      </div>
      <keep-alive>
        <tree-comments
          :key="comment.id"
          :id="'tree' + comment.id"
          class="ml-12"
          :comments="comment.replies"
          :hierarchy="nextHierarchy"
          @index="getHierarchy($event, comment.id)"
          v-if="isOpen(comment.id)"
          :selectedComment="selectedComment"
        ></tree-comments>
      </keep-alive>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { isEmptyArr, isReadyArr } from '~/modules/utils'
declare module 'vue/types/vue' {
  interface Vue {
    getHierarchy: (obj: {}) => void;
  }
}
export default Vue.extend({
  props: {
    comments: {},
    hierarchy: {
      type: Number
    },
    selectedComment: {
      type: Object
    }
  },
  data: () => ({
    selectedComments: []
  }),
  computed: {
    isEmptyArr,
    isReadyArr,
    nextHierarchy() {
      return this.hierarchy + 1;
    },
  },
  methods: {
    selectComment(id: number, index: number) {
      if(this.selectedComment.id == id) id = 0
      const obj = {...{id}, ...{index}}
      this.getHierarchy(obj)
    },
    getHierarchy(obj: {}, parent=0) {
      this.$emit('index', {...obj, ...{parent}})
    },
    isOpen(key: number) {
      return this.selectedComments.some((commentKey) => {
        return commentKey === key
      })
    },
    checkInput(el: string) {
      (this.$refs[el] as HTMLInputElement[])[0].click()
    }
  }
})
</script>

<style lang="scss" ascoped>
  .select {
    background-color: #295daa5e
  }
  .icon {
    transform: rotate(-90deg);
  }
  .active {
    transform: rotate(0deg);
  }

</style>