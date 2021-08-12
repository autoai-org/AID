import {ApolloClient, InMemoryCache} from '@apollo/client';
import axios from 'axios'
import { setServer } from '../store/connectivity/server'
import {store} from '../store/store'

let serverEndpoint = "http://localhost:17415"

class RestClient {
    private axios
    constructor() {
        this.axios = axios
    }
    get(url: string) {
        return this.axios.get(url)
    }
    query(gql:string) {
        return this.axios({
            method: 'post',
            url: serverEndpoint+"/api/query",
            data: {
                query: gql
            }
        })
    }
}

let restclient = new RestClient()

function setServerEndpoint(endpoint:string) {
    serverEndpoint = endpoint
    store.dispatch(setServer(endpoint))
}

function getServerEndpoint() {
    return store.getState().connectivity.server
}

function initServerEndpoint() {
    if (store.getState().connectivity.server) {
        serverEndpoint = store.getState().connectivity.server
    }
}

export { 
    serverEndpoint,
    restclient,
    setServerEndpoint,
    getServerEndpoint,
    initServerEndpoint
};