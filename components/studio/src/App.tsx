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
  }
  render() {
    if (this.state.connected) {
      return (
        <MainLayout></MainLayout>
      );
    }
    else {
      return (
        <ConnectPage></ConnectPage>
      );
    }
  }
}

export default Application;
