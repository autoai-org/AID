import { ipcRenderer } from "electron";
import store from '../store'
function isElectron() {
    const userAgent = navigator.userAgent.toLowerCase();
    return userAgent.indexOf(' electron/') > -1
}

function init_ipc_hooks() {
    ipcRenderer.on('disconnected', (event, args)=>{
        store.commit('setConnected', false)
    })
}

export {
    init_ipc_hooks,
    isElectron
}