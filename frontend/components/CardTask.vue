<template>
  <div>
    <div>
      <FilterBar :noTask="noTask"></FilterBar>
    </div>
    <div class="text-center scroll">
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
            class="mx-auto loader-frame"
            type="table"
            v-if="isEmptyArr(tasks)"
          >
          </v-skeleton-loader>
          <div class="circle-position">
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
            <div v-for="(cell, i) in table.cells" :key="i">
              <!-- selectedUser -->
              <user-cell :styleValue="cell.header.style" :user="task.assigner" v-if="cell.name == 'assigner'" @click.native="openUserDialog(task.assigner.id)"></user-cell>
              <user-cell :styleValue="cell.header.style" :user="task.assignee" v-else-if="cell.name == 'assignee'" @click.native="openUserDialog(task.assignee.id)"></user-cell>
              <!-- <nuxt-link :to="{name: 'profile-user_id', params: {user_id: task.assigner.id}}" v-if="cell.name == 'assigner'">
                <user-cell :styleValue="cell.header.style" :user="task.assigner"></user-cell>
              </nuxt-link>
              <nuxt-link :to="{name: 'profile-user_id', params: {user_id: task.assignee.id}}"  v-else-if="cell.name == 'assignee'">
                <user-cell :styleValue="cell.header.style" :user="task.assignee"></user-cell>
              </nuxt-link> -->
              <nuxt-link :to="{name: 'project-id-task-key-edit', params: {id: $route.params.id, key: task.id}}" v-else>
                <div class="cell" :style="cell.header.style">{{ getName(task[cell.name]) }}</div>
              </nuxt-link>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="!noTask" class="py-4">
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
      <v-row class="py-16 relative" justify="center">
        <v-btn :to="'/project/' + project.id + '/create'" color="#295caa">新しいタスクを作成する</v-btn>
      </v-row>
    </div>
    <v-row justify="space-around">
      <v-col cols="auto">
        <v-dialog v-model="userDialog" transition="dialog-bottom-transition" max-width="600">
          <template v-slot:default="dialog">
            <v-row class="py-8" align="center" justify="center" v-if="isReadyObj(selectedUser)">
              <v-list class="px-4" subheader two-line width="600">
                <v-row class="py-8 mx-4 relative" align="center" justify="center">
                  <div class="relative">
                    <v-avatar size="36px" class="mr-4 avatar-position">
                      <v-img class="object-cover" alt="Avatar" :src="$config.mediaURL + '/users/' + selectedUser.user.image" v-if="selectedUser.user.image">
                        <template v-slot:placeholder>
                          <v-row class="fill-height ma-0" align="center" justify="center">
                            <v-progress-circular indeterminate color="grey lighten-5" />
                          </v-row>
                        </template>
                      </v-img>
                      <v-icon size="44px" dark v-else>
                      mdi-account-circle
                      </v-icon>
                    </v-avatar>
                    <h2>{{ selectedUser.user.name }}</h2>
                  </div>
                  <v-btn icon absolute right color="#295caa" :to="{name: 'profile-edit'}" v-if="user.id == selectedUser.user_id">
                    <v-icon>mdi-application-edit-outline</v-icon>
                  </v-btn>
                </v-row>
                <v-list-item two-line v-for="(userEntity, i) in userEntities" :key="i">
                  <v-list-item-content>
                    <v-list-item-title>{{ userEntity.title }}</v-list-item-title>
                    <v-list-item-subtitle>{{ selectedUser.user[userEntity.key] }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item two-line>
                  <v-list-item-content>
                    <v-list-item-title>組織参加日</v-list-item-title>
                    <v-list-item-subtitle>{{ changeToTimeStampFormat(selectedUser.created_at) }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-spacer></v-spacer>
                <v-btn text @click="dialog.value = false;selectedUser = {}" right>Close</v-btn>
              </v-list>
            </v-row>
          </template>
        </v-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import {isReadyArr, isReadyObj, isEmptyArr, isEmptyObj, changeToTimeStampFormat} from '~/modules/utils'
import * as model from '~/modules/model'
export default Vue.extend({
  data: () => ({
    tab: null,
    userDialog: false,
    selectedUser: {}
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
      'user',
      'project',
    ]),
    isReadyArr,
    isReadyObj,
    isEmptyArr,
    isEmptyObj,
    changeToTimeStampFormat,
    noTask() {
      if(this.isReadyObj(this.project)) {
        return this.isEmptyArr(this.tasks);
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
    userEntities() {
      const userEntities = [
        {key: 'description', title: '自己紹介'},
        {key: 'age', title: '年齢'},
        {key: 'sex', title: '性別'},
        {key: 'email', title: 'メールアドレス'},
        {key: 'address', title: '住所'},
      ]
      return userEntities
    }
  },
  methods: {
    sortTask(index: number) {
      if(index == 0) return;
      this.$store.commit('task/sortTask', index);
    },
    pageChange(index: number) {
      this.$store.commit('task/pageChange', index);
    },
    openUserDialog(userID: number) {
      this.selectedUser = this.project.authority_users.find((user: model.ProjectAuthority) => user.user_id == userID);
      this.userDialog = true;
    },
    getName(v: string | {name: string}) {
      if(!v || typeof v == 'string') return v;
      return v.name;
    }
  }
})
</script>
<style lang="scss" scoped>
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
.loader-frame {
  height: 450px;
  width: 100%;
  min-width: 1560px;
}
.circle-position {
  position: absolute;
  top: 50%;
  left: 50%;
}
.avatar-position {
  position: absolute;
  left: -50px
}
</style>