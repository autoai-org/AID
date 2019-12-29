// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { API, FETCH_PACKAGES, LOAD_PACKAGES } from '../actions/types'


function _get(endpoint: string, onSuccess: Function) {
    return {
        type: API,
        payload: {
            url: endpoint,
            method: "GET",
            data: null,
            onSuccess: onSuccess,
            onFailure: () => {
                console.error("Cannot Load Content from " + endpoint)
            },
            label: FETCH_PACKAGES
        }
    }
}

function GetAllPackages() {
    return _get("https://localhost/packages", (data:object)=>{
        return {
            type: LOAD_PACKAGES,
            payload: data
        }
    })
}

export {
    GetAllPackages,
}