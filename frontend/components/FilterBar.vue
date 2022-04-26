<template>
  <div>
    <v-row align="center" class="pt-8 px-4">
      <v-col class="d-flex" cols="12" sm="4">
      <v-select
          v-model="selectField"
          :items="projectAuthority.project.fields"
          item-text="name"
          item-value="name"
          label="分野"
          outlined
          clearable
          :disabled="noTask || noField"
        ></v-select>
      </v-col>
      <v-col class="d-flex" cols="12" sm="4">
        <v-select
          v-model="selectAssignee"
          :items="projectAuthority.project.authority_users"
          item-text="user.name"
          item-value="user.name"
          label="担当者"
          outlined
          clearable
          item-disabled=""
          :disabled="noTask"
        ></v-select>
      </v-col>
      <v-col class="d-flex" cols="12" sm="4">
      <v-select
        v-model="selectStatus"
        :items="statuses"
        item-text="name"
        item-value="name"
        label="状況"
        multiple
        outlined
        color="indigo--text"
        clearable
        :disabled="noTask"
      >
      <template v-slot:prepend-item>
        <v-list-item @mousedown.prevent>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn @click="selectProgress">作業</v-btn>
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn @click="selectPending">確認</v-btn>
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn @click="reset">解除</v-btn>
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn @click="selectCompletion">完了</v-btn>
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider class="mt-2"></v-divider>
      </template>
      </v-select>
      </v-col>
    </v-row>
  </div>
</template>


<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import {isReadyObj, isEmptyArr} from '~/modules/utils'

declare module 'vue/types/vue' {
  interface Vue {
    storeCondition: (v: {}) => void;
  }
}
export default Vue.extend({
  props: ['noTask'],
  data: () => ({
    tab: null,
  }),
  computed: {
    ...mapGetters('task', [
      'statuses',
      // 'selectStatus',
    ]),
    ...mapGetters([
      'projectAuthority',
    ]),
    isReadyObj,
    isEmptyArr,
    noField() {
      if(this.isReadyObj(this.projectAuthority)) {
        return this.isEmptyArr(this.projectAuthority.project.fields);
      }
      return false;
    },
    selectAssignee: {
      get() {
        return this.$store.getters['task/selectAssignee'];
      },
      set(assignee) {
        this.$store.commit('task/selectAssignee', assignee);
      }
    },
    selectField: {
      get() {
        return this.$store.getters['task/selectField'];
      },
      set(field) {
        this.$store.commit('task/selectField', field);
      }
    },
    selectStatus: {
      get() {
        return this.$store.getters['task/selectStatus'];
      },
      set(status) {
        this.$store.commit('task/selectStatus', status);
      }
    },
  },
  // created() {
  //   const itemStr = sessionStorage.getItem(location.host + this.$route.params.id);
  //   if(!itemStr) return;
  //   const item = JSON.parse(itemStr);
  //   this.$store.commit('task/selectAssignee', item?.assignee);
  //   this.$store.commit('task/selectField', item?.field);
  //   this.$store.commit('task/selectStatus', item?.status);
  // },
  methods: {
    selectProgress() {
      this.$nextTick(() => {
        const status = ['再議', '未対応', '対応中', '調整']
        this.$store.commit('task/selectStatus', status);
      })
    },
    selectPending() {
      this.$nextTick(() => {
        const status = ['相談', '依頼', '中断', '確認']
        this.$store.commit('task/selectStatus', status);
      })
    },
    selectCompletion() {
      this.$nextTick(() => {
        const status = ['完了']
        this.$store.commit('task/selectStatus', status);
      })
    },
    reset() {
      this.$nextTick(() => {
        const resetItems = [] as string[]
        this.$store.commit('task/selectStatus', resetItems);
      })
    },
  }
})
</script>