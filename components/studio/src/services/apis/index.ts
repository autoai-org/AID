import {ApolloClient, InMemoryCache} from '@apollo/client';
import axios from 'axios'

let gqlclient = new ApolloClient ({
    uri:"http://127.0.0.1:10590/query",
    cache: new InMemoryCache()
})

let restclient = axios

export { 
    restclient,
    gqlclient
};