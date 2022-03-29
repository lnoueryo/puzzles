<template>
  <v-app dark>
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-toolbar-title class="mr-4" v-text="title" />
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <div class="d-flex mx-2" v-bind="attrs" v-on="on">
              <div style="display: grid;">
                <small>プロジェクト名</small>
                <strong style="text-indent: 1em;">心の筋トレプロジェクト</strong>
              </div>
              <v-icon class="mx-2">mdi-chevron-down</v-icon>
            </div>
          </template>
          <span>プロジェクトの変更</span>
        </v-tooltip>
      <v-toolbar
        flat
        color="transparent"
      >
      <div style="max-width: 700px;width: 100%;">
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
      </v-toolbar>
      <div class="d-flex mx-4 align-center">
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              class="mx-2"
              icon
              v-bind="attrs"
              v-on="on"
            >
              <v-badge
                color="green"
                :content="messages"
                :value="messages"
                overlap
              >
              <v-icon>mdi-bell</v-icon>
              </v-badge>
            </v-btn>
          </template>
          <span>通知の確認</span>
        </v-tooltip>
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              class="mx-2"
              icon
              v-bind="attrs"
              v-on="on"
            >
              <v-avatar
                size="40px"
              >
                <img
                  alt="Avatar"
                  :src="user.image"
                >
              </v-avatar>
            </v-btn>
          </template>
          <span>プロフィール設定</span>
        </v-tooltip>
        <div class="mx-3" style="display: grid;">
          <small>名前</small>
          <strong style="text-indent: 1em;">{{name}}</strong>
        </div>
        <div class="mx-3" style="display: grid;">
          <small>組織名</small>
          <strong style="text-indent: 1em;">＋Base</strong>
        </div>
      </div>
    </v-app-bar>
    <v-main>
      <Nuxt />
    </v-main>
    <powered-by></powered-by>
  </v-app>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data () {
    return {
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
      title: 'backend',
      search: '',
      messages: 0,
    }
  },
  computed: {
    user() {
      return this.$store.getters['user']
    },
    name() {
      return this.user.familyName + ' ' + this.user.firstName;
    }
  },
}
</script>
<style lang="scss">
a {
  text-decoration: none!important;
}
</style>