<template>
  <form-card @send="onClickSend" style="max-width: 500px" :loading="loading">
    <template v-slot:main>
      <v-form
        ref="form"
        v-model="formReady"
        class="pa-4 pt-6"
      >
        <v-text-field
          v-model="email"
          :rules="[rules.required, rules.email]"
          filled
          color="amber darken-3"
          label="メールアドレス"
          type="email"
          @keyup.enter="onclickSend"
        ></v-text-field>
          <v-select
            v-model="authority"
            :items="items"
            item-text="name"
            item-value="name"
            label="Select"
            persistent-hint
            return-object
            single-line
          ></v-select>
      </v-form>
      <div class="pa-4 red--text">{{ error }}</div>
    </template>
    <template v-slot:button>
      送信
    </template>
  </form-card>
</template>

<script lang="ts">
import Vue from 'vue'
import {checkStatus} from '~/modules/utils'
import * as lib from '~/modules/store'
export default Vue.extend({
  data: () => ({
    error: '',
    formReady: false,
    loading: false,
    email: '',
    authority: { id: 1, name: '管理者' },
    items: [
      { id: 1, name: '管理者' },
      { id: 2, name: '一般' },
      { id: 3, name: 'ゲスト' },
    ],
    rules: {
      email: (v: string) => !!(v || '').match(/@/) || 'Please enter a valid email',
      length: (len: number) => (v: string) => (v || '').length >= len || `Invalid character length, required ${len}`,
      required: (v: string) => !!v || 'This field is required',
    },
  }),
  computed: {
    checkStatus,
    form() {
      return {
        email: this.email,
        authority_id: this.authority.id,
        organization_id: this.$store.getters['organization']?.organization_id,
      }
    },
  },
  methods: {
    async onClickSend() {
      this.loading = true;
      let response: any;
      try {
        response = await this.$store.dispatch('sendEmail', this.form) as lib.Response;
      } catch (error: any) {
        response = error as lib.Response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection');
        this.error = this.checkStatus(response.status, (() => {return this.handleSuccess()}), (() => {
          this.loading = false;
          return response.data.message
        }));
      }
    },
    handleSuccess() {
      const text = 'メールの送信が完了しました。'
      this.$store.dispatch('showSnackbar', text);
      this.$router.push('/');
    }
  }
})
</script>