import {plain_get, plain_post, plain_put} from '@/middlewares/api.mdw'

const discovery_endpoint = "https://discovery.autoai.org"

function get_discovery_version() {
    return plain_get(discovery_endpoint+"/version")
}

export {
    get_discovery_version
}