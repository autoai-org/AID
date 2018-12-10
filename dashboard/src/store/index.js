import Vue from 'vue'
import Vuex from 'vuex'
import attachCognitoModule from '@vuetify/vuex-cognito-module'
Vue.use(Vuex)

const set = property => (store, payload) => (store[property] = payload)

const store = new Vuex.Store({
  state: {
    email: '',
    isLoading: false,
    isReady: false,
    password: '',
    snackbar: {}
  },
  mutations: {
    setEmail: set('email'),
    setIsReady: set('isReady'),
    setIsLoading: set('isLoading'),
    setPassword: set('password'),
    setUser: set('user'),
    setSnackbar: set('snackbar')
  }
})

attachCognitoModule(store, {
  userPoolId: 'us-east-1_IYJ3FvCKZ',
  identityPoolId: 'us-east-1:cfb20347-6f2b-44ce-bac9-29ca3af55274',
  userPoolWebClientId: '105hsloscm5ucmmgiokdmtg0dh',
  region: 'us-east-1'
}, 'cognito')

store.dispatch('cognito/fetchSession')
  .finally(() => store.commit('setIsReady', true))

export default store
