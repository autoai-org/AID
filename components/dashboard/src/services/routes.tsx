// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import React from 'react'
import { Route, Switch } from 'react-router'
import Home from '../pages/home/Home'
import Packages from '../pages/packages/Packages'

const routes = (
  <div>
    <Switch>
      <Route exact path="/" component={Home} />
      <Route path="/packages" component={Packages} />
    </Switch>
  </div>
)

export default routes
