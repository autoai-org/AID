// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import '@/components/auth'
import Vuetify from 'vuetify'
import VueSocketIO from 'vue-socket.io'
import TreeView from 'vue-json-tree-view'
import VueTour from 'vue-tour'
import colors from 'vuetify/es5/util/colors'
import { ConfigService } from '@/services/config'

import 'vuetify/dist/vuetify.min.css'
import '@/assets/styles/main.css'
import 'vue-tour/dist/vue-tour.css'

import i18n from '@/i18n'
const configService = new ConfigService()

Vue.use(new VueSocketIO({
  debug: false,
  connection: configService.endpoint
}))
Vue.use(Vuetify)
Vue.use(TreeView)
Vue.use(VueTour)

Vue.config.productionTip = false

Vue.use(Vuetify, {
  theme: {
    primary: colors.indigo.darken1,
    secondary: colors.red.lighten4,
    accent: colors.indigo.base
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  components: { App },
  template: '<App/>'
})
