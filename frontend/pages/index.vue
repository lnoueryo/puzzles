<template>
  <div v-if="isReadyObj(organization)">
    <v-carousel cycle height="100" hide-delimiters show-arrows-on-hover interval="4500" class="mb-4">
      <v-carousel-item v-for="(slide, i) in slides" :key="i">
        <v-sheet :color="colors[i]" height="100%">
          <v-row class="fill-height" align="center" justify="center">
            <div class="text-h2">
              {{ slide }} Slide
            </div>
          </v-row>
        </v-sheet>
      </v-carousel-item>
    </v-carousel>
    <v-row justify="center" align="center" class="my-6">
      <v-avatar size="36px" v-if="organization.image">
        <img alt="Avatar" :src="$config.mediaURL + '/organizations/' + organization.image">
      </v-avatar>
      <v-icon size="36px" v-else>
        mdi-account-group
      </v-icon>
      <strong class="mx-2" style="font-size: 30px">{{ organization.name }}</strong>
    </v-row>
    <v-row class="my-10" justify="center">
      <v-icon>mdi-cog-play-outline</v-icon>
      <strong class="mx-2" style="font-size: 30px">プロジェクト</strong>
    </v-row>
    <v-row class="my-10" justify="center" v-if="projectSlides.length != 0">
      <v-carousel height="375" hide-delimiter-background :show-arrows="projectSlides.length > 1" show-arrows-on-hover>
        <v-carousel-item v-for="(projects, i) in projectSlides" :key="i">
          <v-sheet height="100%" color="transparent">
            <v-row class="fill-height" align="center" justify="center">
              <!-- <hover-card :key="i" v-bind="project"></hover-card> -->
            <template v-for="(_, i) in projects">
              <hover-card :key="i" v-bind="_"></hover-card>
            </template>
            </v-row>
          </v-sheet>
        </v-carousel-item>
      </v-carousel>
    </v-row>
    <v-row class="my-2" justify="center" v-else>
      <v-btn to="/project/create" color="#295caa">プロジェクトの作成</v-btn>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { isReadyObj } from '~/modules/utils'
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
  }),
  computed: {
    isReadyObj,
    organization() {
      return this.$store.getters['organization'].organization;
    },
    projectSlides() {
      return this.$store.getters['projectSlides'];
    },
  },
})
</script>
<style scoped>
.v-card {
  transition: all .2s ease-in-out;
  background-color: #FF8F00;
}

.v-card:not(.on-hover) {
  background-color: #272727;
 }

.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}
</style>