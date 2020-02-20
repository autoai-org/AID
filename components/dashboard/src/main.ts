import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import { alert } from '@/middlewares/store.mdw'

Vue.config.productionTip = false
Vue.prototype.$alert = alert

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
