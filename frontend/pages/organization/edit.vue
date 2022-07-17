<template>
  <div v-if="isAuthorized">
    <form-organization v-model="organization" @submit="dialog = true" :loading="loading">
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
    </form-organization>
    <dialog-update v-model="dialog" :form="dialogForm" @submit="onClickSubmit" @loading="loading = $event">
      更新の確認
    </dialog-update>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isEmptyObj, isEmptyArr, checkStatus, isReadyObj, deepCopy } from '~/modules/utils'
import * as model from '~/modules/model'
export default Vue.extend({
  data:() => ({
    dialog: false,
    error: '',
    formReady: false,
    isAuthorized: false,
    loading: false,
    organization: {

    } as model.Organization,
  }),
  computed: {
    ...mapGetters([
      'project',
      'organizationAuthority',
      'user',
    ]),
    checkStatus,
    isEmptyObj,
    isEmptyArr,
    isReadyObj,
    deepCopy,
    dialogForm() {
      return [
        {title: '組織名', newData: this.organization.name, oldData: this.organizationAuthority.organization.name},
        {title: 'プロジェクトの概要', newData: this.organization.description, oldData: this.organizationAuthority.organization.description},
        {title: '設立日', newData: this.organization.founded, oldData: this.organizationAuthority.organization.founded},
        {title: '電話番号', newData: this.organization.number, oldData: this.organizationAuthority.organization.number},
        {title: 'イメージの変更', newData: this.organization.image_data || '/organizations/' + this.organization.image, oldData: '/organizations/' + this.organizationAuthority.organization.image, image: true},
      ];
    }
  },
  async created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.organizationAuthority)) return;
      clearInterval(timer)
      const authority = this.organizationAuthority.auth_id;
      if(authority != 1) return this.$router.back();
      this.organization = this.deepCopy(this.organizationAuthority.organization);
      this.isAuthorized = true;
    }, 100);
  },
  methods: {
    async onClickSubmit() {
      this.dialog = false;
      let response;
      try {
        response = await this.$store.dispatch('updateOrganization', this.organization);
      } catch (error: any) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/error/bad-connection')
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