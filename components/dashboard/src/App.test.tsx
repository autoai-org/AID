import React from 'react';
import { render } from '@testing-library/react';
import { Provider } from 'react-redux'

import App from './App';
import configureStore, { history } from './services/store'

const store = configureStore({})

test('renders learn react link', () => {
  const { getByText } = render(
    <Provider store={store}>
    <App history={history} />
    </Provider>

  );
  const linkElement = getByText(/123/i);
  expect(linkElement).toBeInTheDocument();
});
