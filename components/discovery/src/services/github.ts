import axios from 'axios'
const GITHUB_ENDPOINT = "https://api.github.com/"
import config from '../config'

async function fetch_github_summary(vendor, name) {
    let resp = await axios.get(GITHUB_ENDPOINT+"repos/" + vendor + "/" + name, {
        headers: {
            Authorization: "token " + config.GITHUB_TOKEN
        }
    })
    return resp.data
}

export {
    fetch_github_summary
}