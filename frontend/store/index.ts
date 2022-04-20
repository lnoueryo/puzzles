import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as lib from '~/modules/store';

const initialState = () => {
  return {
    user: lib.user,
    pageReady: false,
    projectReady: false,
    snackbar: false,
    snackbarText: '',
    apiHost: process.env.apiHost,
  }
}

export const state = initialState()

export type RootState = ReturnType<typeof initialState>

export const getters: GetterTree<RootState, RootState> = {
  user: state => state.user.user,
  organization: state => state.user.organization,
  projects: state => state.user.projects,
  projectAuthority: state => state.user.selectedProject,
  project: state => state.user.selectedProject.project,
  projectIndex: state => state.user.projectIndex,
  projectSlides: state => state.user.projectSlides,
  pageReady: state => state.pageReady,
  projectReady: state => state.projectReady,
  snackbar: state => state.snackbar,
  snackbarText: state => state.snackbarText,
  mediaUser: state => state.apiHost + '/media/users/',
}

export const mutations: MutationTree<RootState> = {
  reset: (state) => Object.assign(state, initialState()),
  userData: (state, userData: lib.User) => state.user.insertUser(userData),
  selectProject: (state, params) => state.user.selectProject(params),
  pageReady: (state, pageReady) => state.pageReady = pageReady,
  projectReady: (state, projectReady) => state.projectReady = projectReady,
  initUser: state => state.user.init(),
  snackbar: (state, show) => state.snackbar = show,
  snackbarText: (state, text) => state.snackbarText = text,
  updateUser: (state, user) => state.user.updateUser(user),
}

export const actions: ActionTree<RootState, RootState> = {
  resetAll({dispatch}) {
    dispatch('reset');
    dispatch('project/reset');
    dispatch('task/reset');
  },
  reset({commit}) {
    commit('reset');
    commit('initUser');
  },
  async login({}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/login', form);
        resolve(response);
      } catch (error: any) {
        return reject(error.response);
      }
    })
  },
  async session({commit}) {
    console.log('Get Session')
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.get('/api/session');
        // console.log((new Blob([JSON.stringify (response)])).size);
        resolve(response);
        commit('userData', response.data);
        if(!response.data.name) return this.$router.push('/profile/edit');
      } catch (error: any) {
        reject(error.response);
      }
    })
  },
  async logout({dispatch}) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/logout');
        resolve(response)
        dispatch('reset')
      } catch (error: any) {
        reject(error.response)
      }
    })
  },
  async registerUser({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        console.log(form)
        const response = await this.$axios.put('/api/user/update', form);
        resolve(response);
        commit('updateUser', response.data);
      } catch (error: any) {
        reject(error.response);
      }
    })
  },
  async sendEmail({}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.post('/api/invite', form);
        // console.log((new Blob([JSON.stringify (response)])).size);
        resolve(response);
        // commit('userData', response.data);
      } catch (error: any) {
        reject(error.response);
      }
    })
  },
  showSnackbar({commit}, text) {
    commit('snackbarText', text);
    commit('snackbar', true);
  }
}
