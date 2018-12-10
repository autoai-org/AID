class ConfigService {
  constructor () {
    this.read()
  }
  read () {
    const persistConfig = JSON.parse(localStorage.getItem('cvpm-config'))
    if (persistConfig === null) {
      this.loadDefault()
    } else {
      this.endpoint = persistConfig['endpoint']
      this.developerMode = persistConfig['developerMode']
      this.hintMode = persistConfig['hintMode']
    }
  }
  loadDefault () {
    this.endpoint = 'https://instance.cvtron.xyz'
    this.developerMode = false
  }
  write () {
    const persistConfig = {
      'endpoint': this.endpoint,
      'developerMode': this.developerMode,
      'hintMode': this.hintMode
    }
    localStorage.setItem('cvpm-config', JSON.stringify(persistConfig))
    this.read()
  }
}

const configService = new ConfigService()

const discoveryConfig = {
  'endpoint': 'http://127.0.0.1:3000'
}

export {
  discoveryConfig,
  ConfigService,
  configService
}
