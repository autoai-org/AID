import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    packages: [],
    logs:[],
    isLoading: false,
  },
  mutations: {
    setpackages(state, packages) {
      state.packages = packages
    },
    setlogs(state, logs) {
      state.logs = logs
    },
    beginRequests(state) {
      state.isLoading = true
    },
    endRequests (state) {
      state.isLoading = false
    }
  },
  actions: {
  },
  modules: {
  }
})
