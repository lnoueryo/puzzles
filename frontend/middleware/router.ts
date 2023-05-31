import { Middleware, Context } from '@nuxt/types'
// import { storeCondition } from '~/modules/model';
import { checkStatus, isReadyObj, isEmptyObj } from '~/modules/utils'
const status = checkStatus();
const readyObj = isReadyObj();
const emptyObj = isEmptyObj();
const allowedPath = new Set(['/login', '/expiry', '/success', '/error/bad-connection', '/data/csv'])
let projectID: string;
let overlap = false;
const router: Middleware = async({store, route, redirect}) => {
  if(emptyObj(store.getters.user)) store.commit('pageReady', false);
  if(allowedPath.has(route.path)) return;
  getSession(store, redirect);
  breadCrumbs(store, route)
  selectProject(store, route);
  selectUser(store, route);
  if(Object.keys(route.params).length === 0) return;
  await getTask(store, route, redirect);
}

const getSession = async(store: any, redirect: any) => {
  if(readyObj(store.getters.user)) return;
  let response
  try {
    response = await store.dispatch('session');
  } catch (error) {
    response = error;
  } finally {
    if(!response || 'status' in response === false) return redirect('/error/bad-connection')
    return status(response.status, () => {
      store.commit('pageReady', true);
    });
  }
}

const breadCrumbs = (store: any, route: any) => {
  let path = ''
  const paths = [{text: 'Home', disabled: false, href: '/'}]
  route.matched[0].path.split('/').forEach((devidedPath: string) => {
    if(devidedPath == '') return;
    if(devidedPath[0] == ':') {
      const params = route.params[devidedPath.replace(':', '')]
      path += '/' + params
      paths[paths.length -1].href = path;
      return;
    }
    path += '/' + devidedPath
    paths.push({text: devidedPath, disabled: false, href: path})
  })
  paths[paths.length -1].disabled = true;
  store.commit('breadCrumbs', paths);
}

const selectProject = (store: any, route: any) => {
  let timer = setInterval(() => {
    if(emptyObj(store.getters.user)) return;
    clearInterval(timer);
    store.commit('pageReady', true);
    store.commit('selectProject', route.params);
    store.commit('projectReady', true);
  }, 100)
}

const selectUser = (store: any, route: any) => {
  let timer = setInterval(() => {
    if(emptyObj(store.getters.user)) return;
    clearInterval(timer);
    store.commit('selectUser', route.params);
    store.commit('pageReady', true);
  }, 100)
}

const getTask = async(store: any, route: any, redirect: any) => {
  if(route.params.id == projectID) return setCondition(store);
  if(overlap) {
    overlap = false;
    return;
  }
  overlap = true;
  let timer = setInterval(async() => {
    if(emptyObj(store.getters.project)) return;
    clearInterval(timer);
    let response;
    try {
      response = await store.dispatch('task/getTasks', store.getters.project.id);
      projectID = store.getters.project.id;
      setCondition(store);
    } catch (error: any) {
      response = error.response
    } finally {
      overlap = false;
      if('status' in response === false) return redirect('/error/bad-connection')
      status(response.status, () => {}, () => {
        alert('エラーです。');
      })
    }
  }, 100)
}

const setCondition = (store: any) => {
  store.dispatch('task/setCondition');
}

export default router