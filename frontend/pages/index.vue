<template>
  <div v-if="isReadyObj(organization)">
    <v-app-bar
      dense
      dark
      height="80"
    >
      <v-spacer></v-spacer>
      <v-tabs
        class="px-6 tab-width"
        v-model="tabKey"
        centered
        dark
        icons-and-text
        fixed-tabs
        color="#295caa"
      >
        <v-tabs-slider></v-tabs-slider>

        <v-tab
          :href="'#tab-' + (i + 1)"
          v-for="(tab, i) in tabs"
          :key="i"
        >
          {{ tab.title }}
          <v-icon>{{ tab.icon }}</v-icon>
        </v-tab>

      </v-tabs>
      <v-spacer></v-spacer>
    </v-app-bar>
    <div>

    </div>
    <v-row
      class="py-8 bc2"
      justify="center"
      align="center"
    >
      <v-avatar size="36px" v-if="organization.image">
        <v-img
          class="object-cover"
          alt="Avatar"
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
      </v-avatar>
      <v-icon size="36px" v-else>
        mdi-account-group
      </v-icon>
      <strong id="organization-name" class="mx-2 organization-font">
        {{ organization.name }}
      </strong>
    </v-row>
    <v-tabs-items v-model="tabKey">
      <v-tab-item :value="'tab-1'">
        <project-slides />
      </v-tab-item>
      <v-tab-item :value="'tab-2'">
        <organization-description />
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
import { isReadyObj } from '~/modules/utils'
import * as model from '~/modules/model'
const ProjectSlides = () => import('~/components/views/ProjectSlides.vue')
const OrganizationDescription = () => import('~/components/views/OrganizationDescription.vue')
declare module 'vue/types/vue' {
  interface Vue {
    organization: model.Organization;
    user: model.User
  }
}
export default Vue.extend({
  components: {
    ProjectSlides,
    OrganizationDescription
  },
  // layout: 'dashboard',
  data: () => ({
    model: null,
    organizationImageError: false,
    pageReady: false,
    tabKey: 'tab-1',
    tabs: [
      {title: 'プロジェクト', icon: 'mdi-clipboard-check-multiple-outline', component: 'filter-table'},
      {title: '組織の概要', icon: 'mdi-clipboard-check-multiple-outline', component: 'project'},
    ],
    // colors: [
    //   'indigo',
    //   'warning',
    //   'pink darken-2',
    //   'red lighten-1',
    //   'deep-purple accent-4',
    // ],
    // slides: [
    //   'First',
    //   'Second',
    //   'Third',
    //   'Fourth',
    //   'Fifth',
    // ],
  }),
  computed: {
    isReadyObj,
    authorities() {
      return this.$store.getters['task/authorities']
    },
    organization() {
      return this.$store.getters['organization'].organization;
    },
    organizationImage() {
      const errorImage = require('~/assets/image/organization.png');
      const organizationImage = this.$config.mediaURL + '/organizations/' + this.organization.image;
      return this.organizationImageError ? errorImage : organizationImage;
    }
  },
})
</script>
<style lang="scss" scoped>
  .tab-container {
    width: 500px;
  }
  .organization-font {
    font-size: 30px;
  }
</style>