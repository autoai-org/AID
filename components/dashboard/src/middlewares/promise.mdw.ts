function isPromise(v: any) {
    return v && typeof v.then === 'function';
}

export {
    isPromise
}