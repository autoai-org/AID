// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import store from '@/store'
import dayjs from 'dayjs'
import { Log, Package } from '@/entities'
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

function _get_plain(path: string) {
    return new Promise((resolve, reject) => {
        _apiRequest(path, "get", {}, {},
            (res: Log) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
    })
}

function _put_plain(path: string, payload?: object) {
    return new Promise((resolve, reject) => {
        _apiRequest(path, "put", payload || {}, {},
            (res: any) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
    })
}

function _post_plain(path: string, data: object) {
    return new Promise((resolve, reject) => {
        _apiRequest(path, "post", data, {},
            (res: Log) => {
                resolve(res)
            },
            (err: object) => {
                reject(err)
            })
    })
}

function addWebhook(payload: string, status: string) {
    return _put_plain(endpoint + "webhooks", { payload: payload, status: status })
}

function getWebhooks() {
    return _get_plain(endpoint + "webhooks")
}

function getExtLogs(extName: string, logId: string) {
    return _get_plain(endpoint + "extensions/"+extName+"/logs/"+logId)
}

function getDatasets() {
    return _get_plain(endpoint +"experiments/dataset")
}

function buildImage(vendorName: string, packageName: string, solverName: string) {
    return _put_plain(endpoint + "packages/" + vendorName + "/" + packageName + "/" + solverName + "/images")
}

function installPackage(packageIdentifier: string) {
    return _post_plain(endpoint + "packages", { 'remote_url': packageIdentifier })
}

function createContainer(imageId: string) {
    return _put_plain(endpoint + "images/" + imageId + "/" + "containers")
}

function startContainer(containerId: string) {
    return _put_plain(endpoint + "containers/" + containerId + "/" + "run")
}

function testContainer(file: Blob) {
    const payload = new FormData()
    payload.append('file', file)
    return _post_plain(endpoint + "runnings/1234/infer", payload)
}

function uploadDataset(file: Blob) {
    const payload = new FormData()
    payload.append('file', file)
    return _put_plain(endpoint +"experiments/dataset/file", payload)
}

function addDataset(datasetPath:string, uncompressNow: boolean, name: string) {
    return _put_plain(endpoint+"experiments/dataset", {
        "datasetPath":datasetPath,
        "uncompress": uncompressNow,
        "name": name
    })
}

function fetchLog(id: string) {
    return _get_plain(endpoint + "logs/" + id)
}

function updateConfig(config: object) {
    return _post_plain(endpoint + "configs", config)
}

function fetchMeta(vendorName: string, packageName: string) {
    return _get_plain(endpoint + "packages/" + vendorName + "/" + packageName + "/meta")
}

function fetchSolverDockerfile(vendorName: string, packageName: string, solverName: string) {
    return _get_plain(endpoint + "solvers/" + vendorName + "/" + packageName + "/" + solverName + "/dockerfile")
}

function fetchContainers() {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "/containers", "get", {}, {},
            (res: Array<any>) => {
                res = res.map(function (each: any) {
                    each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                    return each
                })
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

function fetchPackages() {
    return new Promise((resolve, reject) => {
        _apiRequest(endpoint + "packages", "get", {}, {},
            (res: Array<any>) => {
                res = res.map(function (each) {
                    each.Frameworks = []
                    fetchMeta(each.Vendor, each.Name).then(function(res:any){
                        if (res.requirements.includes("opencv")) {
                            each.Frameworks.push("OpenCV")
                        }
                        if (res.requirements.includes("torch")) {
                            each.Frameworks.push("pyTorch")
                        }
                        if (res.requirements.includes("tensorflow")) {
                            each.Frameworks.push("TensorFlow")
                        }
                        if (res.requirements.includes("keras")) {
                            each.Frameworks.push("Keras")
                        }
                    })
                    each.CreatedAt = dayjs(each.CreatedAt).format("DD/MM/YYYY HH:mm")
                    return each
                })
                console.log(res)
                resolve(res)
                store.commit('set' + "packages", res)
            },
            (err: object) => {
                reject(err)
                console.error(err)
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

export {
    getExtLogs,
    fetchSolverDockerfile,
    fetchContainers,
    buildImage,
    createContainer,
    uploadDataset,
    getDatasets,
    fetchLog,
    fetchAllObjects,
    installPackage,
    deleteLog,
    fetchConfig,
    addDataset,
    updateConfig,
    fetchMeta,
    testContainer,
    startContainer,
    fetchImages,
    addWebhook,
    getWebhooks,
    fetchPackages
}