<template>
    <div style="width: 100%;max-width: 600px;margin: auto;" v-if="isReadyObj(project)">
      <v-row class="py-8" align="center" justify="center" style="position: relative">
        <h2>{{ project.name }}</h2>
        <v-btn absolute right color="#295caa" v-if="projectAuthority.type.name == '管理者'" :to="{name: 'project-id-edit', params: {id: $route.params.id}}">
          <v-icon left>mdi-application-edit-outline</v-icon>
          編集
        </v-btn>
      </v-row>
      <v-row class="px-4" align="center" justify="center">
        <v-img :aspect-ratio="16/9" :src="projectImage" @error="projectImageError = true">
          <template v-slot:placeholder>
            <v-row class="fill-height ma-0" align="center" justify="center">
              <v-progress-circular indeterminate color="grey lighten-5" />
            </v-row>
          </template>
        </v-img>
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

          <v-list-item v-for="authUser in project.authority_users" :key="authUser.id">
              <v-list-item-avatar>
                <v-img :src="$config.mediaURL + '/users/' + authUser.user.image" v-if="authUser.user.image">
                  <template v-slot:placeholder>
                    <v-row class="fill-height ma-0" align="center" justify="center">
                      <v-progress-circular indeterminate color="grey lighten-5" />
                    </v-row>
                  </template>
                </v-img>
                <v-icon size="44px" dark v-else>
                mdi-account-circle
                </v-icon>
              </v-list-item-avatar>

            <v-list-item-content class="mr-4" style="width: 400px">
              <v-list-item-title v-text="authUser.user.name"></v-list-item-title>
              <v-list-item-subtitle v-text="authUser.user.email"></v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-content class="text-center">
              <v-list-item-subtitle v-text="authUser.type.name"></v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-action v-if="projectAuthority.type.name == '管理者'">
              <v-menu right bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn icon v-bind="attrs" v-on="on">
                    <v-icon>mdi-dots-vertical</v-icon>
                  </v-btn>
                </template>

                <v-list>
                  <v-list-item v-for="(option, i) in userOptions" :key="i" @click="option.event(authUser)" :disabled="authUser.user_id == user.id">
                    <v-list-item-title>{{ option.title }}</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>
            </v-list-item-action>
          </v-list-item>
        </v-list>
      </v-row>
      <v-row justify="center">
        <v-dialog v-model="addUserDialog" scrollable max-width="400px">
          <v-card>
            <v-card-title>ユーザー追加</v-card-title>
            <v-divider></v-divider>
            <v-card-text>
              <v-radio-group v-model="selectedUser" column >
                <v-radio :value="auth" v-for="(auth, i) in unregisteredUsers" :key="i" color="#295caa">
                  <template v-slot:label>
                    <div v-if="auth.user.name">{{auth.user.name}}</div>
                    <div v-else>{{auth.user.email}}　<strong class="red--text">招待中</strong></div>
                  </template>
                </v-radio>
              </v-radio-group>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="addUserDialog = false;selectedUser = ''">
                戻る
              </v-btn>
              <v-btn class="white--text" color="#295caa" @click="onAddUser" :disabled="isEmptyObj(selectedUser)">
                変更
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
      <v-row justify="center">
        <v-dialog v-model="authorityDialog" scrollable max-width="300px">
          <v-card>
            <v-card-title>権限変更</v-card-title>
            <v-divider></v-divider>
            <v-card-text>
              <v-radio-group v-model="changeAuthority" column>
                <v-radio :value="authority" :label="authority.name" v-for="(authority, i) in authorities" :key="i"/>
              </v-radio-group>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="authorityDialog = false;selectedUser = ''">
                戻る
              </v-btn>
              <v-btn class="white--text" color="#295caa" @click="onChangeAuthority">
                変更
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
      <v-row justify="center">
        <v-dialog v-model="deleteDialog" scrollable max-width="300px">
          <v-card>
            <v-card-title>プロジェクトからユーザーを削除</v-card-title>
            <v-divider></v-divider>
            <v-card-text class="pt-2">
              このユーザーをプロジェクトから削除します。よろしいですか？
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="deleteDialog = false;selectedUser = ''">
                戻る
              </v-btn>
              <v-btn class="white--text" color="red darken-2" @click="onDeleteAuthority">
                削除
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
import { isReadyObj, isEmptyObj, checkStatus } from '~/modules/utils'
import * as model from '~/modules/model'
declare module 'vue/types/vue' {
  interface Vue {
    openAuthorityDialog: () => void
    openDeleteDialog: () => void
  }
}
export default Vue.extend({
  data: () => ({
    addUserDialog: false,
    authorityDialog: false,
    deleteDialog: false,
    changeAuthority: {} as model.Authority,
    selectedUser: {},
    projectImageError: false
    
  }),
  computed: {
    ...mapGetters([
      'user',
      'organization',
      'project',
      'projectAuthority',
    ]),
    checkStatus,
    isReadyObj,
    isEmptyObj,
    userOptions() {
      return [
        {title: '権限変更', event: this.openAuthorityDialog},
        {title: '削除', event: this.openDeleteDialog},
      ]
    },
    unregisteredUsers() {
      return this.organization.organization.users.filter((oUser: model.OrganizationAuthority) => {
        return !this.project.authority_users.some((user: model.ProjectAuthority) => user.user_id == oUser.user_id)
      });
    },
    authorities() {
      return this.$store.getters['task/authorities']
    },
    projectImage() {
      return this.projectImageError ? require('~/assets/image/project.png') : this.$config.mediaURL + '/projects/' + this.project.image
    }
  },
  created() {
    const itemStr = sessionStorage.getItem(location.host + window.$nuxt.$route.params.id);
    if(!itemStr) return;
    const item = JSON.parse(itemStr);
  },
  methods: {
    openAuthorityDialog(authorityUser: model.ProjectAuthority) {
      this.selectedUser = authorityUser;
      this.changeAuthority = authorityUser.type;
      this.authorityDialog = true;
    },
    openDeleteDialog(authorityUser: model.ProjectAuthority) {
      this.selectedUser = authorityUser;
      this.deleteDialog = true;
    },
    async onChangeAuthority() {
      const selectedUser = {...this.selectedUser, ...{auth_id: this.changeAuthority.id, type: this.changeAuthority}}
      let response
      try {
        response = await this.$store.dispatch('project/updateProjectAuthority', selectedUser);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
        this.authorityDialog = false;
        this.selectedUser = {} as model.ProjectAuthority;
        }, () => {
          alert('エラーです。');
        })
      }
    },
    async onAddUser() {
      const selectedUser = {...this.selectedUser, ...{id: 0, auth_id: 2, project_id: Number(this.$route.params.id), type: {id: 2, name: '一般'}}}
      let response
      try {
        console.log(selectedUser)
        response = await this.$store.dispatch('project/createProjectAuthority', selectedUser);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
        this.addUserDialog = false;
        this.selectedUser = {} as model.ProjectAuthority;
        }, () => {
          alert('エラーです。');
        })
      }
    },
    async onDeleteAuthority() {
      let response
      try {
        response = await this.$store.dispatch('project/deleteProjectAuthority', this.selectedUser);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
        this.deleteDialog = false;
        this.selectedUser = {} as model.ProjectAuthority;
        }, () => {
          alert('エラーです。');
        })
      }
    },
  }
})
</script>