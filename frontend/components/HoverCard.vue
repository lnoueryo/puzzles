<template>
  <div>
    <v-hover v-slot="{ hover }">
      <v-card :elevation="hover ? 12 : 2" :class="{ 'on-hover': hover }" width="315" height="300" class="ma-4">
        <nuxt-link :to="{name: 'project-id', params: {id: id}}">
          <v-img :aspect-ratio="16/9" :src="projectImage" @error="projectImageError = true" style="background-color: #00000040">
            <template v-slot:placeholder>
              <v-row class="fill-height ma-0" align="center" justify="center">
                <v-progress-circular indeterminate color="grey lighten-5" />
              </v-row>
            </template>
            <div style="background-color: #00000040" class="fill-height repeating-gradient">
              <div class="d-flex justify-end">
                <v-btn icon :class="{ 'show-btns': hover }" color="transparent">
                  <v-icon>mdi-star</v-icon>
                </v-btn>
                <v-btn icon :class="{ 'show-btns': hover }" color="transparent">
                  <v-icon>mdi-pin</v-icon>
                </v-btn>
              </div>
              <v-card-title class="text-h6 white--text">
                <v-row class="fill-height flex-column" justify="space-between">
                  <p class="ml-4 mt-4 subheading text-left">
                    <strong>{{ name }}</strong>
                  </p>
                </v-row>
              </v-card-title>
            </div>
          </v-img>
        </nuxt-link>
        <div class="d-flex align-center justify-space-between ma-2">
          <h3>メンバー：</h3>
          <div>
            <v-tooltip bottom v-for="(_, i) in authority_users" :key="i">
              <template v-slot:activator="{ on, attrs }">
                <v-btn class="mr-1" icon v-bind="attrs" v-on="on" v-if="i < 4" @click="$router.push({name: 'profile-user_id', params: {user_id: _.user_id}})">
                  <v-avatar style="object-fit: cover;" size="36px">
                    <v-img alt="Avatar" :src="$config.mediaURL + '/users/' + _.user.image" v-if="_.user.image">
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
                </v-btn>
              </template>
              <span>{{ _.user.name }}</span>
            </v-tooltip>
            <v-btn class="mr-1" icon v-if="5 < authUserLength" style="background-color: #ffffff1f" @click="usersDialog = true">
              <v-avatar size="36px">
                <span>+{{ authUserLength - 4 }}</span>
              </v-avatar>
            </v-btn>
          </div>
        </div>
        <div style="margin-top: auto">
          <v-btn text :class="{ 'show-btns': hover }" color="transparent">
            <v-icon left>
              mdi-pencil
            </v-icon>
              ダッシュボード
          </v-btn>
          <v-btn text :class="{ 'show-btns': hover }" color="transparent">
            <v-icon left>
              mdi-plus
            </v-icon>
              課題追加
          </v-btn>
        </div>
      </v-card>
    </v-hover>
    <v-row justify="space-around">
      <v-col cols="auto">
        <v-dialog
          transition="dialog-bottom-transition"
          max-width="550"
          v-model="usersDialog"
        >
          <template v-slot:default="dialog">
            <v-card>
              <v-toolbar color="#295caa" dark>
                <div class="text-h5" style="margin: auto">
                  {{name}}
                </div>
                <v-btn icon absolute right @click="usersDialog = false">
                  <v-icon small>mdi-close</v-icon>
                </v-btn>
              </v-toolbar>
              <v-list subheader two-line>
                <v-list-item v-for="authUser in authority_users" :key="authUser.id"  @click="toProfile(authUser)" link class="px-8">
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
                    <v-list-item-title v-text="authUser.user.name" v-if="authUser.user.name"></v-list-item-title>
                    <v-list-item-title v-text="'招待中'" style="color: red" v-else></v-list-item-title>
                    <v-list-item-subtitle v-text="authUser.user.email"></v-list-item-subtitle>
                  </v-list-item-content>

                  <v-list-item-content class="text-center">
                    <v-list-item-subtitle v-text="authUser.type.name"></v-list-item-subtitle>
                  </v-list-item-content>

                </v-list-item>
              </v-list>
              <v-card-actions class="justify-end">
                <v-btn
                  text
                  @click="dialog.value = false"
                >閉じる</v-btn>
              </v-card-actions>
            </v-card>
          </template>
        </v-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import * as model from '~/modules/model'
export default Vue.extend({
  props: {
    id: Number,
    image: String,
    user: Object,
    name: String,
    authority_users: Object
  },
  data: () => ({
    projectImageError: false,
    usersDialog: false,
  }),
  computed: {
    authUserLength() {
      return this.authority_users?.length || 0;
    },
    projectImage() {
      return this.projectImageError ? require('~/assets/image/project.png') : this.$config.mediaURL + '/projects/' + this.image
    }
  },
  methods: {
    isPicture(src: string) {
      if(src) {
        return this.$config.mediaURL + '/projects/' + src;
      }
      return require('~/assets/image/project.png');
    },
    toProfile(authUser: model.OrganizationAuthority) {
      if(!authUser.user.name) return;
      if(this.user.id == authUser.user_id) {
        this.$router.push({name: 'profile'});
        return;
      }
      this.$router.push({name: 'profile-user_id', params: {user_id: String(authUser.user_id)}})
    }
  },
})
</script>
<style scoped>
.v-card {
  transition: all .2s ease-in-out;
  background-color: #295daa6e;
}

.v-card:not(.on-hover) {
  background-color: #272727;
 }

/* .show-btns {
  color: rgba(255, 255, 255, 1) !important;
} */
</style>