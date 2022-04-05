<template>
  <div v-if="isAuthorized">
    <project-form v-model="selectedProject" @submit="onClickSend" :loading="loading">
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
  </div>
</template>

<script lang="ts">

import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isEmptyObj, isEmptyArr, resizeFile, checkStatus, isReadyObj } from '~/modules/utils'
import * as lib from '~/modules/store'
declare module 'vue/types/vue' {
  interface Vue {
    preprocessProjectAuthority: () => void;
    projectForm: () => lib.ProjectAuthority
  }
}
export default Vue.extend({
  data:() => ({
    isAuthorized: false,
    loading: false,
    selectedProject: {
      project: {
        organization_id: '',
        name: '',
        description: '',
        image: '',
        image_data: '',
        milestones: [{id: 0, name: ''}],
        fields: [{id: 0, name: ''}],
        users: [] as lib.User[]
      },
    },
    rules: {
      length: (len: number) => (v: string) => (v || '').length <= len || `最大20文字までです`,
      required: (v: string) => !!v || '必ずご記入ください',
      requiredSelect: (v: lib.User[]) => v.length != 0 || '1名は選択してください',
    },
    error: '',
  }),
  computed: {
    ...mapGetters([
      'user',
      'organization',
    ]),
    isEmptyObj,
    isEmptyArr,
    isReadyObj,
    checkStatus,
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.user)) return;
      clearInterval(timer)
      const authority = this.user.authority;
      const authorityType = '管理者';
      if(authority != authorityType) return this.$router.back();
      this.isAuthorized = true;
    }, 100);
  },
  methods: {
    async onClickSend() {
      this.loading = true;
      let response;
      console.log(this.projectForm())
      try {
        response = await this.$store.dispatch('project/createProject', this.projectForm());
      } catch (error: any) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'project'})
        },
        () => {
          this.loading = false;
        }
        )
      }
    },
    projectForm() {
      const project = {} as lib.Project
      const isFirstMilestone = !!this.selectedProject.project.milestones[0].name;
      const isFirstField = !!this.selectedProject.project.fields[0].name;
      project.milestones = isFirstMilestone ? this.selectedProject.project.milestones : [];
      project.fields = isFirstField ? this.selectedProject.project.fields : [];
      project.organization_id = this.organization.id;
      project.users = [this.user];
      const newProject = {...this.selectedProject.project, ...project}
      return newProject;
    },
  }
})
</script>