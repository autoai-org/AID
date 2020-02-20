import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    packages: [],
    logs: [],
    solvers: [],
    isLoading: false,
    config: {},
    alert_info:"",
    alert_title:""
  },
  mutations: {
    setpackages(state, packages) {
      state.packages = packages
    },
    setlogs(state, logs) {
      state.logs = logs
    },
    setsolvers(state, solvers) {
      state.solvers = solvers
    },
    setconfig(state, config) {
      state.config = config
    },
    beginRequests(state) {
      state.isLoading = true
    },
    endRequests(state) {
      state.isLoading = false
    }
  },
  actions: {
  },
  modules: {
  }
})
