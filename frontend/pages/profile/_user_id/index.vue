<template>
  <div>
    <card-profile :selectedUser="selectedUser" :editable="editable" />
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
        'selectedUser',
        'organizationAuthority',
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
    },
    editable() {
      if(this.user?.id == this.selectedUser?.id) return this.$router.push({name: 'profile'})
      return this.user?.id == this.selectedUser?.id;
    }
  },
})
</script>