import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as model from '~/modules/model';
import { Table } from '~/modules/table';
import { RootState } from '~/store'

const initialState = () => {
  return {
    table: new Table(),
    statuses: model.statuses,
    types: model.types,
    priorities: model.priorities,
    authorities: model.authorities,
    tasks: new model.Tasks(),
    selectedComment: {id: 0, index: 0, parent: 0},
  }
}

export const state = initialState()

export type TaskState = ReturnType<typeof initialState>

export const getters: GetterTree<TaskState, RootState> = {
  table: state => state.table.items,
  statuses: state => state.statuses,
  types: state => state.types,
  priorities: state => state.priorities,
  authorities: state => state.authorities,
  selectAssignee: state => state.tasks.tasks.selectAssignee,
  selectField: state => state.tasks.tasks.selectField,
  selectMilestone: state => state.tasks.tasks.selectMilestone,
  selectVersion: state => state.tasks.tasks.selectVersion,
  selectStatus: state => state.tasks.tasks.selectStatus,
  listNumArr: state => state.tasks.tasks.listNumArr,
  selectList: state => state.tasks.tasks.listNumArr[state.tasks.tasks.listIndex],
  pageIndex: state => state.tasks.tasks.pageIndex,
  allTasks: state => state.tasks.tasks.all,
  tasks: state => state.tasks.main,
  totalTasks: state => state.tasks.filterTasks.length,
  totalPageNum: state => state.tasks.totalPageNum,
  currentDisplayedTasksNum: state => state.tasks.currentDisplayedTasksNum,
  task: state => state.tasks.selectedTask,
  selectedComment: state => state.selectedComment
}

export const mutations: MutationTree<TaskState> = {
  reset: (state) => Object.assign(state, initialState()),
  tasks: (state, tasks: model.Task[]) => state.tasks.insertTasks(tasks),
  cellKey: (state, cellKey) => {
    const cell = state.table.cells[cellKey.index];
    cell.header.active = cellKey.active;
    state.tasks.sortBy(cell);
  },
  sortTask: (state, index) => {
    const cell = state.table.cells[index];
    if(cell.header.active != 1) {
      if(cell.header.active === 2) {
        state.tasks.resetSort();
        const cellKey = {index: 0, active: 0}
        state.tasks.storeCondition({cellKey})
        return state.table.cells = state.table.resetActive(state.table.cells)
      };
      state.table.cells = state.table.resetActive(state.table.cells);
    }
    cell.header.active += 1;
    state.tasks.sortBy(cell);
    const cellKey = {index, active: cell.header.active};
    state.tasks.storeCondition({cellKey});
  },
  selectField: (state, field) => {
    state.tasks.storeCondition({field});
    return state.tasks.selectField(field);
  },
  selectMilestone: (state, milestone) => {
    state.tasks.storeCondition({milestone});
    return state.tasks.selectMilestone(milestone);
  },
  selectVersion: (state, version) => {
    state.tasks.storeCondition({version});
    return state.tasks.selectVersion(version);
  },
  selectAssignee: (state, assignee) => {
    state.tasks.storeCondition({assignee});
    return state.tasks.selectAssignee(assignee);
  },
  selectStatus: (state, status) => {
    state.tasks.storeCondition({status});
    return state.tasks.selectStatus(status);
  },
  // selectTask: (state, params) => state.tasks.selectTask(params),
  listIndex: (state, index) => state.tasks.changeListIndex(index),
  pageChange: (state, index) => state.tasks.tasks.pageIndex += index,
  selectComment: (state, id) => state.selectedComment = id,
}

export const actions: ActionTree<TaskState, RootState> = {
  reset({commit}) {
    commit('reset');
  },
  /** 選択されたタスクの最大表示数より、最大表示数が格納されている配列のインデックスを取得 */
  async listIndex({commit, getters}, v: number) {
    const index = getters['listNumArr'].findIndex((listNum: number) => listNum === v);
    commit('listIndex', index);
  },
  async getTasks({commit, dispatch}, id: number) {
    commit('reset')
    console.log('Get Task')
    return new Promise(async(resolve, reject) => {
      try {
        const t0 = performance.now();
        const response = await this.$axios.get('/api/task', {
          params: {
            id: id,
            page: 0
          }
        })
        const t1 = performance.now();
        console.log(`Call to doSomething took ${t1 - t0} milliseconds.`);
        const size = new TextEncoder().encode(JSON.stringify(response)).length
        const kiloBytes = size / 1024;
        const megaBytes = kiloBytes / 1024;
        resolve(response);
        commit('tasks', response.data.tasks);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async updateTask({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/task/update', form);
        commit('reset')
        commit('tasks', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async createTask({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/task/create', form);
        commit('reset')
        commit('tasks', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },

  /** セッションストレージからタスクの絞り込み条件を取得 */
  setCondition({commit}) {
    const itemStr = sessionStorage.getItem(location.host + window.$nuxt.$route.params.id);
    if(!itemStr) return;
    const item = JSON.parse(itemStr);
    if(item?.assignee) commit('selectAssignee', item?.assignee);
    if(item?.field) commit('selectField', item?.field);
    if(item?.milestone) commit('selectMilestone', item?.milestone);
    if(item?.version) commit('selectVersion', item?.version);
    if(item?.status) commit('selectStatus', item.status);
    if(item?.cellKey?.index && item?.cellKey?.active) commit('cellKey', item.cellKey);
  },
}

