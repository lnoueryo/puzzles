<template>
  <div>
    <div v-for="(comment, i) in comments" :key="i">
      <div :class="['d-flex', 'comment', 'justify-space-between', {'select' : selectedComment.id === comment.id}]" @click="selectComment(comment, hierarchy)">
        <div class="d-flex align-center w100">
          <v-btn icon @click="checkInput('tree' + comment.id)" v-if="isReadyArr(comment.replies)">
            <v-icon :class="[{'active': isOpen(comment.id)}, 'icon']">mdi-triangle-small-down</v-icon>
          </v-btn>
          <input
            :ref="'tree' + comment.id"
            :id="'tree' + comment.id"
            type="checkbox"
            :value="comment.id"
            v-model="selectedComments"
            class="hide-display"
          >
          <user-cell :class="{'ml-9' : isEmptyArr(comment.replies)}" :user="comment.user"></user-cell>
          <div>{{ comment.content }}</div>
          <div class="mla" v-if="comment.user_id == user.id">
            <v-btn @click="finishEditComment()" v-if="editMode">終了</v-btn>
            <v-btn @click="editComment(comment, hierarchy)" v-else>編集</v-btn>
            <v-btn @click="deleteComment(comment)">削除</v-btn>
          </div>
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
          @addID="selectedComments.push($event)"
          :editMode="editMode"
        ></tree-comments>
      </keep-alive>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { isEmptyArr, isReadyArr, checkStatus } from '~/modules/utils'
import * as model from '~/modules/model'
declare module 'vue/types/vue' {
  interface Vue {
    getHierarchy: (obj: {}) => void;
    comments: [];
    createIDArray: (comments: model.Comment[]) => number[]
    selectComment: (comment: model.Comment, index: number) => void
  }
}
export default Vue.extend({
  props: {
    comments: [],
    hierarchy: Number,
    selectedComment: Object,
    editMode: Boolean
  },
  data: () => ({
    selectedComments: [],
    commentsIDs: [] as number[],
  }),
  computed: {
    isEmptyArr,
    isReadyArr,
    checkStatus,
    nextHierarchy() {
      return this.hierarchy + 1;
    },
    user() {
      return this.$store.getters.user;
    },
  },
  created() {
    this.commentsIDs = this.createIDArray(this.comments);
  },
  methods: {
    createIDArray(comments: model.Comment[]): number[] {
      return comments.map((comment: model.Comment) => comment.id as number) as number[];
    },
    selectComment(comment: model.Comment, index: number) {
      const newComment = JSON.parse(JSON.stringify(comment))
      if(this.selectedComment.id == newComment.id) newComment.id = 0
      const obj = {...newComment, ...{index}}
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
    },
    editComment(comment: model.Comment, index: number) {
      this.$store.dispatch('comment/editMode', true)
      this.$store.commit('comment/content', comment.content)
      this.selectComment(comment, index)
    },
    finishEditComment() {
      this.$store.dispatch('comment/editMode', false)
    },
    async deleteComment(comment: model.Comment) {
      let response;
      try {
        response = await this.$store.dispatch('comment/deleteComment', comment);
      } catch (error: any) {
        response = error.response;
      } finally {
        console.log(response)
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {},
        () => {
          // this.loading = false;
          alert('エラーです。');
        })
      }
    }
  }
})
</script>

<style lang="scss" scoped>
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