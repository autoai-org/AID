import React from 'react';
import MainLayout from './layouts/main'
import './App.css';

class Application extends React.Component {
  constructor(props:any) {
    super(props)
  }
  render() {
    return (
      <MainLayout></MainLayout>
    );
  }
}

export default Application;
