<template>
  <div>
    <v-row align="center" class="pt-8 px-4">
      <v-col class="d-flex" cols="12" sm="3">
        <v-select
          v-model="selectAssignee"
          :items="assigneeItems"
          item-text="user.name"
          item-value="user.name"
          label="担当者"
          outlined
          clearable
          item-disabled=""
          :disabled="noTask"
          color="#295caa"
        ></v-select>
      </v-col>
      <v-col class="d-flex" cols="12" sm="3">
      <v-select
          v-model="selectField"
          :items="project.fields"
          item-text="name"
          item-value="name"
          label="分野"
          outlined
          clearable
          color="#295caa"
          :disabled="noTask || noField"
        ></v-select>
      </v-col>
      <v-col class="d-flex" cols="12" sm="3">
      <v-select
          v-model="selectMilestone"
          :items="project.milestones"
          item-text="name"
          item-value="name"
          label="マイルストーン"
          outlined
          clearable
          color="#295caa"
          :disabled="noTask || noMilestone"
        ></v-select>
      </v-col>
      <v-col class="d-flex" cols="12" sm="3">
      <v-select
          v-model="selectVersion"
          :items="project.versions"
          item-text="name"
          item-value="name"
          label="バージョン"
          outlined
          clearable
          color="#295caa"
          :disabled="noTask || noVersion"
        ></v-select>
      </v-col>
      <v-col class="d-flex" cols="12" sm="3">
      <v-select
        v-model="selectStatus"
        :items="statuses"
        item-text="name"
        item-value="name"
        label="状況"
        multiple
        outlined
        clearable
        color="#295caa"
        :disabled="noTask"
      >
      <template v-slot:prepend-item>
        <v-list-item @mousedown.prevent>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn text color="cyan darken-2" @click="selectProgress">作業</v-btn>
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn text color="brown lighten-2" @click="selectPending">確認</v-btn>
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn text color="amber lighten-4" @click="selectCompletion">完了</v-btn>
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-content>
            <v-list-item-title>
              <v-btn text color="red darken-2" @click="reset">解除</v-btn>
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
import * as model from '~/modules/model'

declare module 'vue/types/vue' {
  interface Vue {
    storeCondition: (v: {}) => void;
  }
}
export default Vue.extend({
  props: {
    noTask: Boolean
  },
  data: () => ({
    tab: null,
  }),
  computed: {
    ...mapGetters('task', [
      'statuses',
    ]),
    ...mapGetters([
      'project',
    ]),
    isReadyObj,
    isEmptyArr,
    noField() {
      if(this.isReadyObj(this.project)) {
        return this.isEmptyArr(this.project.fields);
      }
      return false;
    },
    noMilestone() {
      if(this.isReadyObj(this.project)) {
        return this.isEmptyArr(this.project.milestones);
      }
      return false;
    },
    noVersion() {
      if(this.isReadyObj(this.project)) {
        return this.isEmptyArr(this.project.versions);
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
    selectMilestone: {
      get() {
        return this.$store.getters['task/selectMilestone'];
      },
      set(milestone) {
        this.$store.commit('task/selectMilestone', milestone);
      }
    },
    selectVersion: {
      get() {
        return this.$store.getters['task/selectVersion'];
      },
      set(version) {
        this.$store.commit('task/selectVersion', version);
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
    assigneeItems() {
      return this.project.authority_users.filter((user: model.OrganizationAuthority) => {
        return user.user.name;
      })
    }
  },
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