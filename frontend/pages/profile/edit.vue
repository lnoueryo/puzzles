<template>
  <form-card @send="onClickSend" style="max-width: 500px">
    <v-form
      ref="form"
      v-model="formReady"
      class="pa-4 pt-6"
    >
      <v-text-field
        v-model="profile.name"
        :rules="[rules.required]"
        filled
        color="amber darken-3"
        label="名前"
        type="text"
      ></v-text-field>
      <v-text-field
        v-model="profile.age"
        :rules="[rules.required]"
        filled
        color="amber darken-3"
        label="年齢"
        type="text"
      ></v-text-field>
      <v-select
        v-model="profile.sex"
        :items="sexes"
        label="性別"
        persistent-hint
        single-line
      ></v-select>
      <v-text-field
        v-model="profile.address"
        :rules="[rules.required]"
        filled
        color="amber darken-3"
        label="住所"
        type="text"
      ></v-text-field>
      <v-text-field
        v-model="profile.discription"
        :rules="[rules.required]"
        filled
        color="amber darken-3"
        label="自己紹介"
        type="text"
      ></v-text-field>
      <v-text-field
        v-model="profile.password"
        :rules="[rules.length(20)]"
        filled
        color="amber darken-3"
        counter="6"
        label="パスワード"
        style="min-height: 96px"
        type="password"
        v-if="!user.name"
      ></v-text-field>
      <cropper v-model="profile.image_data" :width="450" :currentImage="$config.mediaURL + '/users/' + profile.image"></cropper>
    </v-form>
  </form-card>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import {checkStatus, resizeFile, isEmptyObj} from '~/modules/utils'
declare module 'vue/types/vue' {
  interface Vue {
    checkStatus: (status: number, func: Function, error: Function) => string;
    handleSuccess: () => void;
  }
}
export default Vue.extend({
  name: 'login',
  layout: 'login',
  data: () => ({
    error: '',
    profile: {
      name: '',
      age: '' as number | string,
      sex: '',
      address: '',
      description: '',
      image: '',
      image_data: '',
      password: '',
    },
    formReady: false,
    isLoading: false,
    rules: {
      email: (v: string) => !!(v || '').match(/@/) || 'Please enter a valid email',
      length: (len: number) => (v: string) => (v || '').length >= len || `Invalid character length, required ${len}`,
      required: (v: string) => !!v || 'This field is required',
    },
  }),
  computed: {
    ...mapGetters([
      'user',
      'organization',
    ]),
    checkStatus,
    isEmptyObj,
    sexes() {
      return [
        '男',
        '女',
        'その他',
      ]
    }
  },
  created() {
    let timer = setInterval(() => {
      if(this.isEmptyObj(this.user)) return;
      console.log(this.$store.getters['mediaUser'])
      clearInterval(timer);
      this.profile = {...this.profile, ...this.user};
      this.profile.age = !this.profile.age ? '' : this.profile.age
    })
  },
  methods: {
    async onClickSend() {
      let response;
      try {
        response = await this.$store.dispatch('registerUser', this.form());
      } catch (error) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.error = this.checkStatus(response.status, (() => {return this.handleSuccess()}));
      }
    },
    form() {
      this.profile.age = Number(this.profile.age);
      return this.profile;
    },
    handleSuccess() {
      this.$router.push({name: 'index'});
    },
  }
})
</script>