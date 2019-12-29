// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

function isPromise(v: any) {
    return v && typeof v.then === 'function';
}

export {
    isPromise
}