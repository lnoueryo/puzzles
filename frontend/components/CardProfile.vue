<template>
  <div class="card-frame" v-if="isReadyObj(selectedUser)">
    <v-row class="py-8" align="center" justify="center">
      <v-list class="px-4" subheader two-line width="600">
        <v-row class="py-8 mx-4 relative" align="center" justify="center">
          <div class="relative">
            <v-avatar size="36px" class="mr-4 avatar-position">
              <v-img class="object-cover" alt="Avatar" :src="$config.mediaURL + '/users/' + selectedUser.user.image" v-if="selectedUser.user.image">
                <template v-slot:placeholder>
                  <v-row class="fill-height ma-0" align="center" justify="center">
                    <v-progress-circular indeterminate color="grey lighten-5" />
                  </v-row>
                </template>
              </v-img>
              <v-icon size="44px" dark v-else>
              mdi-account-circle
              </v-icon>
            </v-avatar>
            <h2 id="user-name">{{ selectedUser.user.name }}</h2>
          </div>
          <v-btn id="edit-profile" icon absolute right color="#295caa" :to="{name: 'profile-edit'}" v-if="editable">
            <v-icon>mdi-application-edit-outline</v-icon>
          </v-btn>
        </v-row>
        <v-list-item two-line v-for="(userEntity, i) in userEntities" :key="i">
          <v-list-item-content>
            <v-list-item-title>{{ userEntity.title }}</v-list-item-title>
            <v-list-item-subtitle>{{ selectedUser.user[userEntity.key] }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item two-line>
          <v-list-item-content>
            <v-list-item-title>組織参加日</v-list-item-title>
            <v-list-item-subtitle>{{ changeToTimeStampFormat(selectedUser.created_at) }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-row>
  </div>
</template>


<script lang="ts">
import Vue from 'vue'
import { isReadyObj, changeToTimeStampFormat } from '~/modules/utils'
export default Vue.extend({
  props: {
    selectedUser: Object,
    editable: Boolean
  },
  computed: {
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
<style lang="scss" scoped>
  .card-frame {
    width: 100%;
    max-width: 600px;
    margin: auto;
  }
  .avatar-position {
    position: absolute;
    left: -50px
  }
</style>