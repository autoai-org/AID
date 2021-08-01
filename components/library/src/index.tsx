import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import Landing from './pages/Landing'
import SearchResult from './pages/SearchResult'
import Login from './pages/Authentication/Login'
import Signup from './pages/Authentication/Register'
import Publish from './pages/Publish'
import Profile from './pages/Profile'
import NotFound from './pages/NotFound'
import reportWebVitals from './reportWebVitals'
import ModelDetail from './components/Models/ModelDetail'
import { FirebaseAuthProvider } from '@react-firebase/auth'
import firebase from "firebase/app"
import "firebase/auth"
import { config } from './services/firebase'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"
import ReactGA, { ga } from 'react-ga'
import * as Sentry from "@sentry/react";
import { Integrations } from "@sentry/tracing";

Sentry.init({
  dsn: "https://af50717a734e41ada1ada89d286667e9@o465689.ingest.sentry.io/5888093",
  integrations: [new Integrations.BrowserTracing()],

  // Set tracesSampleRate to 1.0 to capture 100%
  // of transactions for performance monitoring.
  // We recommend adjusting this value in production
  tracesSampleRate: 1.0,
});

const TRACKING_ID = "281386144";
ReactGA.initialize(TRACKING_ID);

ReactDOM.render(
  <React.StrictMode>
    <FirebaseAuthProvider firebase={firebase} {...config}>
      <Router>
        <Switch>
          <Route path="/signup">
            <Signup />
          </Route>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/search">
            <SearchResult />
          </Route>
          <Route path="/profile">
            <Profile />
          </Route>
          <Route path="/publish">
            <Publish />
          </Route>
          <Route path="/model/:vendor/:name">
            <ModelDetail />
          </Route>
          <Route exact path="/">
            <Landing />
          </Route>
          <Route component={NotFound} />
        </Switch>
      </Router>
    </FirebaseAuthProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

function sendToAnalytics({ id, name, value }:any) {
  ga('send', 'event', {
    eventCategory: 'Web Vitals',
    eventAction: name,
    eventValue: Math.round(name === 'CLS' ? value * 1000 : value), // values must be integers
    eventLabel: id, // id unique to current page load
    nonInteraction: true, // avoids affecting bounce rate
  });
}

reportWebVitals(sendToAnalytics);