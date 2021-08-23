import axios, { Method } from 'axios'
import { setServer, getServer } from '../store/connectivity/server'
import {store} from '../store/store'

class RestClient {
    private axios
    private serverEndpoint
    constructor(serverEndpoint: string) {
        this.axios = axios
        this.serverEndpoint = serverEndpoint
    }
    setEndpoint(url:string) {
        this.serverEndpoint = url
    }
    getEndpoint() {
        return this.serverEndpoint
    }
    get(url: string) {
        return this.axios.get(this.serverEndpoint + url)
    }
    query(gql:string) {
        console.log(this.serverEndpoint)
        if (this.serverEndpoint==="") {
            this.setEndpoint(getServer(store.getState()))
        }
        console.log(getServer(store.getState()))
        return this.axios({
            method: 'post',
            url: this.serverEndpoint+"/api/query",
            data: {
                query: gql
            }
        })
    }
    infer(method: Method, url: string, payload: any) {
        const options = {
            method: method,
            url: this.serverEndpoint+url,
            data: payload
        }
        return this.axios.request(options)
    }
    subscribe(url:string) {
        return new EventSource(this.serverEndpoint+"/api/"+url)
    }
}

let restclient: RestClient = new RestClient("")

function setServerEndpoint(endpoint:string) {
    restclient.setEndpoint(endpoint)
    store.dispatch(setServer(endpoint))
}

function getServerEndpoint() {
    return store.getState().connectivity.server
}

function initServerEndpoint() {
    if (store.getState().connectivity.server) {
        let serverEndpoint = store.getState().connectivity.server
        restclient.setEndpoint(serverEndpoint)
    }
}

export { 
    restclient,
    setServerEndpoint,
    getServerEndpoint,
    initServerEndpoint
};