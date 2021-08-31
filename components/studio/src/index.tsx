import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { store, persistor } from './services/store/store';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react'
import Loading from './components/Loading'
import firebase from "firebase/app"
import "firebase/auth"
import { config } from './services/apis/firebase'
import { FirebaseAuthProvider } from '@react-firebase/auth'
import {transitions, positions, Provider as AlertProvider} from 'react-alert'
import AlertTemplate from './services/utilities/alert';
const options = {
  // you can also just use 'bottom center'
  position: positions.TOP_CENTER,
  timeout: 3000,
  offset: '20px',
  // you can also just use 'scale'
  transition: transitions.FADE,
}

ReactDOM.render(

  <React.StrictMode>
    <AlertProvider template={AlertTemplate} {...options}>
    <FirebaseAuthProvider firebase={firebase} {...config}>
    <Provider store={store}>
      <PersistGate loading={<Loading />} persistor={persistor}>
          <App />
      </PersistGate>
    </Provider>
    </FirebaseAuthProvider>
    </AlertProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

reportWebVitals(console.log);
