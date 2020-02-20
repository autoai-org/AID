// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import store from '@/store/index'

function alert(info:string, title:string) {
    store.state.alert_info = info
    store.state.alert_title = title
}

export {
    alert
}