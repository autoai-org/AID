import axios from 'axios'
import store from '@/store'

function _get (endpoint) {
  store.state.isLoading = true
  return new Promise((resolve, reject) => {
    axios.get(endpoint).then(function (res) {
      resolve(res)
    }).catch(function (err) {
      reject(err)
    }).finally(function () {
      store.state.isLoading = false
    })
  })
}

function _post (endpoint, data) {
  store.state.isLoading = true
  return new Promise((resolve, reject) => {
    axios.post(endpoint, data).then(function (res) {
      resolve(res)
    }).catch(function (err) {
      reject(err)
    }).finally(function () {
      store.state.isLoading = false
    })
  })
}

export {
  _get,
  _post
}
