<template>
  <div v-if="pageReady">
    <v-app-bar
      dense
      dark
      height="80"
    >
      <div class="text-center d-flex justyfy-space-between" style="max-width: 200px;width: 100%">
        <v-btn class="mr-3" :to="'/project/' + projectAuthority.project_id + '/create'" color="#295caa">
          <v-icon>mdi-clipboard-plus-outline</v-icon>
          タスク作成
        </v-btn>
        <v-btn :to="{name: 'project-id-edit', params: {id: $route.params.id}}" color="#295caa">
          <v-icon>mdi-clipboard-plus-outline</v-icon>
          プロジェクト編集
        </v-btn>
      </div>
      <v-spacer></v-spacer>
      <v-tabs
        v-model="tab"
        centered
        dark
        icons-and-text
        fixed-tabs
        color="#295caa"
        class="px-6"
        style="width: 500px"
      >
        <v-tabs-slider></v-tabs-slider>

        <v-tab href="#tab-1">
          全てのタスク
          <v-icon>mdi-clipboard-check-multiple-outline</v-icon>
        </v-tab>

        <v-tab href="#tab-2">
          あなたのタスク
          <v-icon>mdi-account-box</v-icon>
        </v-tab>

        <v-tab href="#tab-3">
          プロジェクト概要
          <v-icon>mdi-account-box</v-icon>
        </v-tab>

      </v-tabs>
      <v-spacer></v-spacer>
      <div class="text-right" style="max-width: 200px;width: 100%">
      <v-btn icon>
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
      </v-menu>
      </div>
    </v-app-bar>

    <v-tabs-items v-model="tab">
      <v-tab-item v-for="i in 3" :key="i" :value="'tab-' + i">
        <filter-table></filter-table>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import {isReadyObj, isEmptyObj, checkStatus} from '~/modules/utils'
export default Vue.extend({
  data: () => ({
    tab: null,
    pageReady: false,
  }),
  computed: {
    ...mapGetters([
      'user',
      'projectAuthority',
      'project',
      'projectIndex',
      'projectReady',
    ]),
    isReadyObj,
    isEmptyObj,
    checkStatus,
  },
  created() {
    let timer = setInterval(() => {
      if(!this.projectReady) return;
      clearInterval(timer);
      if(this.projectIndex === -1) return this.$router.push('/');
      this.pageReady = true;
    }, 100)
  }
})
</script>
