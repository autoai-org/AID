// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import store from '@/store'
import dayjs from 'dayjs'
import { Package, Log } from '@/entities'
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
        (res: Array<Package>) => {
            res = res.map(function (each) {
                each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                return each
            })
            store.commit('setPackages', res)
        },
        (err: object) => {
            console.error(err)
        })
}

function buildImage(packageName: string, solverName: string) {
    _apiRequest(endpoint + "packages/" + packageName + "/solvers" + solverName + "/images", "put", {}, {},
        (res: Array<Package>) => {
            store.commit('setPackages', res)
        },
        (err: object) => {
            console.error(err)
        })
}

function fetchAllLogs() {
    _apiRequest(endpoint + "logs", "get", {}, {},
        (res: Array<Log>) => {
            res = res.map(function (each) {
                each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                return each
            })
            store.commit('setLogs', res)
        },
        (err: object) => {
            console.error(err)
        })
}

function fetchLog(id: number) {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "logs/" + id, "get", {}, {},
            (res: Log) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })

    })
}

export {
    _apiRequest,
    fetchAllPackages,
    buildImage,
    fetchAllLogs,
    fetchLog
}