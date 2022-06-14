<template>
  <div>
    <v-card
     class="mx-auto my-4 px-8 py-4"
     max-width="700px"
     v-if="isReadyObj(user)"
    >
      <div class="pb-2">
        <v-btn
         icon
         @click="$router.push({name: 'profile'})"
         v-if="user.name"
        >
          <v-icon>mdi-arrow-left</v-icon>
        </v-btn>
        <v-spacer></v-spacer>
      </div>
      <v-form
       class="pa-4 pt-6"
       ref="form"
       v-model="formReady"
      >
        <v-text-field
          v-model="profile.name"
          type="text"
          :rules="[rules.required]"
          filled
          color="#295caa"
          label="名前"
        ></v-text-field>
        <v-text-field
          v-model="profile.age"
          type="text"
          :rules="[rules.required, isNumber(profile.age), rules.maxAge(3)]"
          filled
          color="#295caa"
          label="年齢"
        ></v-text-field>
        <v-select
          v-model="profile.sex"
          :items="sexes"
          label="性別"
          persistent-hint
          single-line
          outlined
          :rules="[rules.required]"
        ></v-select>
        <v-text-field
          v-model="profile.address"
          type="text"
          filled
          color="#295caa"
          label="住所"
        ></v-text-field>
        <v-text-field
          v-model="profile.description"
          type="text"
          filled
          color="#295caa"
          label="自己紹介"
        ></v-text-field>
        <v-text-field
          class="password-height"
          v-model="profile.password"
          type="password"
          :rules="[rules.required, rules.length(8)]"
          filled
          color="#295caa"
          label="パスワード"
          v-if="!user.name"
        ></v-text-field>
        <v-cropper
         v-model="profile.image_data"
         ratio="1:1"
         :pixel="900"
         :width="450"
         :currentImage="currentImage"
        ></v-cropper>
        <div class="px-4 py-2 red--text accent-3 text-center" style="height: 80px">{{ this.error }}</div>
        <v-card-actions>
          <v-spacer />
          <v-btn text @click="$router.push(to)" v-if="user.name">
            戻る
          </v-btn>
          <v-btn
           class="white--text"
           color="#295caa"
           depressed
           :loading="loading"
           :disabled="!formReady || loading"
           @click="onClickSend"
          >
            {{ user.name ? '更新' : '登録'}}
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import VCropper from '~/components/VCropper.vue'
import {checkStatus, isNumber, isEmptyObj, isReadyObj} from '~/modules/utils'
declare module 'vue/types/vue' {
  interface Vue {
    checkStatus: (status: number, func: Function, error: Function) => string;
    form: () => {},
    handleSuccess: () => void;
  }
}
export default Vue.extend({
  components: { VCropper },
  data: () => ({
    error: '',
    formReady: false,
    loading: false,
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
    rules: {
      length: (len: number) => (v: string) => (v || '').length >= len || `パスワードは${len}文字以上です`,
      maxAge: (len: number) => (v: string) => (String(v) || '').length < len || `間違っている可能性があります`,
      required: (v: string) => !!v || '記入必須です',
    },
  }),
  computed: {
    ...mapGetters([
      'user',
    ]),
    checkStatus,
    isEmptyObj,
    isReadyObj,
    isNumber,
    currentImage() {
      return this.$config.mediaURL + '/users/' + this.profile.image;
    },
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
      this.loading = true;
      let response;
      try {
        response = await this.$store.dispatch('registerUser', this.form());
      } catch (error) {
        response = error;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.error = this.checkStatus(response.status, (() => {return this.handleSuccess()}), () => this.loading = false);
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
<style lang="scss" scoped>
  .password-hright {
    min-height: 96px;
  }
</style>