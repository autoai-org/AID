import {ApolloClient, InMemoryCache} from '@apollo/client';

let client = new ApolloClient ({
    uri:"http://127.0.0.1:10590/query",
    cache: new InMemoryCache()
})

export default client;