import axios from 'axios'

function _get (endpoint) {
  return new Promise((resolve, reject) => {
    axios.get(endpoint).then(function (res) {
      resolve(res)
    }).catch(function (err) {
      reject(err)
    })
  })
}

function _post (endpoint, data) {
  return new Promise((resolve, reject) => {
    axios.post(endpoint, data).then(function (res) {
      resolve(res)
    }).catch(function (err) {
      reject(err)
    })
  })
}

export {
  _get,
  _post
}
