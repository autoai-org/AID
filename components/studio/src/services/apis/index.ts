import {ApolloClient, InMemoryCache} from '@apollo/client';
import axios from 'axios'

let gqlclient = new ApolloClient ({
    uri:"http://127.0.0.1:17415/query",
    cache: new InMemoryCache()
})

let restclient = axios

export { 
    restclient,
    gqlclient
};