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
    data: object,
    headers: object,
    onSuccess: Function,
    onError: Function) {
    store.commit('beginRequests')
    return axios.request({
        url,
        method,
        data,
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

function buildImage(packageName: string, solverName: string) {
    _apiRequest(endpoint + "packages/" + packageName + "/solvers" + solverName + "/images", "put", {}, {},
        (res: Array<Package>) => {
            store.commit('setPackages', res)
        },
        (err: object) => {
            console.error(err)
        })
}

function installPackage(packageIdentifier: string) {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "packages", "post", {
            'remote_url': packageIdentifier
        }, {},
            (res: Log) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
    })
}

function fetchAllObjects(objectName: string) {
    _apiRequest(endpoint + objectName, "get", {}, {},
        (res: Array<Log>) => {
            res = res.map(function (each) {
                each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                return each
            })
            store.commit('set' + objectName, res)
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
    buildImage,
    fetchLog,
    fetchAllObjects,
    installPackage
}