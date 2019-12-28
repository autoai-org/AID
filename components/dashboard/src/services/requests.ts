import axios from 'axios'

function _get(endpoint: string) {
    return new Promise((resolve, reject) => {
        axios.get(endpoint).then(function (res) {
            resolve(res)
        }).catch(function (err) {
            reject(err)
        }).finally(function () {
        })
    })
}

function _post(endpoint: string, data: object) {
    return new Promise((resolve, reject) => {
        axios.post(endpoint, data).then(function (res) {
            resolve(res)
        }).catch(function (err) {
            reject(err)
        }).finally(function () {
        })
    })
}

export {
    _get,
    _post
}