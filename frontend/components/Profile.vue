<template>
  <div style="width: 100%;max-width: 600px;margin: auto;" v-if="isReadyObj(selectedUser)">
    <v-row class="py-8" align="center" justify="center">
      <v-list class="px-4" subheader two-line width="600">
        <v-row class="py-8 mx-4" align="center" justify="center" style="position: relative">
          <div style="position: relative;">
            <v-avatar size="36px" class="mr-4" style="position: absolute;left: -50px">
              <v-img alt="Avatar" style="object-fit: cover;" :src="$config.mediaURL + '/users/' + selectedUser.user.image" v-if="selectedUser.user.image">
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
            <h2>{{ selectedUser.user.name }}</h2>
          </div>
          <v-btn icon absolute right color="#295caa" :to="{name: 'profile-edit'}" v-if="!disabled">
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
  props: ['selectedUser', 'disabled'],
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