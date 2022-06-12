<template>
  <div v-if="pageReady">
    <v-row justify="center">
      <v-col cols="12" sm="10" md="10" lg="10">
        <form-task v-model="selectedTask" @submit="onClickSubmit" :loading="loading">
          <template v-slot:back>
            戻る
          </template>
          <template v-slot:submit>
            作成
          </template>
        </form-task>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isReadyObj, isEmptyArr, isEmptyObj, checkStatus, changeToDateISOFormat, changeToTimeStampFormat } from '~/modules/utils'
import * as model from '~/modules/model'
declare module 'vue/types/vue' {
  interface Vue {
    changeToHalf: (string: string) => any;
    isNumber: (v: any) => boolean;
    validateForm: (v: model.Task) => string;
  }
}
interface Rule {
  length: Function
  required: Function
  requiredSelect: Function
  isNumber: Function
}
export default Vue.extend({
  data: () => ({
    selectedTask: {} as model.Task,
    loading: false,
  }),
  computed: {
    ...mapGetters('task', [
      'statuses',
      'types',
      'priorities',
    ]),
    ...mapGetters([
      'user',
      'project',
    ]),
    isReadyObj,
    isEmptyArr,
    checkStatus,
    isEmptyObj,
    changeToDateISOFormat,
    changeToTimeStampFormat,
    pageReady() {
      return this.isReadyObj(this.project);
    },
    taskForm() {
      const additionalInfo = {
        id: this.$route.params.key,
        assigner_id: this.user.id,
        project_id: Number(this.$route.params.id),
      }
      const cleansedData = {
        estimated_time: Number(this.selectedTask.estimated_time),
        start_time: this.selectedTask.start_time ? new Date(this.selectedTask.start_time) : null,
        deadline: this.selectedTask.deadline ? new Date(this.selectedTask.deadline) : null,
      }
      const assigner = this.user;
      const assignee = this.project.authority_users.find((user: model.ProjectAuthority) => user.user_id === this.selectedTask.assignee_id).user;
      const status = this.statuses.find((status: {id: number}) => status.id === this.selectedTask.status_id);
      const type = this.types.find((type: {id: number}) => type.id === this.selectedTask.type_id);
      const priority = this.priorities.find((priority: {id: number}) => priority.id === this.selectedTask.priority_id);
      const field = this.project.fields.find((field: model.Field) => field.id === this.selectedTask.field_id) || {};
      const milestone = this.project.milestones.find((milestone: model.Milestone) => milestone.id === this.selectedTask.milestone_id) || {};
      const actual_time = 0
      const requiredDataforDisplay = {
        assigner,
        assignee,
        status,
        type,
        field,
        milestone,
        priority,
        actual_time,
        comments: [],
      }
      const newTask = {...this.selectedTask, ...additionalInfo, ...cleansedData, ...requiredDataforDisplay}
      return newTask;
    },
  },
  created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.user)) return;
      clearInterval(timer);
      const additionalInfo = {
        assignee_id: this.user.id,
        status_id: 1,
        type_id: 1,
        priority_id: 2,
        estimated_time: 0,
        deadline: this.changeToDateISOFormat('', 5)
      }
      this.selectedTask = {...this.selectedTask, ...additionalInfo}
    }, 100)
  },
  methods: {
    async onClickSubmit() {
      this.loading = true;
      // const validation = this.validateForm(this.taskForm);
      // if(validation) {
      //   return;
      // }

      let response
      try {
        response = await this.$store.dispatch('task/createTask', this.taskForm);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'project-id', params: {id: this.$route.params.id}});
        }, () => {
          this.loading = false;
          alert('エラーです。');
        })
      }
    },
    validateForm(form: model.Task) {
      if(!form.title) {
        return '課題のタイトルが空白です'
      }
      // if(!form.field_id) {
      //   return '分野を選択してください'
      // }
      // if(!form.milestone_id) {
      //   return 'マイルストーンを選択してください'
      // }
      if(!form.deadline) {
        return '期日を選択してください'
      }
      return;
    },
  }
})
</script>

<style lang="scss" scoped>
.v-application--is-ltr .v-text-field .v-label {
    color: cadetblue;
}
</style>
