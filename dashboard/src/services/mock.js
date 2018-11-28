import axios from 'axios'
import { discovery_config } from './config'

class Discovery {
    constructor(endpoint) {
        self.endpoint = endpoint
    }
    getSystemStatus () {
        return new Promise((resolve, reject) => {
            resolve({
                'status':'ok'
            })
        })
    }
}

const discovery = Discovery(discovery_config.endpoint)

export {
    discovery
}