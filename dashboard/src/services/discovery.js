import axios from 'axios'
import { discoveryConfig } from './config'

class Discovery {
  constructor (endpoint) {
    self.endpoint = endpoint
  }
}

const discovery = Discovery(discoveryConfig.endpoint)

export {
    discovery
}
