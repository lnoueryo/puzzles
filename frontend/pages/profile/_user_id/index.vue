<template>
  <div>
    <profile :selectedUser="selectedUser" disabled="true" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isReadyObj, changeToTimeStampFormat } from '~/modules/utils'
export default Vue.extend({
  computed: {
    ...mapGetters([
    'user',
    'organization',
    'selectedUser',
    ]),
    isReadyObj,
    changeToTimeStampFormat,
    userEntities() {
      const userEntities = [
        {key: 'description', title: '自己紹介'},
        {key: 'age', title: '年齢'},
        {key: 'sex', title: '性別'},
        {key: 'email', title: 'メールアドレス'},
        {key: 'address', title: '住所'},
      ]
      return userEntities
    }
  }
})
</script>