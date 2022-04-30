import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as lib from '~/modules/store';
import { RootState } from '~/store'

const initialState = () => {
  return {
    user: lib.user,
    name: '' as string
  }
}

export const state = initialState()

export type ProjectState = ReturnType<typeof initialState>

export const getters: GetterTree<ProjectState, RootState> = {
  // project: state => state.project.main,
}

export const mutations: MutationTree<ProjectState> = {
  reset: (state) => Object.assign(state, initialState()),
  createProject: (state, project: lib.Project) => state.user.createProject(project),
  updateProject: (state, project: lib.Project) => state.user.updateProject(project),
}

export const actions: ActionTree<ProjectState, RootState> = {
  reset({commit}) {
    commit('reset');
  },
  async selectProject({commit}, id: number) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.get('/api/project', {
          params: {id: id}
        })
        resolve(response);
        commit('task/tasks', response.data.tasks, { root: true });
        // commit('project', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async selectEditProject({commit}, id: number) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.get('/api/project/edit', {
          params: {id: id}
        });
        resolve(response);
        commit('project', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async createProject({commit, rootState}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/project/create', form);
        commit('createProject', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async updateProject({commit, dispatch}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/project/update', form);
        console.log(form)
        if(form.field_delete || form.milestone_delete) dispatch('task/getTasks', window.$nuxt.$route.params.id, {root: true});
        commit('updateProject', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
}

