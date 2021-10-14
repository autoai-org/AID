import axios from 'axios'
const GITHUB_ENDPOINT = "https://api.github.com/"
import config from '../config'
import { Octokit } from "@octokit/rest";
import toml from 'toml'
import base64 from 'base-64'
import { isEmpty } from '../utilities/object_checker'
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

async function getConfiguration(vendor, name) {
    let configuration
    try {
        configuration = await octokit.rest.repos.getContent({
            owner: vendor,
            repo: name,
            path: 'aid.toml',
        })
        configuration = tomlParse(configuration.data['content'])
    } catch(error) {
        console.log(error)
        configuration = "An error occurred"
    } finally {
        return configuration
    }
}
async function getPretrained(vendor, name) {
    let pretrained
    try {
        pretrained = await octokit.rest.repos.getContent({
            owner: vendor,
            repo: name,
            path: 'pretrained.toml',
        })
        pretrained = tomlParse(pretrained.data['content'])
        if (isEmpty(pretrained)) {
            throw new Error("empty pretrained.toml file");
        }
    } catch(error){
        pretrained = {
            "models": []
        }
    } finally {
        return pretrained
    }
}

async function getReadme(vendor, name) {
    let readme
    try {
        readme = await octokit.rest.repos.getReadme({
            owner: vendor,
            repo: name,
        })
        readme = base64.decode(readme.data['content'])
    } catch(error){
        readme = "**No Readme Found**"
    } finally {
        return readme
    }
}

async function getGitHubDetails(vendor, name) {
    const commits = await octokit.rest.repos.listCommits({
        owner: vendor,
        repo: name,
    });
    return {
        'configuration': await getConfiguration(vendor, name),
        'pretrained': await getPretrained(vendor, name),
        'commits': commits.data,
        'readme': await getReadme(vendor, name)
    }

}

export {
    fetchGithubSummary,
    getGitHubDetails
}