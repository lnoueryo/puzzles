import { Middleware } from '@nuxt/types'
import { storeCondition } from '~/modules/store';
import { checkStatus, isReadyObj, isEmptyObj } from '~/modules/utils'
const status = checkStatus();
const readyObj = isReadyObj();
const emptyObj = isEmptyObj();
const allowedPath = ['/login', '/expiry', '/success', '/bad-connection']
let projectID: string;
const router: Middleware = async({store, route}) => {
  if(allowedPath.includes(route.path)) return;
  store.commit('pageReady', false);
  getSession(store);
  selectProject(store, route);
  if (Object.keys(route.params).length === 0) return;
  getTask(store, route);
  selectTask(store, route);
}

const getSession = async(store: any) => {
  if(readyObj(store.getters.user)) return;
  let response
  try {
    response = await store.dispatch('session');
  } catch (error) {
    response = error;
  } finally {
    if('status' in response === false) return window.$nuxt.$router.push('/bad-connection')
    return status(response.status, () => {
      store.commit('pageReady', true);
    });
  }
}

const selectProject = (store: any, route: any) => {
  let timer = setInterval(() => {
    if(emptyObj(store.getters.user)) return;
    clearInterval(timer);
    store.commit('selectProject', route.params);
    store.commit('pageReady', true);
    store.commit('projectReady', true);
  }, 100)
}

const getTask = (store: any, route: any) => {
  if(route.params.id == projectID) return setCondition(store);
  let timer = setInterval(async() => {
    if(emptyObj(store.getters.project)) return;
    clearInterval(timer);
    let response;
    try {
      response = await store.dispatch('task/getTasks', route.params.id);
      projectID = store.getters.project.id;
      setCondition(store);
    } catch (error: any) {
      response = error.response
    } finally {
      if('status' in response === false) return window.$nuxt.$router.push('/bad-connection')
      status(response.status, () => {}, () => {
        alert('エラーです。');
      })
    }
  }, 100)
}

const setCondition = (store: any) => {
  store.dispatch('task/setCondition');
}

const selectTask = (store: any, route: any) => {
  store.commit('task/selectTask', {});
  if('key' in route.params === false) return;
  let timer = setInterval(() => {
    if(store.getters['task/allTasks'].length === 0) return;
    clearInterval(timer);
    store.commit('task/selectTask', route.params);
  }, 100)
}

export default router