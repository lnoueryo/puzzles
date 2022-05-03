import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as lib from '~/modules/store';
import {Tasks} from '~/modules/store/table';
import { RootState } from '~/store'

const initialState = () => {
  return {
    table: lib.preprocessTable(),
    statuses: lib.statuses,
    types: lib.types,
    priorities: lib.priorities,
    authorities: lib.authorities,
    tasks: new Tasks(),
    selectedComment: {id: 0, index: 0, parent: 0},
  }
}

export const state = initialState()

export type TaskState = ReturnType<typeof initialState>

export const getters: GetterTree<TaskState, RootState> = {
  table: state => state.table,
  statuses: state => state.statuses,
  types: state => state.types,
  priorities: state => state.priorities,
  authorities: state => state.authorities,
  selectAssignee: state => state.tasks.tasks.selectAssignee,
  selectField: state => state.tasks.tasks.selectField,
  selectMilestone: state => state.tasks.tasks.selectMilestone,
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
  tasks: (state, tasks: lib.Task[]) => state.tasks.insertTasks(tasks),
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
        lib.storeCondition({cellKey})
        return state.table.cells = lib.resetActive(state.table.cells)
      };
      state.table.cells = lib.resetActive(state.table.cells);
    }
    cell.header.active += 1;
    state.tasks.sortBy(cell);
    const cellKey = {index, active: cell.header.active};
    lib.storeCondition({cellKey});
  },
  selectField: (state, field) => {
    lib.storeCondition({field});
    return state.tasks.selectField(field);
  },
  selectMilestone: (state, milestone) => {
    lib.storeCondition({milestone});
    return state.tasks.selectMilestone(milestone);
  },
  selectAssignee: (state, assignee) => {
    lib.storeCondition({assignee});
    return state.tasks.selectAssignee(assignee);
  },
  selectStatus: (state, status) => {
    lib.storeCondition({status});
    return state.tasks.selectStatus(status);
  },
  selectTask: (state, params) => state.tasks.selectTask(params),
  listIndex: (state, index) => state.tasks.changeListIndex(index),
  pageChange: (state, index) => state.tasks.tasks.pageIndex += index,
  addTask: (state, task) => state.tasks.addTask(task),
  updateTask: (state, task) => state.tasks.updateTask(task),
  selectComment: (state, id) => state.selectedComment = id,
}

export const actions: ActionTree<TaskState, RootState> = {
  reset({commit}) {
    commit('reset');
  },
  async listIndex({commit, getters}, v: number) {
    const index = getters['listNumArr'].findIndex((listNum: number) => listNum === v);
    commit('listIndex', index);
  },
  async getTasks({commit, dispatch}, id: number) {
    console.log('Get Task')
    return new Promise(async(resolve, reject) => {
      try {
        commit('reset')
        const t0 = performance.now();
        const response = await this.$axios.get('/api/task', {
          params: {id: id}
        })
        const t1 = performance.now();
        console.log(`Call to doSomething took ${t1 - t0} milliseconds.`);
        resolve(response);
        commit('tasks', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async updateTask({commit, rootGetters }, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/task/update', form);
        response.data.assignee = rootGetters.project.authority_users.find((authority_user: lib.ProjectAuthority) => {
          return authority_user.user_id == response.data.assignee_id;
        }).user
        commit('updateTask', response.data);
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
        commit('addTask', response.data);
        resolve(response);
        // commit('tasks', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  setCondition({commit}) {
    const itemStr = sessionStorage.getItem(location.host + window.$nuxt.$route.params.id);
    if(!itemStr) return;
    const item = JSON.parse(itemStr);
    if(item?.assignee) commit('selectAssignee', item?.assignee);
    if(item?.field) commit('selectField', item?.field);
    if(item?.status) commit('selectStatus', item.status);
    if(item?.cellKey?.index && item?.cellKey?.active) commit('cellKey', item.cellKey);
  },
}

