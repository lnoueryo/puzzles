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
    selectedProject: {} as lib.ProjectAuthority,
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
      let fields = this.selectedProject.project.fields.map((field, i) => {
        return {title: '分野' + (i + 1), newData: field.name, oldData: ''};
      })
      const newFieldLength = fields.length - 1;
      this.projectAuthority.project.fields.forEach((field: lib.Field, i: number) => {
        if(newFieldLength >= i) {
          fields[i] = {...fields[i], ...{oldData: field.name}}
        } else {
          fields.push({title: '分野' + (i + 1), newData: '', oldData: field.name})
        }
      });
      let milestones = this.selectedProject.project.milestones.map((milestone, i) => {
        return {title: 'マイルストーン' + (i + 1), newData: milestone.name, oldData: ''};
      })
      const newMilestoneLength = milestones.length - 1;
      this.projectAuthority.project.milestones.forEach((milestone: lib.Milestone, i: number) => {
        if(newMilestoneLength >= i) {
          milestones[i] = {...milestones[i], ...{oldData: milestone.name}}
        } else {
          milestones.push({title: 'マイルストーン' + (i + 1), newData: '', oldData: milestone.name})
        }
      });
      let administers = this.selectedProject.project.authority_users.map((authority_user, i) => {
        return {title: 'プロジェクト管理者' + (i + 1), newData: authority_user.user.name, oldData: ''};
      })
      const authorityType = '管理者';
      const newAdministerLength = administers.length - 1;
      this.projectAuthority.project_users.forEach((authority_user: lib.ProjectAuthority, i: number) => {
        if(authority_user.type.name != authorityType) return
        if(newAdministerLength >= i) {
          administers[i] = {...administers[i], ...{oldData: authority_user.user.name}}
        } else {
          administers.push({title: 'プロジェクト管理者' + (i + 1), newData: '', oldData: authority_user.user.name})
        }
      });
      return [
        {title: 'プロジェクト名', newData: this.selectedProject.project.name, oldData: this.projectAuthority.project.name},
        ...fields,
        ...milestones,
        ...administers,
        {title: 'プロジェクトの概要', newData: this.selectedProject.project.description, oldData: this.projectAuthority.project.description},
      ];
    }
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.projectAuthority)) return;
      clearInterval(timer)
      const authority = this.user.authority;
      const authorityType = '管理者';
      if(authority != authorityType) return this.$router.back();
      this.preprocessProjectAuthority()
    }, 100);
  },
  methods: {
    preprocessProjectAuthority() {
      const authorityType = '管理者';
      this.selectedProject = JSON.parse(JSON.stringify(this.projectAuthority));
      console.log(this.selectedProject.project_users.filter((user) => user.type.name == authorityType))
      this.selectedProject.project.authority_users = this.selectedProject.project_users.filter((user) => user.type.name == authorityType);
      if(this.isEmptyArr(this.projectAuthority.project.milestones)) this.selectedProject.project.milestones.push({id: 0, name: ''})
      if(this.isEmptyArr(this.projectAuthority.project.fields)) this.selectedProject.project.fields.push({id: 0, name: ''})
      // console.log(this.selectedProject.project.fields)
      this.isAuthorized = true;
    },
    async onClickSubmit() {
      console.log(this.projectForm())
      let response;
      try {
        response = await this.$store.dispatch('project/updateProject', this.projectForm());
      } catch (error: any) {
        response = error;
      } finally {
        // const blob = new Blob([JSON.stringify(response.data, null, '  ')], {type: 'application\/json'});
        // const url = URL.createObjectURL(blob);
        // location.href = url;
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'project-id-task', params: {id: this.$route.params.id}})
        },
        () => {
          this.loading = false;
        }
        )
      }
    },
    projectForm() {
      const project_authority = {} as lib.ProjectAuthority;
      const project = {} as lib.Project
      const isFirstField = !!this.selectedProject.project.fields[0].name;
      const isFirstMilestone = !!this.selectedProject.project.milestones[0].name;
      project.fields = isFirstField ? this.selectedProject.project.fields : [];
      project.milestones = isFirstMilestone ? this.selectedProject.project.milestones : [];
      let fieldDelete = this.projectAuthority.project.fields.length > this.selectedProject.project.fields.length;
      if(!fieldDelete) fieldDelete = !!this.projectAuthority.project.fields[0]?.name && !isFirstField;
      let milestoneDelete = this.projectAuthority.project.milestones.length > this.selectedProject.project.milestones.length;
      if(!milestoneDelete) milestoneDelete = !!this.projectAuthority.project.milestones[0]?.name && !isFirstMilestone;
      const adminUserNum = this.selectedProject.project.authority_users.map((user) => user.user_id)
      project_authority.project_users = this.selectedProject.project_users.map((user: lib.ProjectAuthority) => {
        const newUser = {} as lib.ProjectAuthority
        newUser.auth_id = 2;
        if(adminUserNum.includes(newUser.user_id)) {
          newUser.auth_id = 1;
          newUser.active = true;
        }
        return {...user, ...newUser};
      });
      const newProject = {...this.selectedProject.project, ...project}
      project_authority.project = newProject
      const request = {
        project_authority: {...this.selectedProject, ...project_authority},
        field_delete: fieldDelete,
        milestone_delete: milestoneDelete,
      }
      return request;
    },
    onDeleteField(index: number) {
      if(this.selectedProject.project.fields.length == 1) return;
      this.selectedProject.project.fields.splice(index, 1)
    },
    onDeleteMilestone(index: number) {
      if(this.selectedProject.project.milestones.length == 1) return;
      this.selectedProject.project.milestones.splice(index, 1)
    },
  }
})
</script>