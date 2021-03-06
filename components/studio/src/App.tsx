import React from 'react';
import MainLayout from './layouts/MainLayout'
import './App.css';

type AppProps = {
}

type AppState  = {
  connected: boolean,
}

class Application extends React.Component<AppProps, AppState> {
  constructor(props:AppProps) {
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
        <MainLayout/>
      );
    }
}

export default Application;
