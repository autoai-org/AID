import axios from 'axios'

function _request (method, url, data) {
  return axios.request({
    method: method,
    url: url,
    data: data
  })
}

export {
  _request
}
