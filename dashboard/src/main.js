// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import '@/components/auth'
import Vuetify from 'vuetify'
import TreeView from 'vue-json-tree-view'
import VueTour from 'vue-tour'
import colors from 'vuetify/es5/util/colors'
import 'vuetify/dist/vuetify.min.css'
import '@/assets/styles/main.css'
import 'vue-tour/dist/vue-tour.css'

import zhHans from 'vuetify/es5/locale/zh-Hans'
import en from 'vuetify/es5/locale/en'
import de from 'vuetify/es5/locale/de'

import i18n from '@/i18n'

Vue.use(Vuetify)
Vue.use(TreeView)
Vue.use(VueTour)

Vue.config.productionTip = false

Vue.use(Vuetify, {
  lang: {
    locales: { en, zhHans, de },
    current: 'zhHans'
  },
  theme: {
    primary: colors.indigo.darken1,
    secondary: colors.red.lighten4,
    accent: colors.indigo.base
  }
})

/* eslint-disable no-new */
new Vue({
  i18n,
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
