import React from 'react'
import { Route, Switch } from 'react-router'
import Home from '../pages/home/Home'
import Dashboard from '../pages/dashboard/Dashboard'

const routes = (
  <div>
    <Switch>
      <Route exact path="/" component={Home} />
      <Route path="/dashboard" component={Dashboard} />
    </Switch>
  </div>
)

export default routes
