// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuetify from 'vuetify'
import VueSocketIO from 'vue-socket.io'
import TreeView from 'vue-json-tree-view'

import 'vuetify/dist/vuetify.min.css'
import '@/assets/styles/main.css'

import i18n from '@/i18n'

Vue.use(new VueSocketIO({
  debug: true,
  connection: 'http://192.168.1.11:10590'
}))
Vue.use(Vuetify)
Vue.use(TreeView)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  components: { App },
  template: '<App/>'
})
