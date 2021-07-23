// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import { analyseGithubRepo } from './github'

export const serverEndpoint = 'https://discovery.autoai.org/'
// export const serverEndpoint = 'http://127.0.0.1:3001/'
function _get(endpoint:string) {
    return new Promise((resolve, reject)=>{
        axios.get(endpoint).then((res)=>{
            resolve(res)
        }).catch((err)=>{
            reject(err)
        })
    })
}
function _post(endpoint:string, payload:object) {
    return new Promise((resolve, reject)=>{
        axios.post(endpoint, payload).then((res)=>{
            resolve(res)
        }).catch((err)=>{
            reject(err)
        })
    })
}
export function findModelsByKeyword(keyword:string) {
    if (keyword==="") {
        return _get(serverEndpoint+"model/")
    }
    return _get(serverEndpoint+"model/keyword/"+keyword)
}

export async function analyseRepos(repos:any) {
    for (let i=0;i<repos.length;i++) {
        console.log(repos[i])
        if (repos[i].githubURL !== "" && repos[i].githubURL) {
            let vendor = repos[i].githubURL.split("/")[3]
            let name = repos[i].githubURL.split("/")[4]
            let result = await analyseGithubRepo(vendor, name, repos[i]._id)
            repos[i].analysis=result
        }
    }
    return repos
}

export function getModelInfo(vendor:string, name:string) {
    return _get(serverEndpoint+"model/info/"+vendor+"/"+name)
}

export function postNewRepo(repo:any) {
    return _post(serverEndpoint+"model/", repo)
}