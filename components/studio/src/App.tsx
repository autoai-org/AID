import React from 'react';
import MainLayout from './layouts/MainLayout'
import SignIn from './pages/Authentication/Login'
import Register from './pages/Authentication/Register'
import './App.css';
import { initServerEndpoint } from './services/apis'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

type AppProps = {
}

type AppState = {
  connected: boolean,
}

class Application extends React.Component<AppProps, AppState> {
  componentDidMount() {
    initServerEndpoint()
  }
  constructor(props: AppProps) {
    super(props);
    this.state = {
      connected: false,
    }
    this.handler = this.handler.bind(this)
  }
  handler() {
    this.setState({
      connected: true
    })
  }
  render() {
    
    return (
      <Router>
        <Switch>
          <Route path="/signin"><SignIn /></Route>
          <Route path="/signup"><Register /></Route>
          <Route path="/"><MainLayout /></Route>
        </Switch>
      </Router>
    );
  }
}

export default Application;
