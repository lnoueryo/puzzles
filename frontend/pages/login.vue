<template>
  <form-card @send="onClickSend" style="max-width: 500px" v-if="pageReady">
    <v-form
      ref="form"
      v-model="formReady"
      class="pa-4 pt-6"
    >
      <v-text-field
        v-model="organization"
        :rules="[rules.required]"
        filled
        color="amber darken-3"
        label="組織 ID"
        type="text"
        @keyup.enter="onClickSend"
      ></v-text-field>
      <v-text-field
        v-model="email"
        :rules="[rules.required, rules.email]"
        filled
        color="amber darken-3"
        label="メールアドレス"
        type="email"
        @keyup.enter="onClickSend"
      ></v-text-field>
      <v-text-field
        v-model="password"
        :rules="[rules.length(20)]"
        filled
        color="amber darken-3"
        counter="6"
        label="パスワード"
        style="min-height: 96px"
        type="password"
        @keyup.enter="onClickSend"
      ></v-text-field>
    </v-form>
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
    formReady: false,
    isLoading: false,
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
        organization: this.organization,
        email: this.email,
        password: this.password,
      }
    }
  },
  async beforeCreate() {
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
        console.log(response.status)
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
      let response;
      try {
        response = await this.$store.dispatch('login', this.form);
      } catch (error) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.error = this.checkStatus(response.status, (() => {return this.handleSuccess()}), ((): string => {return '組織ID、メールアドレス、またはパスワードが違います。'}));
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