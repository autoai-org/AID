// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import { analyseGithubRepo } from './github'
const serverEndpoint = 'https://discovery.autoai.org/'

function _get(endpoint:string) {
    return new Promise((resolve, reject)=>{
        axios.get(endpoint).then((res)=>{
            resolve(res)
        }).catch((err)=>{
            reject(err)
        })
    })
}

function findModelsByKeyword(keyword:string) {
    if (keyword==="") {
        return _get(serverEndpoint+"model/")
    }
    return _get(serverEndpoint+"model/keyword/"+keyword)
}

async function analyseRepos(repos:any) {
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

export {
    serverEndpoint,
    analyseRepos,
    findModelsByKeyword
}
