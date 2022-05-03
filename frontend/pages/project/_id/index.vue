<template>
  <div v-if="pageReady">
    <v-app-bar
      dense
      dark
      height="80"
    >
      <div class="text-center d-flex justyfy-space-between" style="max-width: 200px;width: 100%">
        <v-btn class="mr-3" :to="'/project/' + project.id + '/create'" color="#295caa">
          <v-icon left>mdi-clipboard-plus-outline</v-icon>
          タスク作成
        </v-btn>
      </div>
      <v-spacer></v-spacer>
      <v-tabs
        v-model="tabKey"
        centered
        dark
        icons-and-text
        fixed-tabs
        color="#295caa"
        class="px-6"
        style="width: 500px"
      >
        <v-tabs-slider></v-tabs-slider>

        <v-tab :href="'#' + tab.component" v-for="(tab, i) in tabs" :key="i" @click="$router.replace({...$route, ...{query: {tab: tab.component}}})">
          {{ tab.title }}
          <v-icon>{{ tab.icon }}</v-icon>
        </v-tab>

      </v-tabs>
      <v-spacer></v-spacer>
      <div class="text-right" style="max-width: 200px;width: 100%">
        <!-- <v-btn icon>
          <v-icon>mdi-heart</v-icon>
        </v-btn>

        <v-btn icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>

        <v-menu
          left
          bottom
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </template>

          <v-list>
            <v-list-item
              v-for="n in 5"
              :key="n"
              @click="() => {}"
            >
              <v-list-item-title>Option {{ n }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu> -->
      </div>
    </v-app-bar>

    <v-tabs-items v-model="tabKey">
      <v-tab-item v-for="(tab, i) in tabs" :key="i + 1" :value="tab.component">
        <keep-alive>
          <div :is="tab.component"></div>
        </keep-alive>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import {isReadyObj, isEmptyObj} from '~/modules/utils'
export default Vue.extend({
  data: () => ({
    tabKey: 'task',
    pageReady: false,
    tabs: [
      {title: '全てのタスク', icon: 'mdi-clipboard-check-multiple-outline', component: 'task'},
      {title: 'プロジェクトの概要', icon: 'mdi-clipboard-check-multiple-outline', component: 'project'},
    ],
    addUserDialog: false,
  }),
  computed: {
    ...mapGetters([
      'user',
      'project',
      'projectIndex',
      'projectReady',
    ]),
    isReadyObj,
    isEmptyObj,
  },
  created() {
    if('tab' in this.$route.query === false) {
      const query = {tab: 'task'};
      this.$router.replace({query});
    }
    this.tabKey = this.$route.query.tab as string;
    let timer = setInterval(() => {
      if(!this.projectReady) return;
      clearInterval(timer);
      if(this.projectIndex === -1) return this.$router.push('/');
      this.pageReady = true;
    }, 100)
  }
})
</script>
