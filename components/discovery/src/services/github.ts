import axios from 'axios'
const GITHUB_ENDPOINT = "https://api.github.com/"
import config from '../config'
import { Octokit } from "@octokit/rest";
import toml from 'toml'
import base64 from 'base-64'
const octokit = new Octokit({
    auth: config.GITHUB_TOKEN,
})
function tomlParse(inputBase64: string) {
    return toml.parse(base64.decode(inputBase64))
}
async function fetchGithubSummary(vendor, name) {
    const resp = await axios.get(GITHUB_ENDPOINT + "repos/" + vendor + "/" + name, {
        headers: {
            Authorization: "token " + config.GITHUB_TOKEN
        }
    })
    return resp.data
}

async function getGitHubDetails(vendor, name) {
    const configuration = await octokit.rest.repos.getContent({
        owner: vendor,
        repo: name,
        path: 'aid.toml',
    })
    const pretrained = await octokit.rest.repos.getContent({
        owner: vendor,
        repo: name,
        path: 'pretrained.toml',
    })
    const commits = await octokit.rest.repos.listCommits({
        owner: vendor,
        repo: name,
    });
    const readme = await octokit.rest.repos.getReadme({
        owner: vendor,
        repo: name,
    });
    return {
        'configuration': tomlParse(configuration.data['content']),
        'pretrained': tomlParse(pretrained.data['content']),
        'commits': commits.data,
        'readme': base64.decode(readme.data['content'])
    }
}

export {
    fetchGithubSummary,
    getGitHubDetails
}