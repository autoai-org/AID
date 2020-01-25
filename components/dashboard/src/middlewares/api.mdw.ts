// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import store from '@/store'
import dayjs from 'dayjs'
import { Log } from '@/entities'
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
        onSuccess(res.data)
    }).catch((err) => {
        onError(err)
    }).finally(() => {
        store.commit('endRequests')
    })
}

function buildImage(vendorName: string, packageName: string, solverName: string) {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "packages/" + vendorName + "/" + packageName + "/" + solverName + "/images", "put", {}, {},
            (res: any) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
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
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + objectName, "get", {}, {},
            (res: Array<Log>) => {
                res = res.map(function (each) {
                    each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                    return each
                })
                resolve(res)
                store.commit('set' + objectName, res)
            },
            (err: object) => {
                reject(err)
                console.error(err)
            })
    })
}

function fetchLog(id: string) {
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

function deleteLog(id: string) {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "logs/" + id, "delete", {}, {},
            (res: Log) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
    })
}

function fetchConfig() {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "configs", "get", {}, {},
            (res: object) => {
                resolve(res)
                store.commit('setconfig', res)
            },
            (err: object) => {
                reject(err)
                console.error(err)
            })
    })
}

function updateConfig(config: object) {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "configs", "post", config, {},
            (res: object) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
    })
}

function fetchMeta(vendorName: string, packageName: string) {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "packages/" + vendorName + "/" + packageName +"/meta", "get", {}, {},
            (res: object) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
                console.error(err)
            })
    })
}

function fetchImages() {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "images", "get", {}, {},
            (res: Array<any>) => {
                res = res.map(function (each) {
                    each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                    return each
                })
                resolve(res)
            },
            (err: object) => {
                reject(err)
                console.error(err)
            })
    })}

export {
    _apiRequest,
    buildImage,
    fetchLog,
    fetchAllObjects,
    installPackage,
    deleteLog,
    fetchConfig,
    updateConfig,
    fetchMeta,
    fetchImages
}