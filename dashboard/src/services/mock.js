import axios from 'axios'
import { discoveryConfig } from './config'

class Discovery {
  constructor (endpoint) {
    self.endpoint = endpoint
  }
  getSystemStatus () {
    return new Promise((resolve, reject) => {
      resolve({
        'status': 'ok'
      })
    })
  }
}

const discovery = Discovery(discoveryConfig.endpoint)

export {
    discovery
}
