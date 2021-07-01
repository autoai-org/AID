import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { ApolloProvider } from '@apollo/client/react';
import { gqlclient } from './services/apis';
import { store, persistor } from './services/store/store';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react'
import Loading from './components/Loading'
ReactDOM.render(

  <React.StrictMode>
    <Provider store={store}>
    <PersistGate loading={<Loading/>} persistor={persistor}>
      <ApolloProvider client={gqlclient}>
        <App />
      </ApolloProvider>
      </PersistGate>
    </Provider>
  </React.StrictMode>,
  document.getElementById('root')
);

reportWebVitals(console.log);
