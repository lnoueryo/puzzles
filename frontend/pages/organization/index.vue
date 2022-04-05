<template>
  <form-card @send="onClickSend" style="max-width: 500px">
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
    isLoading: false,
    email: 'r.inoue@goga.co.jp',
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
      }
    }
  },
  created() {
    console.log(process.browser)
  },
  methods: {
    async onClickSend() {
      let response: any;
      try {
        response = await this.$store.dispatch('sendEmail', this.form) as lib.Response;
      } catch (error: any) {
        response = error as lib.Response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection');
        this.error = this.checkStatus(response.status, (() => {return this.handleSuccess()}), (() => {return response.data.message}));
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