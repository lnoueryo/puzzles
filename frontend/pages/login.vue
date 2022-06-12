<template>
  <form-card @send="onClickSend" style="max-width: 500px" v-if="pageReady" :loading="loading" :formReady="formReady">
    <template v-slot:main>
      <v-form ref="form" v-model="formReady" class="pa-4 pt-6">
        <v-text-field
          id="organizaiton"
          v-model="organization"
          :rules="[rules.required]"
          filled
          color="#295caa"
          label="組織 ID"
          type="text"
          @keyup.enter="onClickSend"
        ></v-text-field>
        <v-text-field
          id="email"
          v-model="email"
          :rules="[rules.required, rules.email]"
          filled
          color="#295caa"
          label="メールアドレス"
          type="email"
          @keyup.enter="onClickSend"
        ></v-text-field>
        <v-text-field
          id="password"
          v-model="password"
          filled
          color="#295caa"
          label="パスワード"
          style="min-height: 96px"
          type="password"
          @keyup.enter="onClickSend"
        ></v-text-field>
        <div style="color: red">
          {{ error }}
        </div>
      </v-form>
    </template>
    <template v-slot:button>
      ログイン
    </template>
  </form-card>
</template>

<script lang="ts">
import Vue from 'vue'
import {checkStatus} from '~/modules/utils'
declare module 'vue/types/vue' {
  interface Vue {
    checkStatus: (status: number, func: Function, error: Function) => string;
    handleSuccess: () => void;
    pageReady: boolean
  }
}
export default Vue.extend({
  name: 'login',
  layout: 'login',
  data: () => ({
    pageReady: false,
    error: '',
    organization: '',
    email: undefined,
    password: undefined,
    formReady: true,
    loading: false,
    rules: {
      email: (v: string) => !!(v || '').match(/@/) || 'メールアドレスの形式ではありません',
      required: (v: string) => !!v || '必須項目です',
    },
  }),
  computed: {
    checkStatus,
    form() {
      return {
        organization: this.organization,
        email: this.email,
        password: this.password,
      }
    }
  },
  async beforeCreate() {
    console.log('Hello')
    this.$store.dispatch('resetAll');
      let response;
      try {
        response = await this.$store.dispatch('session');
      } catch (error) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        if(response.status == 304) {
          this.pageReady= true;
          return;
        }
        this.checkStatus(response.status, (() => {return this.$router.push('/')}), (() => {}));
      }
  },
  created() {
    const storageJson = JSON.parse(localStorage.getItem(window.location.host) as string);
    if(storageJson) {
      this.organization = storageJson.organization;
      this.email = storageJson.email;
    }
  },
  methods: {
    async onClickSend() {
      this.loading = true;
      let response;
      try {
        response = await this.$store.dispatch('login', this.form);
      } catch (error) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.error = this.checkStatus(response.status, (() => {return this.handleSuccess()}), ((): string => {
          this.loading = false;
          return '組織ID、メールアドレス、またはパスワードが違います。'
        }));
      }
    },
    handleSuccess() {
      const jsonString = JSON.stringify({organization: this.organization, email: this.email})
      localStorage.setItem(window.location.host, jsonString);
      this.$router.push({name: 'index'});
    }
  }
})
</script>