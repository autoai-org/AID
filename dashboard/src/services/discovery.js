import axios from 'axios'
import { discovery_config } from './config'

class Discovery {
    constructor(endpoint) {
        self.endpoint = endpoint
    }
}

const discovery = Discovery(discovery_config.endpoint)

export {
    discovery
}