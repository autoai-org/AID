// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT


class HTTPRequest {
    method: string
    url: string
    contentType: string
    payload: { [key: string]: any }
    constructor(method: string, url: string, contentType: string, payload: object) {
        this.contentType = contentType
        this.method = method;
        this.url = url;
        this.payload = payload;
    }
    _preparePayload() {
        if ('file' in this.payload) {
            let payload = new FormData();
            for (let key in this.payload) {
                payload.append(key, this.payload[key]);
            }
            var blob = new Blob([this.payload['file']], { type: "text/plain" })
            payload.delete('file')
            payload.append('file', blob)
            this.payload = payload
        }
    }
    do(): Promise<any> {
        this._preparePayload();
        let xhr = new XMLHttpRequest();
        xhr.open(this.method, this.url, true);
        xhr.setRequestHeader("Content-Type", this.contentType);
        return new Promise((resolve, reject) => {
            xhr.onload = function () {
                if (xhr.status >= 200 && xhr.status < 300) {
                    resolve(xhr.response);
                } else {
                    reject(xhr.statusText);
                }
            }
            xhr.onerror = function () {
                reject(xhr.statusText);
            }
            xhr.send();
        })
    }
  }
  
  class Client {
    host: string
    payload: FormData;
    endpoints: { [key: string]: string };
    constructor(host:string) {
        this.host = host;
        this.endpoints = {};
        this.payload = new FormData();
    }
    flush() {
        this.payload = new FormData();
    }
    setHost(host: string) {
        this.host = host;
    }
    async create_request(vendor:string, packageName:string, solver:string, payload:object): Promise<HTTPRequest> {
        let endpoint;
        if (vendor+packageName+solver in this.endpoints) {
            endpoint = this.endpoints[vendor+packageName+solver];
        } else {
            let params = {'vendor': vendor, 'package': packageName, 'solver': solver}
            let request = new HTTPRequest("POST", this.host+":17415/preflight","application/json", params)
            let res = JSON.parse(await request.do())
            endpoint = res['port']
            this.endpoints[vendor + packageName + solver] = endpoint
        }
        return new HTTPRequest("POST", endpoint, "multipart/form-data", payload)
    }
  }
  
  export {Client}