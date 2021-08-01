// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
import axios, { Method } from "axios"
import { serverEndpoint } from ".";
class HTTPClient {
    host: string;
    constructor(host: string) {
        this.host = host
    }
    request(method: Method, url: string, payload:any): Promise<any> {
        const options = {
            method: method,
            url: this.host + url,
            data: payload
        }
        return axios.request(options)
    }
}

const client = new HTTPClient(serverEndpoint)

export default client;