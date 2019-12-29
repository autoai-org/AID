import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    packages: [],
    isLoading: false,
  },
  mutations: {
    setPackages(state, packages) {
      state.packages = packages
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
