// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { restclient } from './index'

function connectServer(url:string) {
    return restclient.get(url)
}

export {
    connectServer
}