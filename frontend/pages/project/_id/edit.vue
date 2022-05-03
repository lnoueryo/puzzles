<template>
  <div v-if="isAuthorized">
    <project-form v-model="selectedProject" @submit="dialog = true" :loading="loading">
      <template slot="back">
        <div>
          戻る
        </div>
      </template>
      <template slot="submit">
        <div>
          更新
        </div>
      </template>
    </project-form>
    <update-dialog v-model="dialog" :form="dialogForm" @submit="onClickSubmit" @loading="loading = $event">
      更新の確認
    </update-dialog>
  </div>
</template>

<script lang="ts">

import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isEmptyObj, isEmptyArr, checkStatus, isReadyObj } from '~/modules/utils'
import * as lib from '~/modules/store'
declare module 'vue/types/vue' {
  interface Vue {
    preprocessProjectAuthority: () => void;
  }
}
interface ProjectAuthority extends lib.ProjectAuthority {
  disabled: boolean
  project: lib.Project
}
export default Vue.extend({
  data:() => ({
    isAuthorized: false,
    formReady: false,
    loading: false,
    dialog: false,
    selectedProject: {} as lib.Project,
    projectUsers: [] as lib.ProjectAuthority[],
    rules: {
      length: (len: number) => (v: string) => (v || '').length <= len || `最大20文字までです`,
      required: (v: string) => !!v || '必ずご記入ください',
      requiredSelect: (v: lib.User[]) => v.length != 0 || '1名は選択してください',
    },
    error: '',
  }),
  computed: {
    ...mapGetters([
      'project',
      'projectAuthority',
      'user',
    ]),
    isEmptyObj,
    isEmptyArr,
    isReadyObj,
    checkStatus,
    organizationUsers() {
      return this.$store.getters['organization'].users;
    },
    dialogForm() {
      // imageの比較もする
      let fields = this.selectedProject.fields.map((field, i) => {
        return {title: '分野' + (i + 1), newData: field.name, oldData: ''};
      })
      const newFieldLength = fields.length - 1;
      this.project.fields.forEach((field: lib.Field, i: number) => {
        if(newFieldLength >= i) {
          fields[i] = {...fields[i], ...{oldData: field.name}}
        } else {
          fields.push({title: '分野' + (i + 1), newData: '', oldData: field.name})
        }
      });
      let milestones = this.selectedProject.milestones.map((milestone, i) => {
        return {title: 'マイルストーン' + (i + 1), newData: milestone.name, oldData: ''};
      })
      const newMilestoneLength = milestones.length - 1;
      this.project.milestones.forEach((milestone: lib.Milestone, i: number) => {
        if(newMilestoneLength >= i) {
          milestones[i] = {...milestones[i], ...{oldData: milestone.name}}
        } else {
          milestones.push({title: 'マイルストーン' + (i + 1), newData: '', oldData: milestone.name})
        }
      });
      // const filteredNewProject = this.selectedProject.authority_users.filter((authority_user) => authority_user.auth_id == 1)
      // let administers = filteredNewProject.map((authority_user, i) => {
      //   return {title: 'プロジェクト管理者' + (i + 1), newData: authority_user.user.name, oldData: ''};
      // })
      // const authorityType = 1;
      // const newAdministerLength = administers.length - 1;
      // this.project.authority_users.forEach((authority_user: lib.ProjectAuthority, i: number) => {
      //   if(authority_user.auth_id != authorityType) return
      //   if(newAdministerLength >= i) {
      //     administers[i] = {...administers[i], ...{oldData: authority_user.user.name}}
      //   } else {
      //     administers.push({title: 'プロジェクト管理者' + (i + 1), newData: '', oldData: authority_user.user.name})
      //   }
      // });
      return [
        {title: 'プロジェクト名', newData: this.selectedProject.name, oldData: this.project.name},
        ...fields,
        ...milestones,
        // ...administers,
        {title: 'イメージの変更', newData: this.selectedProject.image_data || this.selectedProject.image, oldData: this.selectedProject.image, image: true},
        {title: 'プロジェクトの概要', newData: this.selectedProject.description, oldData: this.project.description},
      ];
    }
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.project)) return;
      clearInterval(timer)
      const authority = this.projectAuthority.auth_id;
      if(authority != 1) return this.$router.back();
      this.preprocessProjectAuthority()
    }, 100);
  },
  methods: {
    preprocessProjectAuthority() {
      this.selectedProject = JSON.parse(JSON.stringify(this.project));
      if(this.isEmptyArr(this.project.milestones)) this.selectedProject.milestones.push({id: 0, name: ''})
      if(this.isEmptyArr(this.project.fields)) this.selectedProject.fields.push({id: 0, name: ''})
      this.isAuthorized = true;
    },
    async onClickSubmit() {
      this.dialog = false;
      let response;
      try {
        response = await this.$store.dispatch('project/updateProject', this.projectForm());
      } catch (error: any) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'project-id', params: {id: this.$route.params.id}})
        },
        () => {
          this.loading = false;
        }
        )
      }
    },
    projectForm() {
      const project = {} as lib.Project
      const isFirstField = !!this.selectedProject.fields[0].name;
      const isFirstMilestone = !!this.selectedProject.milestones[0].name;
      project.fields = isFirstField ? this.selectedProject.fields : [];
      project.milestones = isFirstMilestone ? this.selectedProject.milestones : [];
      let fieldDelete = this.project.fields.length > this.selectedProject.fields.length;
      if(!fieldDelete) fieldDelete = !!this.project.fields[0]?.name && !isFirstField;
      let milestoneDelete = this.project.milestones.length > this.selectedProject.milestones.length;
      if(!milestoneDelete) milestoneDelete = !!this.project.milestones[0]?.name && !isFirstMilestone;
      // project.authority_users = this.selectedProject.authority_users;
      // console.log(this.selectedProject)
      const newProject = {...this.selectedProject, ...project}
      const request = {
        project: newProject,
        field_delete: fieldDelete,
        milestone_delete: milestoneDelete,
      }
      return request;
    },
    onDeleteField(index: number) {
      if(this.selectedProject.fields.length == 1) return;
      this.selectedProject.fields.splice(index, 1)
    },
    onDeleteMilestone(index: number) {
      if(this.selectedProject.milestones.length == 1) return;
      this.selectedProject.milestones.splice(index, 1)
    },
  }
})
</script>