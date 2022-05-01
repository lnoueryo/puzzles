<template>
    <div style="width: 100%;max-width: 600px;margin: auto;">
      <v-row class="py-8" align="center" justify="center" style="position: relative">
        <h2>{{ project.name }}</h2>
        <v-btn absolute right color="#295caa" v-if="projectAuthority.type.name == '管理者'" :to="{name: 'project-id-edit', params: {id: $route.params.id}}">
          <v-icon left>mdi-application-edit-outline</v-icon>
          編集
        </v-btn>
      </v-row>
      <v-row class="px-4" align="center" justify="center">
        <v-img :aspect-ratio="16/9" :src="$config.mediaURL + '/projects/' + project.image"></v-img>
      </v-row>
      <v-row class="pt-8" align="center" justify="center">
        <h2>概要</h2>
      </v-row>
      <v-row class="pb-8" align="center" justify="center">
        <div v-if="project.description">{{ project.description }}</div>
        <div v-else>なし</div>
      </v-row>
      <v-row class="py-8" align="center" justify="center">
        <div class="d-flex justify-center align-center" style="position: relative;width: 100%">
          <h2>ユーザー</h2>
          <v-btn absolute right color="#295caa" v-if="projectAuthority.type.name == '管理者'" @click="addUserDialog = true">
            <v-icon left>mdi-account-multiple-plus</v-icon>
            メンバー追加
          </v-btn>
        </div>
        <v-list subheader two-line>
          <v-subheader inset>ユーザー</v-subheader>

          <v-list-item
            v-for="authUser in project.authority_users"
            :key="authUser.id"
          >
              <v-list-item-avatar>
                <v-img :src="$config.mediaURL + '/users/' + authUser.user.image"></v-img>
              </v-list-item-avatar>

            <v-list-item-content class="mr-4" style="width: 400px">
              <v-list-item-title v-text="authUser.user.name"></v-list-item-title>
              <v-list-item-subtitle v-text="authUser.user.email"></v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-content class="text-center">
              <v-list-item-subtitle v-text="authUser.type.name"></v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-action v-if="projectAuthority.type.name == '管理者'">
              <v-menu left bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn icon v-bind="attrs" v-on="on">
                    <v-icon>mdi-dots-vertical</v-icon>
                  </v-btn>
                </template>

                <v-list>
                  <v-list-item v-for="(option, i) in userOptions" :key="i" @click="option.event()">
                    <v-list-item-title>{{ option.title }}</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>
            </v-list-item-action>
          </v-list-item>
        </v-list>
      </v-row>
      <v-row justify="center">
        <v-dialog v-model="addUserDialog" scrollable max-width="300px">
          <v-card>
            <v-card-title>Select Country</v-card-title>
            <v-divider></v-divider>
            <v-card-text style="height: 300px;">
              <v-radio-group v-model="selectUser" column >
                <v-radio :value="auth" :label="auth.user.name" v-for="(auth, i) in unregisteredUsers" :key="i">
                </v-radio>
              </v-radio-group>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
              <v-btn color="blue darken-1" text @click="addUserDialog = false;selectUser = ''">
                Close
              </v-btn>
              <v-btn color="blue darken-1" text @click="addUserDialog = false;selectUser = ''">
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import * as lib from '~/modules/store'
declare module 'vue/types/vue' {
  interface Vue {
    changeAuthority: () => void
  }
}
export default Vue.extend({
  data: () => ({
    addUserDialog: false,
    selectUser: '',
  }),
  computed: {
    ...mapGetters([
      'user',
      'organization',
      'project',
      'projectAuthority',
    ]),
    userOptions() {
      return [
        {title: '権限変更', event: this.changeAuthority},
        {title: '削除', event: this.changeAuthority},
      ]
    },
    unregisteredUsers() {
      return this.organization.organization.users.filter((oUser: lib.OrganizationAuthority) => {
      return !this.project.authority_users.some((user: lib.ProjectAuthority) => user.user_id == oUser.user_id)
    });
    }
  },
  // 0.30000000447034836 milliseconds.
  // 0.40000000447034836 milliseconds.
  // 0.40000000447034836 milliseconds.
  // 0.60000000447034836 milliseconds.

  // 0.4 milliseconds.
  // 0.5 milliseconds.
  // 0.6 milliseconds.
  // 0.7 milliseconds.
  // 0.8 milliseconds.
  created() {
    const itemStr = sessionStorage.getItem(location.host + window.$nuxt.$route.params.id);
    if(!itemStr) return;
    const item = JSON.parse(itemStr);
  },
  methods: {
    changeAuthority() {
      console.log('AAA')
    }
  }
})
</script>