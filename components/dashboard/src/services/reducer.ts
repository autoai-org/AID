// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { connectRouter } from 'connected-react-router'
import { combineReducers } from 'redux';
import { History } from 'history';

const createRootReducer = (history: History<any>) => combineReducers({
    router: connectRouter(history)
})

export default createRootReducer
