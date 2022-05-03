<template>
  <div v-if="pageReady">
    <v-row justify="center">
      <v-col cols="12" sm="10" md="10" lg="10">
        <task-form v-model="selectedTask" @submit="dialog = true">
          <template v-slot:back>
            戻る
          </template>
          <template v-slot:submit>
            更新
          </template>
        </task-form>
      </v-col>
    </v-row>
    <update-dialog v-model="dialog" :form="dialogForm" @submit="onClickSubmit" @loading="loading = $event">
      {{ taskForm.key }}
    </update-dialog>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isReadyObj, isEmptyArr, checkStatus, isEmptyObj, changeToISOFormat, changeToTimeStampFormat, changeToDateISOFormat } from '~/modules/utils'
import * as lib from '~/modules/store'

declare module 'vue/types/vue' {
  interface Vue {
    preprocessTask: () => void;
    doThisInput: () => void;
  }
}

export default Vue.extend({
  data:() =>({
    pageReady: false,
    loading: false,
    dialog: false,
    selectedTask: {} as lib.Task,
    rules: {
      length: (len: number) => (v: string) => (v || '').length <= len || `最大20文字までです`,
      required: (v: string) => !!v || '必ずご記入ください',
      requiredSelect: (v: lib.User[]) => v.length != 0 || '1名は選択してください',
    },
  }),
  computed: {
    ...mapGetters('task', [
      'statuses',
      'types',
      'priorities',
      'allTasks',
      'task'
    ]),
    ...mapGetters([
      'user',
      'project',
      'projectAuthority',
    ]),
    isReadyObj,
    isEmptyObj,
    isEmptyArr,
    checkStatus,
    changeToISOFormat,
    changeToTimeStampFormat,
    changeToDateISOFormat,
    min() {
      const today = new Date();
      const numberOfDaysToAdd = 1;
      today.setDate(today.getDate() + numberOfDaysToAdd);
      return today.toISOString()
    },
    taskForm() {
      const additionalInfo = {
        id: Number(this.$route.params.key),
        assigner_id: this.user.id,
        project_id: Number(this.$route.params.id),
      }
      const cleansedData = {
        estimated_time: Number(this.selectedTask.estimated_time),
        start_time: this.selectedTask.start_time ? new Date(this.selectedTask.start_time) : null,
        deadline: this.selectedTask.deadline ? new Date(this.selectedTask.deadline) : null,
      }
      const status = this.statuses.find((status: {id: number}) => status.id === this.selectedTask.status_id);
      const type = this.types.find((type: {id: number}) => type.id === this.selectedTask.type_id);
      const priority = this.priorities.find((priority: {id: number}) => priority.id === this.selectedTask.priority_id);
      const field = this.project.fields.find((field: lib.Field) => field.id === this.selectedTask.field_id) || {};
      const milestone = this.project.milestones.find((milestone: lib.Milestone) => milestone.id === this.selectedTask.milestone_id) || {};
      const actual_time = 0
      const created_at = new Date(this.selectedTask.created_at);
      const requiredDataforDisplay = {
        status,
        type,
        field,
        milestone,
        priority,
        actual_time,
        comments: [],
        created_at,
      }
      const newTask = {...this.selectedTask, ...additionalInfo, ...cleansedData, ...requiredDataforDisplay}
      return newTask;
    },
    dialogForm() {
    const newAssignee = this.project.authority_users.find((user: lib.ProjectAuthority) => user.user_id === this.selectedTask.assignee_id);
    const newStatus = this.statuses.find((status: lib.Status) => status.id === this.selectedTask.status_id);
    const newType = this.types.find((type: lib.Type) => type.id === this.selectedTask.type_id);
    const newPriority = this.priorities.find((priority: lib.Priority) => priority.id === this.selectedTask.priority_id);
    const newMilestone = this.project.milestones.find((milestone: lib.Milestone) => milestone.id === this.selectedTask.milestone_id);
    const newField = this.project.fields.find((field: lib.Field) => field.id === this.selectedTask.field_id);
      return [
        {title: '課題のタイトル', newData: this.selectedTask.title, oldData: this.task.title },
        {title: '担当者', newData: newAssignee.user.name, oldData: this.task.assignee.name },
        {title: '状況', newData: newStatus.name, oldData: this.task.status },
        {title: 'タスクの種類', newData: newType?.name, oldData: this.task.type },
        {title: '優先順位', newData: newPriority.name, oldData: this.task.priority },
        {title: 'マイルストーン', newData: newMilestone?.name, oldData: this.task.milestone },
        {title: '分野', newData: newField?.name, oldData: this.task.field },
        {title: '期日', newData: this.selectedTask.deadline, oldData: this.changeToDateISOFormat(this.task.deadline) },
        {title: '推定時間', newData: this.selectedTask.estimated_time, oldData: this.task.estimated_time },
        {title: 'タスクの詳細', newData: this.selectedTask.detail, oldData: this.task.detail },
      ]
    }
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.task)) return;
      clearInterval(timer);
      this.preprocessTask();
    }, 100)
  },
  methods: {
    preprocessTask: function(): void {
      this.selectedTask = JSON.parse(JSON.stringify(this.task));
      this.selectedTask.deadline = this.changeToDateISOFormat(this.selectedTask.deadline);
      this.pageReady = true;
    },
    changeTimeFormat(time: string) {
      const dateObj = new Date(time);
      const year = dateObj.getFullYear();
      const month = dateObj.getMonth() + 1;
      const date = dateObj.getDate();
      const dateStr = year + '-' + month + '-' + date;
      return dateStr;
    },
    async onClickSubmit() {
      this.loading = true;
      // const validation = this.validateForm(this.form);
      // if(validation) {
      //   return;
      // }
      let response
      try {
        response = await this.$store.dispatch('task/updateTask', this.taskForm);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'project-id', params: {id: this.$route.params.id}});
        },
        () => {
          this.loading = false;
        }
        )
      }
    },
    validateForm() {
      if(!this.selectedTask.title) {
        return '課題のタイトルが空白です'
      }
      // if(!form.field_id) {
      //   return '分野を選択してください'
      // }
      // if(!form.milestone_id) {
      //   return 'マイルストーンを選択してください'
      // }
      if(!this.selectedTask.deadline) {
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
