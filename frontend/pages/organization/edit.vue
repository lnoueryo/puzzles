<template>
  <div v-if="isAuthorized">
    <organization-form v-model="organizationAuthority.organization" @submit="dialog = true" :loading="loading">
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
    </organization-form>
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
    organizationAuthority: {} as lib.OrganizationAuthority,
    formReady: false,
    loading: false,
    dialog: false,
    error: '',
  }),
  computed: {
    ...mapGetters([
      'project',
      'organization',
      'user',
    ]),
    isEmptyObj,
    isEmptyArr,
    isReadyObj,
    checkStatus,
    dialogForm() {
      return [
        {title: '組織名', newData: this.organizationAuthority.organization.name, oldData: this.organization.organization.name},
        {title: 'プロジェクトの概要', newData: this.organizationAuthority.organization.description, oldData: this.organization.organization.description},
        {title: '設立日', newData: this.organizationAuthority.organization.founded, oldData: this.organization.organization.founded},
        {title: '電話番号', newData: this.organizationAuthority.organization.number, oldData: this.organization.organization.number},
        {title: 'イメージの変更', newData: this.organizationAuthority.organization.image_data || this.organizationAuthority.organization.image, oldData: this.organization.organization.image, image: true},
      ];
    }
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.organization)) return;
      clearInterval(timer)
      const authority = this.organization.auth_id;
      if(authority != 1) return this.$router.back();
      this.organizationAuthority = JSON.parse(JSON.stringify(this.organization))
      this.isAuthorized = true;
    }, 100);
  },
  methods: {
    async onClickSubmit() {
      this.dialog = false;
      let response;
      try {
        response = await this.$store.dispatch('updateOrganization', this.organizationAuthority.organization);
      } catch (error: any) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
          this.$router.push({name: 'index', params: {id: this.$route.params.id}})
        },
        () => {
          this.loading = false;
        }
        )
      }
    },
  }
})
</script>