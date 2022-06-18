import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as model from '~/modules/model';

const initialState = () => {
  return {
    user: model.user,
    pageReady: false,
    projectReady: false,
    snackbar: false,
    snackbarText: '',
    apiHost: process.env.apiHost,
    breadCrumbs: {text: '', disabled: false, href: ''}
  }
}

export const state = initialState()

export type RootState = ReturnType<typeof initialState>

export const getters: GetterTree<RootState, RootState> = {
  user: state => state.user.user,
  organizationAuthority: state => state.user.organizationAuthority,
  projects: state => state.user.projects,
  project: state => state.user.selectedProject,
  selectedUser: state => state.user.selectedUser,
  projectAuthority: state => state.user.projectAuthority,
  projectIndex: state => state.user.projectIndex,
  projectSlides: state => state.user.projectSlides,
  pageReady: state => state.pageReady,
  projectReady: state => state.projectReady,
  snackbar: state => state.snackbar,
  snackbarText: state => state.snackbarText,
  mediaUser: state => state.apiHost + '/media/users/',
  breadCrumbs: state => state.breadCrumbs,
}

export const mutations: MutationTree<RootState> = {
  reset: (state) => Object.assign(state, initialState()),
  insertUserData: (state, userData: model.MainUserInfo) => state.user.insertUser(userData),
  selectProject: (state, params) => state.user.selectProject(params),
  selectUser: (state, params) => state.user.selectUser(params),
  pageReady: (state, pageReady) => state.pageReady = pageReady,
  projectReady: (state, projectReady) => state.projectReady = projectReady,
  resetUser: state => state.user.reset(),
  snackbar: (state, show) => state.snackbar = show,
  snackbarText: (state, text) => state.snackbarText = text,
  updateUser: (state, user) => state.user.updateUser(user),
  breadCrumbs: (state, breadCrumbs) => state.breadCrumbs = breadCrumbs
}

export const actions: ActionTree<RootState, RootState> = {
  resetAll({dispatch}) {
    dispatch('reset');
    dispatch('project/reset');
    dispatch('task/reset');
  },
  reset({commit}) {
    commit('reset');
    commit('resetUser');
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
    console.info('Get Session')
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.get('/api/session');
        resolve(response);
        commit('insertUserData', response.data);
        if(!response.data.user.name) return this.$router.push('/profile/edit');
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
        const response = await this.$axios.put('/api/user/update', form);
        resolve(response);
        commit('insertUserData', response.data);
      } catch (error: any) {
        reject(error.response);
      }
    })
  },
  /** 新たにユーザー登録を行うユーザーにメールを送信 */
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
  updateOrganization({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/organization/update', form);
        // console.log((new Blob([JSON.stringify (response)])).size);
        resolve(response);
        commit('insertUserData', response.data);
      } catch (error: any) {
        reject(error.response);
      }
    })
  },
  updateOrganizationAuthority({commit}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await this.$axios.put('/api/organization-authority/update', form);
        // console.log((new Blob([JSON.stringify (response)])).size);
        resolve(response);
        commit('insertUserData', response.data);
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
