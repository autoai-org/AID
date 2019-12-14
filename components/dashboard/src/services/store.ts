import { applyMiddleware, createStore } from 'redux';
import { createLogger } from 'redux-logger'
import createHistory from 'history/createBrowserHistory';
import { routerMiddleware } from 'react-router-redux'
import { composeWithDevTools } from 'redux-devtools-extension/developmentOnly';

import reducer from './reducer'

export const history = createHistory();

const myRouterMiddleware = routerMiddleware(history);

const getMiddleware = () => {
    if (process.env.NODE_ENV === 'production') {
        return applyMiddleware(myRouterMiddleware);
    } else {
        // Enable additional logging in non-production environments.
        return applyMiddleware(myRouterMiddleware, createLogger())
    }
};

export const store = createStore(
    reducer, composeWithDevTools(getMiddleware()));
