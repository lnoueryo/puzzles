<template>
  <v-hover v-slot="{ hover }">
    <v-card :elevation="hover ? 12 : 2" :class="{ 'on-hover': hover }" width="315" height="300" class="ma-4">
      <nuxt-link :to="{name: 'project-id-task', params: {id: id}}">
        <v-img :aspect-ratio="16/9" :src="isPicture(image)" style="background-color: #00000040">
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
              <v-btn class="mr-1" icon v-bind="attrs" v-on="on" v-if="i < 4">
                <v-avatar size="36px">
                  <img alt="Avatar" :src="$config.mediaURL + '/users/' + _.user.image" v-if="_.user.image">
                  <v-icon size="44px" dark v-else>
                    mdi-account-circle
                  </v-icon>
                </v-avatar>
              </v-btn>
            </template>
            <span>{{ _.user.name }}</span>
          </v-tooltip>
          <v-btn class="mr-1" icon v-if="5 < authority_users.length" style="background-color: #ffffff1f">
            <v-avatar size="36px">
              <span>+{{ authority_users.length - 4 }}</span>
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
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  props: ['id', 'image', 'user', 'name', 'authority_users'],
  methods: {
    isPicture(src: string) {
      if(src) {
        return this.$config.mediaURL + '/projects/' + src;
      }
      return require('~/assets/image/project.png');
    }
  },
})
</script>