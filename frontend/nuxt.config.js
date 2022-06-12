import colors from 'vuetify/es5/util/colors'

export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,
  router: {
    middleware: ['router'],
    // extendRoutes(routes, resolve) {
    //   routes.push({
    //     name: 'notFound',
    //     path: '*',
    //     component: resolve(__dirname, 'pages/error/bad-connection.vue')
    //   })
    // }
  },
  // Target: https://go.nuxtjs.dev/config-target
  target: 'static',
  publicRuntimeConfig: {
    apiHost: process.env.NODE_ENV === 'production' ? process.env.API_URL : 'http://localhost:8080',
    mediaURL: process.env.MEDIA_URL,
    // mediaURL: process.env.NODE_ENV === 'production' ? process.env.MEDIA_URL : 'http://localhost:8080/media',
  },
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: '%s - frontend',
    title: 'frontend',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    "~assets/scss/common.scss"
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    [
      '@nuxtjs/firebase',
      {
        config: {
          apiKey: 'AIzaSyD_55QAis9q0SYUIeq0IywW9l5RZfEel4Q',
          authDomain: 'puzzles-345814.firebaseapp.com',
          projectId: 'puzzles-345814',
          storageBucket: 'puzzles-345814.appspot.com',
          messagingSenderId: '1079713727627',
          appId: '1:1079713727627:web:22d158a499f22ec2ca808e',
          measurementId: 'G-KMGMB72GE5'
        },
        services: {
          auth: false // Just as example. Can be any other service.
        }
      }
    ],
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    // https://go.nuxtjs.dev/pwa
    '@nuxtjs/pwa',
    '@nuxtjs/markdownit',
  ],
  markdownit: {
    injected: true,
    breaks: true, // 改行コードに変換する
    html: true, // HTML タグを有効にする
    linkify: true, // URLに似たテキストをリンクに自動変換する
    typography: true, // 言語に依存しないきれいな 置換 + 引用符 を有効にします。
  },

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    prefix: process.env.NODE_ENV === 'production' ? process.env.BASE_URL : 'http://localhost:3000',
    baseURL: process.env.NODE_ENV === 'production' ? process.env.BASE_URL : 'http://localhost:3000', // 正しい
    browserBaseURL: process.env.NODE_ENV === 'production' ? process.env.API_URL : 'http://localhost:8080', // 正しい
    credentials: true
  },

  // PWA module configuration: https://go.nuxtjs.dev/pwa
  pwa: {
    manifest: {
      lang: 'en'
    }
  },

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        },
      }
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
