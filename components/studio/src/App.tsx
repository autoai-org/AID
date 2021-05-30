import React from 'react';
import MainLayout from './layouts/main'
import ConnectPage from './pages/Connect'
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
    if (this.state.connected) {
      return (
        <MainLayout></MainLayout>
      );
    }
    else {
      return (
        <ConnectPage handler={this.handler}></ConnectPage>
      );
    }
  }
}

export default Application;
