// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { applyMiddleware, compose, createStore } from 'redux'
import { createLogger } from 'redux-logger'
import { createBrowserHistory } from 'history'
import { routerMiddleware } from 'react-router-redux'
import apiMiddleware from '../middlewares/api.mdw'

import createRootReducer from './reducer'

export const history = createBrowserHistory()

const myRouterMiddleware = routerMiddleware(history);

const getMiddleware = () => {
    if (process.env.NODE_ENV === 'production') {
        return applyMiddleware(myRouterMiddleware, apiMiddleware);
    } else {
        // Enable additional logging in non-production environments.
        return applyMiddleware(myRouterMiddleware, apiMiddleware, createLogger())
    }
};
export default function configureStore(preloadedState: object) {
    const store = createStore(
        createRootReducer(history), // root reducer with router state
        preloadedState,
        compose(
            getMiddleware(),
        ),
    )
    if (module.hot) {
        // Enable Webpack hot module replacement for reducers
        module.hot.accept('./reducer', () => {
            store.replaceReducer(createRootReducer(history));
        });
    }

    return store
}
