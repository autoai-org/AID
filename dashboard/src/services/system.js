import { ConfigService } from './config'
import { getStatus } from './mock'
import { _get, _post } from './base'

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
    return _get(this.endpoint + '/system')
  }
  getPackages () {
    return _get(this.endpoint + '/repos')
  }
  getRepoMeta (vendor, name) {
    return _get(this.endpoint + '/repo/meta/' + vendor + '/' + name)
  }
  getRunningSolver (vendor, name) {
    return _get(this.endpoint + '/solvers/running/' + vendor + '/' + name)
  }
  runRepoSolver (vendor, name, solver, port) {
    return _post(this.endpoint + '/repo/running', {
      'vendor': vendor,
      'name': name,
      'solver': solver,
      'port': port
    })
  }
  testRepoSolver (vendor, packageName, solver, parameters, file) {
    let payload = new FormData()
    payload.append('file', file)
    for (let i = 0; i < parameters.length; i++) {
      payload.append(parameters[i].key, parameters[i].value)
    }
    return _post(this.endpoint + '/engine/solvers/' + vendor + '/' + packageName + '/' + solver, payload)
  }
  installRepo (type, id) {
    // if type === 'git', id => git url
    return _post(this.endpoint + '/repos', {
      type: type,
      url: id
    })
  }
  // contrib
  // datasets
  getOpenDatasets () {
    return _get(this.endpoint + '/contrib/datasets')
  }
  SyncDatabase (datasetsUrl) {
    return _post(this.endpoint + '/contrib/datasets/registries', {
      url: datasetsUrl
    })
  }
  // files
  getMyFiles () {
    return _get(this.endpoint + '/contrib/files/list')
  }
  // inspector
  getInspectorInfo () {
    return _get(this.endpoint + '/_inspector')
  }
  uploadFile (file, type) {
    let payload = new FormData()
    payload.append('file', file)
    return _post(this.endpoint + '/contrib/files/upload/' + type, payload)
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
