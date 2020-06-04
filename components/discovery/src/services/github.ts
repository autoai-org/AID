import axios from 'axios'
const GITHUB_ENDPOINT = "https://api.github.com/"
import config from '../config'

async function fetchGithubSummary(vendor, name) {
    const resp = await axios.get(GITHUB_ENDPOINT+"repos/" + vendor + "/" + name, {
        headers: {
            Authorization: "token " + config.GITHUB_TOKEN
        }
    })
    return resp.data
}

export {
    fetchGithubSummary
}