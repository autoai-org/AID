// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import axios from 'axios'
import { serverEndpoint } from './api'
const githubEndpoint = "https://raw.githubusercontent.com/"

async function fetch_requirements(vendor: string, name: string) {
    let requirments = await axios.get(githubEndpoint + vendor + "/" + name + "/master/requirements.txt")
    return requirments.data.split("=").join("\n").split("\n").filter(function (each: any) {
        if (isNaN(each) && each !== "" && !each.includes(".")) {
            return true
        }
    })
}

async function fetch_github_summary(packageId: string) {
    return await (await axios.get(serverEndpoint + "model/meta/" + packageId)).data
}

async function analyseGithubRepo(vendor: string, name: string, packageId: string) {
    let requirements = await fetch_requirements(vendor, name)
    let githubSummary = await fetch_github_summary(packageId)
    return { 'requirements': requirements, 'summary': githubSummary }
}

export {
    analyseGithubRepo
}
