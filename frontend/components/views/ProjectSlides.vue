<template>
  <div>
    <v-row class="my-10" justify="center">
      <v-icon>mdi-cog-play-outline</v-icon>
      <strong class="mx-2">プロジェクト</strong>
    </v-row>
    <v-row
      class="my-10"
      justify="center"
      v-if="projectSlides.length != 0"
    >
      <v-carousel
        height="375"
        hide-delimiter-background
        :show-arrows="projectSlides.length > 1"
        show-arrows-on-hover
      >
        <v-carousel-item v-for="(projects, i) in projectSlides" :key="i">
          <v-sheet height="100%" color="transparent">
            <v-row
              class="fill-height"
              align="center"
              justify="center"
            >
            <template v-for="(_, i) in projects">
              <hover-card
                :key="i"
                v-bind="_"
                :user="user"
              />
            </template>
            </v-row>
          </v-sheet>
        </v-carousel-item>
      </v-carousel>
    </v-row>
    <v-row
      class="my-2"
      justify="center"
      v-else
    >
      <v-btn
        to="/project/create"
        color="#295caa"
        v-if="isAdmin"
      >
        プロジェクトの作成
      </v-btn>
      <small v-else>※現在参加しているプロジェクトがありません</small>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import * as model from '~/modules/model'
declare module 'vue/types/vue' {
  interface Vue {
    organizationAuthority: model.OrganizationAuthority;
  }
}
export default Vue.extend({
  computed: {
    user() {
      return this.$store.getters['user'];
    },
    organizationAuthority() {
      return this.$store.getters['organization'];
    },
    projectSlides() {
      return this.$store.getters['projectSlides'];
    },
    isAdmin() {
      return this.organizationAuthority.type.name == '管理者';
    }
  }
})
</script>