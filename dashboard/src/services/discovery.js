// import axios from 'axios'
import { discoveryConfig } from './config'
import { getNews } from './mock'
import { newsService } from './news'
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
    return newsService.fetchNews()
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
