import axios from 'axios'
import { ConfigService } from './config'
import { getStatus } from './mock'
class SystemServiceMock {
  constructor (endpoint) {
    this.endpoint = endpoint
  }
  getStatus () {
    return getStatus()
  }
}

class SystemService {
  constructor (endpoint) {
    this.endpoint = endpoint
  }
  getStatus () {
    return getStatus()
  }
  getPackages () {
    return new Promise((resolve, reject) => {
      axios.get(this.endpoint + '/repos').then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  getRepoMeta (vendor, name) {
    return new Promise((resolve, reject) => {
      axios.get(this.endpoint + '/repo/meta/' + vendor + '/' + name).then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  runRepoSolver (vendor, name, solver) {
    return new Promise((resolve, reject) => {
      axios.post(this.endpoint + '/repo/running')
    })
  }
}

const configService = new ConfigService()

let systemService

if (configService.developerMode) {
  systemService = new SystemServiceMock(configService.endpoint)
} else {
  systemService = new SystemService(configService.endpoint)
}

export {
  systemService
}
