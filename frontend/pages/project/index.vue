<template>
  <div>
    <v-carousel
      cycle
      height="100"
      hide-delimiters
      show-arrows-on-hover
      interval="4500"
      class="mb-4"
    >
      <v-carousel-item
        v-for="(slide, i) in slides"
        :key="i"
      >
        <v-sheet
          :color="colors[i]"
          height="100%"
        >
          <v-row
            class="fill-height"
            align="center"
            justify="center"
          >
            <div class="text-h2">
              {{ slide }} Slide
            </div>
          </v-row>
        </v-sheet>
      </v-carousel-item>
    </v-carousel>
    <v-row justify="center" align="center" class="my-6">
      <v-avatar size="36px" v-if="organization.image">
        <img
          class="object-cover"
          alt="Avatar"
          :src="$config.mediaURL + '/organizations/' + organization.image"
        >
      </v-avatar>
      <v-icon size="36px" v-else>
        mdi-account-group
      </v-icon>
      <strong class="mx-2 organization-font">{{ organization.name }}</strong>
    </v-row>
    <v-row class="my-10" justify="center">
      <v-icon>mdi-cog-play-outline</v-icon>
      <strong class="mx-2 organization-font">プロジェクト</strong>
    </v-row>
    <v-row class="my-10" justify="center" v-if="projectSlides.length != 0">
      <v-carousel height="375" hide-delimiter-background :show-arrows="projectSlides.length > 1" show-arrows-on-hover>
        <v-carousel-item v-for="(projects, i) in projectSlides" :key="i">
          <v-sheet height="100%" color="transparent">
            <v-row class="fill-height" align="center" justify="center">
            <template v-for="(_, i) in projects">
              <hover-card :key="i" v-bind="_"></hover-card>
              <!-- <v-hover v-slot="{ hover }" :key="i">
                <v-card :elevation="hover ? 12 : 2" :class="{ 'on-hover': hover }" width="315" height="300" class="ma-4">
                  <nuxt-link :to="{name: 'project-id-task', params: {id: _.project.id}}">
                    <v-img :src="isPicture(_.project.image)" height="150px" style="background-color: #00000040">
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
                          <v-row
                            class="fill-height flex-column"
                            justify="space-between"
                          >
                            <p class="ml-4 mt-4 subheading text-left">
                              <strong>{{ _.project.name }}</strong>
                            </p>
                          </v-row>
                        </v-card-title>
                      </div>
                    </v-img>
                  </nuxt-link>
                  <div class="d-flex align-center justify-space-between ma-2">
                    <h3>メンバー：</h3>
                    <div>
                      <v-tooltip bottom v-for="(_, i) in _.project_users" :key="i">
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn
                            class="mr-1"
                            icon
                            v-bind="attrs"
                            v-on="on"
                            v-if="i < 4"
                          >
                            <v-avatar
                              size="36px"
                            >
                              <img
                                alt="Avatar"
                                :src="'http://localhost:8080/media/users/' + _.user.image"
                                v-if="_.user.image"
                              >
                              <v-icon size="44px" dark v-else>
                                mdi-account-circle
                              </v-icon>
                            </v-avatar>
                          </v-btn>
                        </template>
                        <span>{{ _.user.name }}</span>
                      </v-tooltip>
                      <v-btn
                        class="mr-1"
                        icon
                        v-if="5 < _.project_users.length"
                        style="background-color: #ffffff1f"
                      >
                        <v-avatar
                          size="36px"
                        >
                          <span>+{{ _.project_users.length - 4 }}</span>
                        </v-avatar>
                      </v-btn>
                    </div>
                  </div>
                  <div style="margin-top: auto">
                    <v-btn
                      text
                      :class="{ 'show-btns': hover }"
                      color="transparent"
                    >
                    <v-icon left>
                      mdi-pencil
                    </v-icon>
                      ダッシュボード
                    </v-btn>
                    <v-btn
                      text
                      :class="{ 'show-btns': hover }"
                      color="transparent"
                    >
                    <v-icon left>
                      mdi-plus
                    </v-icon>
                      課題追加
                    </v-btn>
                  </div>
                </v-card>
              </v-hover> -->
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
    organization() {
      return this.$store.getters['organization']
    },
    projectSlides() {
      return this.$store.getters['projectSlides']
    },
  },
})
</script>
<style lang="scss" scoped>
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
.organization-font {
  font-size: 30px;
}
</style>