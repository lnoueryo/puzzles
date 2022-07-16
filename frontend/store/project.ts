import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as model from '~/modules/model';
import { RootState } from '~/store'

const initialState = () => {
  return {
    user: model.user,
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
  createProject: (state, project: model.Project) => state.user.createProject(project),
  updateProject: (state, project: model.Project) => state.user.updateProject(project),
  createProjectAuthority: (state, projectAuthority: model.ProjectAuthority) => state.user.createProjectAuthority(projectAuthority),
  updateProjectAuthority: (state, projectAuthority: model.ProjectAuthority) => state.user.updateProjectAuthority(projectAuthority),
  deleteProjectAuthority: (state, projectAuthority: model.ProjectAuthority) => state.user.deleteProjectAuthority(projectAuthority),
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
  async createProject({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/project/create', form);
        commit('userData', response.data, {root: true});
        // commit('createProject', response.data);
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
        if(form.field_delete || form.milestone_delete) dispatch('task/getTasks', window.$nuxt.$route.params.id, {root: true});
        commit('updateProject', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async createProjectAuthority({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/project-authority/create', form);
        commit('createProjectAuthority', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async updateProjectAuthority({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/project-authority/update', form);
        commit('updateProjectAuthority', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async deleteProjectAuthority({commit}, projectAuthority) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.delete('/api/project-authority/delete', {params: {id: projectAuthority.id}});
        commit('deleteProjectAuthority', projectAuthority);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
}

