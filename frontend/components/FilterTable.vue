<template>
  <div>
    <div>
      <filter-bar :noTask="noTask"></filter-bar>
    </div>
    <div class="text-center" style="overflow: scroll">
      <div class="thead"  :style="table.thead.style">
        <!-- #006098 -->
        <!-- #009879; -->
        <div class="tr d-flex">
          <div class="cell" :style="cell.header.style" v-for="(cell, i) in table.cells" :key="i" @click="sortTask(i)">
            <v-btn color="transparent" elevation="0" :plain="noTask" :disabled="noTask">
              {{cell.header.title}}
              <v-icon right :class="[{'rotate' : cell.header.active === 2},{'hide' : cell.header.active === 0}]" small v-if="cell.name != 'key' && !noTask">mdi-arrow-up</v-icon>
            </v-btn>
          </div>
        </div>
      </div>
      <div v-if="!noTask">
        <div class="py-16" v-if="isEmptyObj(project) && isEmptyArr(tasks)">
          <v-skeleton-loader
            class="mx-auto"
            type="table"
            style="height: 450px;width: 100%;min-width: 1560px;"
            v-if="isEmptyArr(tasks)"
          >
          </v-skeleton-loader>
          <div style="position: absolute;top: 50%;left: 50%;">
            <v-progress-circular
              :size="50"
              color="primary"
              indeterminate
            ></v-progress-circular>
          </div>
        </div>
        <div class="py-8" v-if="isReadyObj(project) && isEmptyArr(tasks)">
          <div>
            一致する検索結果が得られませんでした。
          </div>
        </div>
        <div class="tbody" :style="table.tbody.style" v-if="isReadyArr(tasks)">
          <div class="tr d-flex align-center" v-for="(task, i) in tasks" :key="i">
            <nuxt-link :to="{name: 'project-id-task-key', params: {id: $route.params.id, key: task.id}}" v-for="(cell, i) in table.cells" :key="i">
              <user-cell :styleValue="cell.header.style" :user="task.assigner" v-if="cell.name == 'assigner'"></user-cell>
              <user-cell :styleValue="cell.header.style" :user="task.assignee" v-else-if="cell.name == 'assignee'"></user-cell>
              <div class="cell" :style="cell.header.style" v-else>{{ task[cell.name] }}</div>
            </nuxt-link>
          </div>
        </div>
      </div>
    </div>
    <div v-if="!noTask">
      <v-row>
        <v-col class="d-flex" cols="12" sm="2">
          <v-select
            v-model="selectList"
            :items="listNumArr"
            outlined
          ></v-select>
        </v-col>
        <v-col class="d-flex" cols="12" sm="2">
          <v-btn icon @click="pageChange(-1)" :disabled="disabledPrevious">
            <v-icon>mdi-arrow-left</v-icon>
          </v-btn>
        </v-col>
        <v-col class="d-flex" cols="12" sm="2">
          <v-btn icon @click="pageChange(1)" :disabled="disabledNext">
            <v-icon>mdi-arrow-right</v-icon>
          </v-btn>
        </v-col>
        <v-col class="d-flex" cols="12" sm="2">
          <span> {{ currentDisplayedTasksNum }} of {{ totalTasks }} </span>
        </v-col>
      </v-row>
    </div>
    <div v-else>
      <v-row class="py-16" style="position: relative;" justify="center">
        <v-btn :to="'/project/' + project.id + '/create'" color="#295caa">新しいタスクを作成する</v-btn>
      </v-row>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import {isReadyArr, isReadyObj, isEmptyArr, isEmptyObj} from '~/modules/utils'
export default Vue.extend({
  data: () => ({
    tab: null,
  }),
  computed: {
    ...mapGetters('task', [
      'statuses',
      'tasks',
      'allTasks',
      'table',
      'listNumArr',
      'pageIndex',
      'currentDisplayedTasksNum',
      'totalTasks',
      'totalPageNum',
    ]),
    ...mapGetters([
      'project',
    ]),
    isReadyArr,
    isReadyObj,
    isEmptyArr,
    isEmptyObj,
    noTask() {
      if(this.isReadyObj(this.project)) {
        return this.isEmptyArr(this.allTasks);
      }
      return false;
    },
    disabledPrevious() {
      return this.pageIndex === 0;
    },
    disabledNext() {
      return this.pageIndex === this.totalPageNum;
    },
    selectList: {
      get() {
        return this.$store.getters['task/selectList'];
      },
      set(v) {
        this.$store.dispatch('task/listIndex', v);
      },
    },
  },
  methods: {
    sortTask(index: number) {
      if(index == 0) return;
      this.$store.commit('task/sortTask', index);
    },
    pageChange(index: number) {
      this.$store.commit('task/pageChange', index);
    },
  }
})
</script>
<style>
.cell {
  display: block;
  position: relative;
  padding: 10px 6px;
  font-size: 13px;
}
.tbody .tr:hover {
  color: #295caa;
  background-color: #295daa18;
  transition: all .2s;
}

::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 5px;
  height: 10px;
  padding: 10px;
}

::-webkit-scrollbar-thumb {
  border-radius: 4px;
  background-color: #285caa2a;
  box-shadow: 0 0 1px rgba(255, 255, 255, .5);
}
.hide {
  opacity: 0
}
.cell:hover .hide {
  opacity: .2
}
.rotate {
  transform: rotate(180deg);
}
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>