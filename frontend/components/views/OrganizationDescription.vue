<template>
  <div>
    <div class="project-container" v-if="isReadyObj(organization)">
      <v-row
        class="py-8 relative"
        align="center"
        justify="center"
      >
        <h2>{{ organization.name }}</h2>
        <v-btn
          :to="{name: 'organization-edit'}"
          absolute
          right
          color="#295caa"
          v-if="isAdmin"
        >
          <v-icon left>mdi-application-edit-outline</v-icon>
          編集
        </v-btn>
      </v-row>
      <v-row
        class="px-4"
        align="center"
        justify="center"
      >
        <v-img
          :aspect-ratio="16/9"
          :src="organizationImage"
          @error="organizationImageError = true"
        >
          <template v-slot:placeholder>
            <v-row
              class="fill-height ma-0"
              align="center"
              justify="center"
            >
              <v-progress-circular indeterminate color="grey lighten-5" />
            </v-row>
          </template>
        </v-img>
      </v-row>
      <v-row
        class="pt-8"
        align="center"
        justify="center"
      >
        <h2>概要</h2>
      </v-row>
      <v-row
        class="pb-8"
        align="center"
        justify="center"
      >
        <div v-if="organization.description">{{ organization.description }}</div>
        <div v-else>なし</div>
      </v-row>
      <v-row
        class="py-8"
        align="center"
        justify="center"
      >
        <div class="d-flex justify-center align-center relative w100">
          <h2>ユーザー</h2>
          <v-btn
            absolute
            right
            color="#295caa"
            v-if="isAdmin"
            @click="$router.push({name: 'organization-create'})"
          >
            <v-icon left>mdi-account-multiple-plus</v-icon>
            メンバー追加
          </v-btn>
        </div>
        <v-list subheader two-line>
          <v-list-item
            v-for="authUser in organization.users"
            :key="authUser.id"
            @click="$router.push({ name: 'profile-user_id', params: { user_id: authUser.id } })"
          >
            <v-list-item-avatar>
              <v-img :src="userImage(authUser.user.image)" v-if="userImage(authUser.user.image)">
                <template v-slot:placeholder>
                  <v-row
                    class="fill-height ma-0"
                    align="center"
                    justify="center"
                  >
                    <v-progress-circular indeterminate color="grey lighten-5" />
                  </v-row>
                </template>
              </v-img>
              <v-icon size="44px" v-else>
                mdi-account-circle
              </v-icon>
            </v-list-item-avatar>

            <v-list-item-content class="mr-4 list-container">
              <v-list-item-title v-text="authUser.user.name" v-if="authUser.user.name" />
              <v-list-item-title
                class="error-color"
                v-text="'招待中'"
                v-else
              />
              <v-list-item-subtitle v-text="authUser.user.email" />
            </v-list-item-content>

            <v-list-item-content class="text-center">
              <v-list-item-subtitle v-text="authUser.type.name" />
            </v-list-item-content>

            <v-list-item-action v-if="isAdmin">
              <v-menu right bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn icon v-bind="attrs" v-on="on">
                    <v-icon>mdi-dots-vertical</v-icon>
                  </v-btn>
                </template>

                <v-list class="user-options">
                  <v-list-item
                    v-for="(option, i) in userOptions"
                    :key="i" @click="option.event(authUser)"
                    :disabled="authUser.user_id == user.id"
                  >
                    <v-list-item-title>{{ option.title }}</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>
            </v-list-item-action>
          </v-list-item>
        </v-list>
      </v-row>
      <v-row justify="center">
        <v-dialog
          v-model="authorityDialog"
          scrollable
          max-width="300px"
        >
          <v-card>
            <v-card-title>権限変更</v-card-title>
            <v-divider></v-divider>
            <v-card-text class="h100">
              <v-radio-group v-model="changeAuthority" column>
                <v-radio
                  color="#295caa"
                  :value="authority"
                  :label="authority.name"
                  v-for="(authority, i) in authorities"
                  :key="i"
                />
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
        <v-dialog
          v-model="deleteDialog"
          scrollable
          max-width="300px"
        >
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
              <v-btn
                class="white--text"
                color="red darken-2"
                @click="onDeleteAuthority"
              >
                削除
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { isReadyObj } from '~/modules/utils'
import * as model from '~/modules/model'
export default Vue.extend({
  data: () => ({
    authorityDialog: false,
    changeAuthority: {} as model.Authority,
    deleteDialog: false,
    organizationImageError: false,
    selectedUser: {},
  }),
  computed: {
    isReadyObj,
    user() {
      return this.$store.getters['user'];
    },
    organizationAuthority() {
      return this.$store.getters['organizationAuthority'];
    },
    organization() {
      return this.$store.getters['organizationAuthority'].organization;
    },
    isAdmin() {
      return this.organizationAuthority.type.name == '管理者';
    },
    organizationImage() {
      const errorImage = require('~/assets/image/organization.png');
      const organizationImage = this.$config.mediaURL + '/organizations/' + this.organization.image;
      return this.organizationImageError ? errorImage : organizationImage;
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
  },
  methods: {
    userImage(image: string) {
      if(!image) return image;
      return this.$config.mediaURL + '/users/' + image;
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
  }
})
</script>
<style lang="scss" scoped>
  .project-container {
    width: 100%;
    max-width: 600px;
    margin: auto;
  }
  .list-container {
    width: 400px;
  }
</style>