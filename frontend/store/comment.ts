import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as lib from '~/modules/store';
import { Comments } from '~/modules/store/comment';
import { RootState } from '~/store'

const initialState = () => {
  return {
    comments: new Comments(),
    selectedComment: {id: 0, index: 0, parent: 0},
  }
}

export const state = initialState()

export type TaskState = ReturnType<typeof initialState>

export const getters: GetterTree<TaskState, RootState> = {
  comments: state => state.comments.comments.all,
  selectedComment: state => state.selectedComment
}

export const mutations: MutationTree<TaskState> = {
  reset: (state) => Object.assign(state, initialState()),
  comments: (state, comments: lib.Comment[]) => state.comments.insertTasks(comments),
//   selectTask: (state, params) => state.comments.selectTask(params),
//   addTask: (state, task) => state.comments.addTask(task),
//   updateTask: (state, task) => state.comments.updateTask(task),
  selectComment: (state, id) => state.selectedComment = id,
}

export const actions: ActionTree<TaskState, RootState> = {
  reset({commit}) {
    commit('reset');
  },
  async getComments({commit, dispatch}, id: number) {
    console.log('Get Comment')
    return new Promise(async(resolve, reject) => {
      try {
        commit('reset')
        const response = await this.$axios.get('/api/comment/show', {
          params: {id: id}
        })
        resolve(response);
        commit('comments', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async createComment({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/comment/create', form);
        commit('updateTask', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async updateComment({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/comment/update', form);
        // commit('updateTask', response.data);
        resolve(response);
        // commit('tasks', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
}

