function isElectron() {
    const userAgent = navigator.userAgent.toLowerCase();
    return userAgent.indexOf(' electron/') > -1
}

export {
    isElectron
}