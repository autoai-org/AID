// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import store from '@/store'

const endpoint: string = "http://localhost:10590/"

function _apiRequest(url: string,
    method: "get" | "post" | "patch" | "delete" | "put",
    params: object,
    headers: object,
    onSuccess: Function,
    onError: Function) {
    store.commit('beginRequests')
    return axios.request({
        url,
        method,
        params,
        headers
    }).then((res) => {
        console.log(res.data)
        onSuccess(res.data)
    }).catch((err) => {
        onError(err)
    }).finally(() => {
        store.commit('endRequests')
    })
}

function fetchAllPackages() {
    _apiRequest(endpoint + "packages", "get", {}, {},
        (res: object) => {
            store.commit('setPackages', res)
        },
        (err: object) => {
            console.error(err)
        })
}

export {
    _apiRequest,
    fetchAllPackages
}