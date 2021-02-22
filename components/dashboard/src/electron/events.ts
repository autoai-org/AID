import store from '../store'
import { isElectron } from './index';

function init_ipc_hooks() {
    /**
     * Everytime importing electron modules, please check if this application is running in electron environment. When it is runnning in browser, it cannot import the electron module.
     * More specifically, do not use global import electron from 'electron'.
     * Use: if (isElectron) { const electronModule = require("electron") }
     */
    if (isElectron()) {
        const electronModule = require("electron")
        electronModule.ipcRenderer.on('disconnected', () => {
            store.commit('setConnected', false)
        })
    }
}

export {
    init_ipc_hooks,
}