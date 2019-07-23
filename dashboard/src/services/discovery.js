// import axios from 'axios'
import { discoveryConfig } from './config'
import { getNews } from './mock'
import { _request } from './request'
class DiscoveryMock {
  constructor (endpoint) {
    self.endpoint = endpoint
  }

  fetchNews (num) {
    return getNews(num)
  }
}

class Discovery {
  constructor (endpoint) {
    self.endpoint = endpoint
  }

  fetchNews () {
    return _request('GET', self.endpoint + '/news')
  }

  putTicket (data) {
    return _request('PUT', self.endpoint + '/tickets', data)
  }
}
let discovery
if (discoveryConfig.developerMode) {
  discovery = new DiscoveryMock(discoveryConfig.endpoint)
} else {
  discovery = new Discovery(discoveryConfig.endpoint)
}

export {
  discovery
}
