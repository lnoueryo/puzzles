<template>
  <div v-if="pageReady">
    <v-row justify="center">
      <v-col
       cols="12"
       sm="10"
       md="10"
       lg="10"
      >
        <FormTask
         v-model="selectedTask"
         @submit="onClickCreate"
         :loading="loading"
        >
          <template v-slot:back>
            戻る
          </template>
          <template v-slot:submit>
            作成
          </template>
        </FormTask>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isReadyObj, isEmptyArr, isEmptyObj, checkStatus, changeToDateISOFormat, changeToTimeStampFormat } from '~/modules/utils'
import * as model from '~/modules/model'
import FormTask from '~/components/FormTask.vue'
declare module 'vue/types/vue' {
  interface Vue {
    changeToHalf: (string: string) => any;
    isNumber: (v: any) => boolean;
    validateForm: (v: model.Task) => string;
  }
}
export default Vue.extend({
  components: { FormTask },
  data: () => ({
    selectedTask: {} as any,
    loading: false,
    pageReady: false
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
    changeToDateISOFormat,
    changeToTimeStampFormat,
    checkStatus,
    isEmptyArr,
    isEmptyObj,
    isReadyObj,
  },

  /** 新しく作成するタスクの初期設定 */
  created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.user)) return;
      clearInterval(timer);
      const additionalInfo = {
        assignee: this.user,
        status: {id: 1, name: ''},
        type: {id: 1},
        priority: {id: 2},
        estimated_time: 0,
        deadline: this.changeToDateISOFormat('', 5)
      }
      console.log(additionalInfo)
      this.selectedTask = {...this.selectedTask, ...additionalInfo}
      this.pageReady = true
    }, 100)
  },
  methods: {
    async onClickCreate() {
      this.loading = true;
      /** バリデーション追加 */
      // const validation = this.validateForm(this.taskForm);
      // if(validation) {
      //   return;
      // }
      const request = {
        id: this.selectedTask.id,
        key: this.selectedTask.key,
        title: this.selectedTask.title,
        detail: this.selectedTask.detail,
        actual_time: this.selectedTask.actual_time,
        start_time: new Date(this.selectedTask.start_time),
        estimated_time: this.selectedTask.estimated_time,
        deadline: new Date(this.selectedTask.deadline),
        project_id: Number(this.$route.params.id),
        assignee_id: this.selectedTask.assignee.id,
        assigner_id: this.selectedTask.assignee.id,
        field_id: this.selectedTask.field?.id,
        milestone_id: this.selectedTask.milestone?.id,
        priority_id: this.selectedTask.priority.id,
        status_id: this.selectedTask.status.id,
        type_id: this.selectedTask.type.id,
        version_id: this.selectedTask.version?.id,
        created_at: new Date(this.selectedTask.created_at),
        updated_at: new Date(this.selectedTask.updated_at),
      }
      let response
      try {
        response = await this.$store.dispatch('task/createTask', request);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/error/bad-connection')
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
