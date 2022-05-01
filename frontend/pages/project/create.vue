<template>
  <div v-if="isAuthorized">
    <project-form v-model="newProject" @submit="onClickSend" :loading="loading">
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
import { isEmptyObj, isEmptyArr, checkStatus, isReadyObj } from '~/modules/utils'
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
    newProject: {
      organization_id: '',
      name: '',
      description: '',
      image: '',
      image_data: '',
      milestones: [{id: 0, name: ''}],
      fields: [{id: 0, name: ''}],
      users: [] as lib.User[]
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
      'projectAuthority',
    ]),
    isEmptyObj,
    isEmptyArr,
    isReadyObj,
    checkStatus,
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.organization)) return;
      clearInterval(timer)
      const authority = this.organization.auth_id;
      if(authority != 1) return this.$router.back();
      this.isAuthorized = true;
    }, 100);
  },
  methods: {
    async onClickSend() {
      this.loading = true;
      let response;
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
      const project = {...this.newProject} as lib.Project
      const isFirstMilestone = !!this.newProject.milestones[0].name;
      const isFirstField = !!this.newProject.fields[0].name;
      project.milestones = isFirstMilestone ? this.newProject.milestones : [];
      project.fields = isFirstField ? this.newProject.fields : [];
      project.organization_id = this.organization.organization_id;
      project.authority_users = [
        {user_id: this.user.id, auth_id: 1, active: true} as any
      ];
      return project;
    },
  }
})
</script>