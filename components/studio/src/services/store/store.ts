import serverReducer from '../store/connectivity/server';
import { persistStore, persistReducer } from 'redux-persist'
import storage from "redux-persist/lib/storage";
import {combineReducers, createStore, applyMiddleware} from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension';

const rootReducer = combineReducers({
    connectivity: serverReducer,
});

const persistConfig = {
    key:'aidstudio',
    storage
}
const persistedReducer = persistReducer(persistConfig, rootReducer)

export const store = createStore(
    persistedReducer,
    composeWithDevTools(
        applyMiddleware()
    ),
)
export const persistor = persistStore(store)
// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

