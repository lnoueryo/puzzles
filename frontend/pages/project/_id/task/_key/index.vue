<template>
  <div v-if="isReadyObj(task)">
    <form-card class="my-6" style="max-width: 70%;padding-bottom: 150px">
      <div class="d-flex justify-space-between align-center pa-4 pl-6">
        <div>
          <small>{{ task.key }}</small>
          <h3>{{ task.title }}</h3>
        </div>
        <div class="d-flex px-3 align-center" style="font-size: 13px">
          <div class="px-7">
            <div>{{ task.status }}</div>
          </div>
          <div class="px-3">
            <div class="text-center">期日</div>
            <div>{{ task.deadline }}</div>
          </div>
          <div class="pl-4">
            <v-btn :to="{name: 'project-id-task-key-edit', params: {id: $route.params.id, key: $route.params.key}}">編集</v-btn>
          </div>
        </div>
      </div>
      <v-divider></v-divider>
      <div class="d-flex align-center justify-space-between pa-2" style="font-size: 13px">
        <user-cell :styleValue="{}" :user="task.assigner"></user-cell>
        <div class="d-flex text-center px-3">
          <div class="px-3">
            <div>分野</div>
            <div>{{ task.field }}</div>
          </div>
          <div class="px-3">
            <div>マイルストーン</div>
            <div>{{ task.milestone }}</div>
          </div>
          <div class="px-3">
            <div>作成日</div>
            <div>{{ task.created_at }}</div>
          </div>
          <div class="px-3">
            <div>更新日</div>
            <div>{{ task.updated_at }}</div>
          </div>
        </div>
      </div>
      <div class="px-6">
        <div class="d-flex py-2" style="width: 100%">
          <div style="width: 100%">
            <div class="mb-2">詳細</div>
            <div class="pa-3" style="min-height: 300px;background-color: #303030;border-radius: 5px;" v-html="task.detail"></div>
          </div>
        </div>
        <div class="d-flex py-2" style="width: 100%">
          <div style="width: 100%">
            <div class="mb-2">コメント</div>
              <tree-comments
                v-if="isReadyArr(task.comments)"
                :comments="task.comments"
                :hierarchy="0"
                :selectedComment="selectedComment"
                @index="receiveCommentKey"
              ></tree-comments>
              <div v-else>
                <div class="pa-2" style="width: 100%;min-height: 100px;background-color: #303030;border-radius: 5px;">なし</div>
              </div>
          </div>
        </div>
      </div>
    </form-card>
    <v-bottom-sheet transition v-model="sheet" persistent inset hide-overlay>
      <v-card tile style="height: 200px">
        <div class="d-flex">
          <div style="width: 50%">
            <v-btn block　@click="onCreateComment">送信</v-btn>
            <v-textarea label="コメント" outlined v-model="comment"></v-textarea>
          </div>
          <div class="px-2" style="width: 50%">
            <div class="d-flex">
              <div class="d-flex align-center justify-space-between" style="width: 50%">
                <div class="text-center" style="width: 25%">状態</div>
                <v-select
                  ref="status"
                  :items="statuses"
                  item-text="name"
                  item-value="id"
                  style="width: 25%;"
                  flat
                  solo
                  hide-details
                  v-model="selectedTask.status_id"
                ></v-select>
              </div>
              <div class="d-flex align-center justify-space-between">
                <div class="text-center" style="width: 30%">実働</div>
                <v-text-field
                  class="d-flex justify-center text-center px-4"
                  flat
                  style="width: 30%;"
                  v-model="selectedTask.actual_time"
                >
                </v-text-field>
                <div class="text-center" style="width: 30%">時間</div>
              </div>
            </div>
            <div class="d-flex">
              <div class="d-flex align-center" style="width: 20%">
                <div class="text-center">担当者</div>
              </div>
              <div class="d-flex align-center justify-space-between" style="width: 80%">
                <user-cell :styleValue="{}" :user="task.assignee"></user-cell>
                <v-btn @click="onUpdateTask">更新</v-btn>
              </div>
            </div>
            <!-- <div>推定時間</div>
            <div>時間</div> -->
          </div>
        </div>
          <!-- <v-list>
            <v-list-item>
              <div>推定時間</div>
              <div>状態</div>
              {{selectedTask.status_id}}
              <v-select
                ref="status"
                v-model="selectedTask.status_id"
                :items="statuses"
                item-text="name"
                item-value="id"
                label="現在の状態"
                style="width: 0px;"
                flat
                solo
                hide-details
              ></v-select>
              <div>実働</div>
              <v-text-field
                solo
                flat
                hide-details
                style="width: 0px;text-align: center;"
                v-model="selectedTask.actual_time"
              >
              </v-text-field>
              <div>時間</div>
              <div>担当者</div>
              <user-cell :styleValue="{}" :user="task.assignee"></user-cell>
              <v-btn>更新</v-btn>
            </v-list-item>
          </v-list> -->
      </v-card>
    </v-bottom-sheet>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isEmptyObj, isEmptyArr, checkStatus, isReadyObj, isReadyArr } from '~/modules/utils'
import * as lib from '~/modules/store'
declare module 'vue/types/vue' {
  interface Vue {
    addReply: (replies: lib.Comment[], commentKey: {id: number, index: number}, currentIndex: number) => lib.Comment;
  }
}
export default Vue.extend({
  data: () => ({
    sheet: true,
    loading: false,
    selectedTask: {
      actual_time: 0,
      status_id: 1,
    },
    comment: '',
    comments: [],
  }),
  computed: {
    ...mapGetters('task', [
      'task',
      'statuses',
      'types',
      'priorities',
      'selectedComment',
    ]),
    ...mapGetters([
      'user',
      'projectAuthority'
    ]),
    isReadyObj,
    isReadyArr,
    isEmptyObj,
    checkStatus,
    taskForm() {
      const additionalInfo = {
        id: this.$route.params.key,
        assigner_id: this.user.id,
        project_id: Number(this.$route.params.id),
      }
      const cleansedData = {
        estimated_time: Number(this.task.estimated_time),
        start_time: this.task.start_time ? new Date(this.task.start_time) : null,
        deadline: this.task.deadline ? new Date(this.task.deadline) : null,
      }
      const status = this.statuses.find((status: {id: number}) => status.id === this.selectedTask.status_id);
      const type = this.types.find((type: {id: number}) => type.id === this.task.type_id);
      const priority = this.priorities.find((priority: {id: number}) => priority.id === this.task.priority_id);
      const field = this.projectAuthority.project.fields.find((field: lib.Field) => field.id === this.task.field_id) || {};
      const milestone = this.projectAuthority.project.milestones.find((milestone: lib.Milestone) => milestone.id === this.task.milestone_id) || {};
      const actual_time = Number(this.selectedTask.actual_time);
      const created_at = new Date(this.task.created_at);
      const status_id = this.selectedTask.status_id;
      const requiredDataforDisplay = {
        status,
        type,
        field,
        milestone,
        priority,
        actual_time,
        comments: [],
        created_at,
        status_id,
      }
      const newTask = {...this.task, ...additionalInfo, ...cleansedData, ...requiredDataforDisplay}
      return newTask;
    },
    newCommentForm() {
      return {
        content: this.comment,
        task_id: this.task?.id,
        user_id: this.user.id,
        parent_id: this.selectedComment.id,
      }
    }
  },
  created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.task)) return;
      clearInterval(timer)
      this.selectedTask.actual_time = this.task.actual_time
      this.selectedTask.status_id = this.task.status_id
      console.log(this.task.status_id)
      console.log(this.selectedTask)
    }, 100)
  },
  methods: {
    async onUpdateTask() {
      let response
      try {
        response = await this.$store.dispatch('task/updateTask', this.taskForm);
      } catch (error: any) {
        response = error.response;
      } finally {
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'project-id-task', params: {id: this.$route.params.id}});
        }, () => {
          this.loading = false;
          alert('エラーです。');
        })
      }
    },
    async onCreateComment() {
      let response;
      try {
        response = this.$store.dispatch('task/createComment', this.newCommentForm);
      } catch (error: any) {
        response = error.response;
      } finally {
        this.checkStatus(response.status, () => {},
        () => {
          this.loading = false;
          alert('エラーです。');
        })
      }
    },
    receiveCommentKey(commentKey: {parent: number}) {
      this.$store.commit('task/selectComment', commentKey)
    },
    onCreateReply() {
      // const comment = this.task.comments.find((reply: lib.Comment) => {
      //   let id = this.selectedComment.parent;
      //   if(this.selectedComment.parent === 0) id = this.selectedComment.id;
      //   return reply.id === id;
      // });
      // let copyComment = JSON.parse(JSON.stringify(comment));
      // if(this.selectedComment.parent === 0) {
      //   copyComment.replies.push(this.newCommentForm);
      // } else {
      //   copyComment.replies = this.addReply(comment.replies, this.selectedComment, 0);
      // }
      console.log(this.selectedComment)
      let response;
      try {
        response = this.$store.dispatch('task/updateComment', this.newCommentForm);
      } catch (error: any) {
        response = error.response;
      } finally {
        this.checkStatus(response.status, () => {},
        () => {
          this.loading = false;
          alert('エラーです。');
        })
      }
    },
    addReply(replies: lib.Comment[], commentKey: {id: number, index: number}, currentIndex: number) {
      if(replies.length == 0 || commentKey.index < currentIndex) return replies;
      return replies.map((reply) => {
        const newReply = JSON.parse(JSON.stringify(reply))
        if(commentKey.id == reply.id) {
          const newComment = this.newCommentForm as lib.Comment;
          newComment.parent_id = false;
          newReply.replies.push(newComment);
        } else if(reply.replies.length != 0) {
          console.log(this.addReply(reply.replies, commentKey, currentIndex + 1))
          newReply.replies = this.addReply(reply.replies, commentKey, currentIndex + 1);
        }
        return newReply;
      }) as lib.Comment[]
    },
  }
})
</script>
