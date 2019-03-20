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
    return new Promise((resolve, reject) => {
      axios.get(this.endpoint + '/system').then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
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
  getRunningSolver (vendor, name) {
    return new Promise((resolve, reject) => {
      axios.get(this.endpoint + '/solvers/running/' + vendor + '/' + name).then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  runRepoSolver (vendor, name, solver, port) {
    return new Promise((resolve, reject) => {
      axios.post(this.endpoint + '/repo/running', {
        'vendor': vendor,
        'name': name,
        'solver': solver,
        'port': port
      }).then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  testRepoSolver (vendor, packageName, solver, parameters, file) {
    return new Promise((resolve, reject) => {
      let payload = new FormData()
      payload.append('file', file)
      for (let i = 0; i < parameters.length; i++) {
        payload.append(parameters[i].key, parameters[i].value)
      }
      axios.post(this.endpoint + '/engine/solvers/' + vendor + '/' + packageName + '/' + solver, payload).then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  installRepo (type, id) {
    // if type === 'git', id => git url
    return new Promise((resolve, reject) => {
      axios.post(this.endpoint + '/repos', {
        type: type,
        url: id
      }).then(function (res) {
        resolve(res)
      }).then(function (err) {
        reject(err)
      })
    })
  }
  // contrib
  // datasets
  getAllDatasets () {
    return new Promise((resolve, reject) => {
      axios.get(this.endpoint + '/contrib/datasets', {
      }).then(function (res) {
        resolve(res)
      }).then(function (err) {
        reject(err)
      })
    })
  }
  SyncDatabase (datasetsUrl) {
    return new Promise((resolve, reject) => {
      axios.post(this.endpoint + '/contrib/datasets/registries', {
        url: datasetsUrl
      }).then(function (res) {
        resolve(res)
      }).then(function (err) {
        reject(err)
      })
    })
  }
  // inspector
  getInspectorInfo () {
    return new Promise((resolve, reject) => {
      axios.get(this.endpoint + '/_inspector').then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
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
