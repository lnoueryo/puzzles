<template>
  <v-app dark>
    <v-app-bar class="justify-content: space-between" :clipped-left="clipped" fixed app elevation="0" v-if="pageReady">
      <div class="d-flex justify-space-between" style="width: 100%">
        <div class="d-flex">
          <nuxt-link to="/">
            <v-toolbar-title class="mr-4" v-text="title" />
          </nuxt-link>
          <v-menu>
            <template v-slot:activator="{ on: menu }">
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <div class="d-flex mx-2" style="cursor: pointer" v-bind="attrs" v-on="{ ...on, ...menu }">
                    <div style="display: grid;">
                      <small>プロジェクト名</small>
                      <div v-if="isReadyObj(project)">
                        <strong style="text-indent: 1em;">{{project.name}}</strong>
                      </div>
                      <div v-else>
                        <strong style="text-indent: 1em;">プロジェクトを選択</strong>
                      </div>
                    </div>
                    <v-icon class="mx-2">mdi-chevron-down</v-icon>
                  </div>
                </template>
                <span>プロジェクトの変更</span>
              </v-tooltip>
            </template>
            <v-list>
              <v-list-item-group active-class="border" color="indigo" :value="list">
                <v-list-item v-if="isEmptyArr(projects)" to="/project/create">
                  <v-list-item-title >プロジェクトの作成</v-list-item-title>
                </v-list-item>
                <v-list-item v-for="(project, index) in projects" :key="index" @click="onSelectProject(project)">
                  <v-list-item-title>{{ project.name }}</v-list-item-title>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </v-menu>
        </div>
        <!-- <v-toolbar flat color="transparent">
          <div style="width: 100%;">
            <v-text-field
              v-model="search"
              class="mx-4"
              hide-details
              label="コメントやタスクを検索"
              prepend-inner-icon="mdi-magnify"
              solo-inverted
              clearable
              color="amber darken-3"
            ></v-text-field>
          </div>
        </v-toolbar> -->
        <div class="d-flex mx-4 align-center">
          <!-- <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn class="mx-2" icon v-bind="attrs" v-on="on">
                <v-badge color="green" :content="messages" :value="messages" overlap>
                <v-icon>mdi-bell</v-icon>
                </v-badge>
              </v-btn>
            </template>
            <span>通知の確認</span>
          </v-tooltip> -->
          <v-menu>
            <template v-slot:activator="{ on: menu }">
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn class="mx-2" icon v-bind="attrs" v-on="{ ...on, ...menu }">
                    <v-avatar size="40px">
                      <img alt="Avatar" :src="$config.mediaURL + '/users/' + user.image">
                    </v-avatar>
                  </v-btn>
                </template>
                <span>プロフィール設定</span>
              </v-tooltip>
            </template>
            <v-list>
              <v-list-item v-for="(config, index) in configs" :key="index" :to="config.url">
                <v-list-item-title>{{ config.name }}</v-list-item-title>
              </v-list-item>
              <v-list-item @click="logout">
                <v-list-item-title>ログアウト</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
          <div class="mx-3" style="display: grid;">
            <small>名前</small>
            <strong style="text-indent: 1em;">{{user.name}}</strong>
          </div>
          <div class="mx-3" style="display: grid;">
            <small>組織名</small>
            <strong style="text-indent: 1em;">{{ authority.organization.name }}</strong>
          </div>
          <v-menu left bottom v-if="authority.auth_id == 1">
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on">
                <v-icon>mdi-dots-vertical</v-icon>
              </v-btn>
            </template>

            <v-list>
              <v-list-item v-for="(item, i) in administer" :key="i" :to="item.url">
                <v-list-item-title>{{ item.name }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </div>
      </div>
    </v-app-bar>
    <!-- <div style="background-color: #272727;width: 100%;padding-top: 64px">
      <v-breadcrumbs :items="breadCrumbs">
        <template v-slot:divider>
          <v-icon>mdi-chevron-right</v-icon>
        </template>
      </v-breadcrumbs>
    </div> -->
    <v-main>
        <Nuxt />
    </v-main>
    <v-footer :absolute="!fixed" app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
    <snack-bar></snack-bar>
  </v-app>
</template>

<script>
import {isEmptyArr, isReadyObj, isEmptyObj, checkStatus} from '~/modules/utils'
export default {
  name: 'DefaultLayout',
  data () {
    return {
      listNum: -1,
      // isSession: false,
      clipped: true,
      drawer: false,
      fixed: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'ダッシュボード',
          to: '/'
        },
        {
          icon: 'mdi-plus',
          title: 'タスク追加',
          to: '/task'
        }
      ],
      miniVariant: false,
      title: 'puzzles',
      search: '',
      messages: 0,
      configs: [
        {
          name: 'プロフィール',
          url: '/profile'
        },
      ],
      administer: [
        {
          name: 'プロジェクトの作成',
          url: '/project/create'
        },
        {
          name: '組織に招待する',
          url: '/organization/create'
        },
      ],
    }
  },
  computed: {
    isReadyObj,
    isEmptyArr,
    isEmptyObj,
    checkStatus,
    user() {
      return this.$store.getters['user'];
    },
    authority() {
      return this.$store.getters['organization'];
    },
    project() {
      return this.$store.getters['project'];
    },
    projects() {
      return this.$store.getters['projects'];
    },
    list() {
      return this.$store.getters['projectIndex'];
    },
    pageReady() {
      return this.$store.getters['pageReady'];
    },
    breadCrumbs() {
      return this.$store.getters['breadCrumbs'];
    }
  },
  // async beforeRouteEnter (to, from, next) {
  //   console.log(this.isReadyObj(this.$store.getters.user))
  //   if(this.isReadyObj(this.$store.getters.user)) return next();
  //   let response
  //   try {
  //     response = await this.$store.dispatch('session');
  //   } catch (error) {
  //     response = error;
  //   } finally {
  //     return status(response.status, () => {
  //       this.$store.commit('pageReady', true);
  //       return next();
  //     },
  //     () => next('/login'));
  //   }
  // },
  async created() {
    let path = ''
    const paths = [{text: 'Home', disabled: false, href: '/'}]
    this.$route.matched[0].path.split('/').forEach((devidedPath) => {
      if(devidedPath == '') return;
      console.log(path)
      if(devidedPath[0] == ':') {
        const params = this.$route.params[devidedPath.replace(':', '')]
        path += '/' + params
        paths[paths.length -1].href = path;
        return;
      }
      path += '/' + devidedPath
      paths.push({text: devidedPath, disabled: false, href: path})
    })
    console.log(paths)
    // window.onload = () => {
    //   this.$store.getters['projectCreateSock'].onmessage = (event) => {
    //     const data = JSON.parse(event.data);
    //     if(data.Name == 'Rio') {
    //       console.log(data)
    //     }
    //     this.$store.dispatch('getProjectData', data)
    //   };
    // };
    // let response
    // try {
    //   response = await this.$store.dispatch('session');
    // } catch (error) {
    //   response = error;
    // } finally {
    //   this.checkStatus(response.status, () => {this.isSession = true});
    // }
  },
  methods: {
    onSelectProject(project) {
      this.$router.push({name: 'project-id', params: {id: project.id}});
    },
    async logout() {
      const OK = 200;
      try {
        const response = await this.$store.dispatch('logout');
        if(response.status == OK) {
          this.isSession = false;
          this.$store.dispatch('reset');
          this.$router.push('/login');
        }
      } catch (error) {
        console.log(error)
      }
    }
  }
}
</script>
<style lang="scss">
.v-application a {
  text-decoration: none;
  color: white!important;
}
/* Change the white to any color */
input:-webkit-autofill, input:-webkit-autofill:hover, input:-webkit-autofill:focus, input:-webkit-autofill:active{
    -webkit-box-shadow: 0 0 0 30px rgba(255, 255, 255, 0) inset !important;
}
input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
    transition: background-color 5000s ease-in-out 0s;
}
input:-webkit-autofill{
    -webkit-text-fill-color: rgb(255, 255, 255) !important;
}
</style>