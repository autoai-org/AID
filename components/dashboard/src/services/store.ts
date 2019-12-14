import { applyMiddleware, compose, createStore } from 'redux'
import { createLogger } from 'redux-logger'
import { createBrowserHistory } from 'history'
import { routerMiddleware } from 'react-router-redux'

import createRootReducer from './reducer'

export const history = createBrowserHistory()

const myRouterMiddleware = routerMiddleware(history);

const getMiddleware = () => {
    if (process.env.NODE_ENV === 'production') {
        return applyMiddleware(myRouterMiddleware);
    } else {
        // Enable additional logging in non-production environments.
        return applyMiddleware(myRouterMiddleware, createLogger())
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
