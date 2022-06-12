<template>
  <div v-if="isReadyObj(organization)">
    <v-app-bar
      dense
      dark
      height="80"
    >
      <v-spacer></v-spacer>
      <v-tabs
        v-model="tabKey"
        centered
        dark
        icons-and-text
        fixed-tabs
        color="#295caa"
        class="px-6 tab-width"
      >
        <v-tabs-slider></v-tabs-slider>

        <v-tab :href="'#tab-' + (i + 1)" v-for="(tab, i) in tabs" :key="i">
          {{ tab.title }}
          <v-icon>{{ tab.icon }}</v-icon>
        </v-tab>

      </v-tabs>
      <v-spacer></v-spacer>
    </v-app-bar>
    <div>

    </div>
    <v-row justify="center" align="center" class="py-8 bc2">
      <v-avatar size="36px" v-if="organization.image">
        <v-img class="object-cover" alt="Avatar" :src="organizationImage" @error="organizationImageError = true">
          <template v-slot:placeholder>
            <v-row class="fill-height ma-0" align="center" justify="center">
              <v-progress-circular indeterminate color="grey lighten-5" />
            </v-row>
          </template>
        </v-img>
      </v-avatar>
      <v-icon size="36px" v-else>
        mdi-account-group
      </v-icon>
      <strong id="organization-name" class="mx-2 organization-font">{{ organization.name }}</strong>
    </v-row>
    <v-tabs-items v-model="tabKey">
      <v-tab-item :value="'tab-1'">
        <v-row class="my-10" justify="center">
          <v-icon>mdi-cog-play-outline</v-icon>
          <strong class="mx-2">プロジェクト</strong>
        </v-row>
        <v-row class="my-10" justify="center" v-if="projectSlides.length != 0">
          <v-carousel height="375" hide-delimiter-background :show-arrows="projectSlides.length > 1" show-arrows-on-hover>
            <v-carousel-item v-for="(projects, i) in projectSlides" :key="i">
              <v-sheet height="100%" color="transparent">
                <v-row class="fill-height" align="center" justify="center">
                <template v-for="(_, i) in projects">
                  <hover-card :key="i" v-bind="_" :user="user"></hover-card>
                </template>
                </v-row>
              </v-sheet>
            </v-carousel-item>
          </v-carousel>
        </v-row>
        <v-row class="my-2" justify="center" v-else>
          <v-btn to="/project/create" color="#295caa" v-if="organizationAuthority.type.name == '管理者'">プロジェクトの作成</v-btn>
          <small v-else>※現在参加しているプロジェクトがありません</small>
        </v-row>
      </v-tab-item>
      <v-tab-item :value="'tab-2'">
        <div class="project-container" v-if="isReadyObj(organization)">
          <v-row class="py-8 relative" align="center" justify="center">
            <h2>{{ organization.name }}</h2>
            <v-btn absolute right color="#295caa"  v-if="organizationAuthority.type.name == '管理者'" :to="{name: 'organization-edit'}">
              <v-icon left>mdi-application-edit-outline</v-icon>
              編集
            </v-btn>
          </v-row>
          <v-row class="px-4" align="center" justify="center">
            <v-img :aspect-ratio="16/9" :src="organizationImage" @error="organizationImageError = true">
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
            <div v-if="organization.description">{{ organization.description }}</div>
            <div v-else>なし</div>
          </v-row>
          <v-row class="py-8" align="center" justify="center">
            <div class="d-flex justify-center align-center relative w100">
              <h2>ユーザー</h2>
              <v-btn absolute right color="#295caa" v-if="organizationAuthority.type.name == '管理者'" @click="$router.push({name: 'organization-create'})">
                <v-icon left>mdi-account-multiple-plus</v-icon>
                メンバー追加
              </v-btn>
            </div>
            <v-list subheader two-line>
              <v-list-item v-for="authUser in organization.users" :key="authUser.id" @click="toProfile(authUser)" link>
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

                <v-list-item-content class="mr-4 list-container">
                  <v-list-item-title v-text="authUser.user.name" v-if="authUser.user.name"></v-list-item-title>
                  <v-list-item-title class="error-color" v-text="'招待中'" v-else></v-list-item-title>
                  <v-list-item-subtitle v-text="authUser.user.email"></v-list-item-subtitle>
                </v-list-item-content>

                <v-list-item-content class="text-center">
                  <v-list-item-subtitle v-text="authUser.type.name"></v-list-item-subtitle>
                </v-list-item-content>

                <v-list-item-action v-if="organizationAuthority.type.name == '管理者'">
                  <v-menu right bottom>
                    <template v-slot:activator="{ on, attrs }">
                      <v-btn icon v-bind="attrs" v-on="on">
                        <v-icon>mdi-dots-vertical</v-icon>
                      </v-btn>
                    </template>

                    <v-list class="user-options">
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
            <v-dialog v-model="authorityDialog" scrollable max-width="300px">
              <v-card>
                <v-card-title>権限変更</v-card-title>
                <v-divider></v-divider>
                <v-card-text class="h100">
                  <v-radio-group v-model="changeAuthority" column>
                    <v-radio color="#295caa" :value="authority" :label="authority.name" v-for="(authority, i) in authorities" :key="i"/>
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
                <v-card-title>組織からユーザーを削除</v-card-title>
                <v-divider></v-divider>
                <v-card-text class="pt-2">
                  このユーザーを組織から削除します。よろしいですか？
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
      </v-tab-item>
    </v-tabs-items>
    <!-- <v-carousel cycle height="100" hide-delimiters show-arrows-on-hover interval="4500" class="mb-4">
      <v-carousel-item v-for="(slide, i) in slides" :key="i">
        <v-sheet :color="colors[i]" height="100%">
          <v-row class="fill-height" align="center" justify="center">
            <div class="text-h2">
              {{ slide }} Slide
            </div>
          </v-row>
        </v-sheet>
      </v-carousel-item>
    </v-carousel> -->
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { isReadyObj, checkStatus } from '~/modules/utils'
import * as model from '~/modules/model'
declare module 'vue/types/vue' {
  interface Vue {
    organization: model.Organization;
    user: model.User
  }
}
export default Vue.extend({
  // layout: 'dashboard',
  data: () => ({
    model: null,
    colors: [
      'indigo',
      'warning',
      'pink darken-2',
      'red lighten-1',
      'deep-purple accent-4',
    ],
    slides: [
      'First',
      'Second',
      'Third',
      'Fourth',
      'Fifth',
    ],
    tabKey: 'tab-1',
    pageReady: false,
    tabs: [
      {title: 'プロジェクト', icon: 'mdi-clipboard-check-multiple-outline', component: 'filter-table'},
      {title: '組織の概要', icon: 'mdi-clipboard-check-multiple-outline', component: 'project'},
    ],
    addUserDialog: false,
    authorityDialog: false,
    deleteDialog: false,
    changeAuthority: {} as model.Authority,
    selectedUser: {},
    organizationImageError: false,
  }),
  computed: {
    isReadyObj,
    checkStatus,
    user() {
      return this.$store.getters['user'];
    },
    organizationAuthority() {
      return this.$store.getters['organization'];
    },
    organization() {
      return this.$store.getters['organization'].organization;
    },
    projectSlides() {
      return this.$store.getters['projectSlides'];
    },
    userOptions() {
      return [
        {title: '権限変更', event: this.openAuthorityDialog},
        {title: '削除', event: this.openDeleteDialog},
      ]
    },
    authorities() {
      return this.$store.getters['task/authorities']
    },
    organizationImage() {
      return this.organizationImageError ? require('~/assets/image/organization.png') : this.$config.mediaURL + '/organizations/' + this.organization.image;
    }
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
        response = await this.$store.dispatch('updateOrganizationAuthority', selectedUser);
      } catch (error: any) {
        response = error.response;
      } finally {
        if('status' in response === false) return this.$router.push('/bad-connection')
        this.checkStatus(response.status, () => {
        this.authorityDialog = false;
        this.selectedUser = {} as model.OrganizationAuthority;
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
    toProfile(authUser: model.OrganizationAuthority) {
      if(!authUser.user.name) return;
      if(this.user.id == authUser.user_id) {
        this.$router.push({name: 'profile'});
        return;
      }
      this.$router.push({name: 'profile-user_id', params: {user_id: String(authUser.user_id)}})
    }
  }
})
</script>
<style lang="scss" scoped>
  .tab-container {
    width: 500px;
  }
  .organization-font {
    font-size: 30px;
  }
  .project-container {
    width: 100%;
    max-width: 600px;
    margin: auto;
  }
  .list-container {
    width: 400px;
  }
</style>