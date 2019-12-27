import React from 'react'
import { ConnectedRouter } from 'connected-react-router'
import { History } from 'history'
import routes from './services/routes'
import { ThemeProvider } from '@material-ui/core/styles';
import globalTheme from './services/theme'

interface AppProps {
  history: History;
}

const App = ({ history }: AppProps) => {
  return (
    <ThemeProvider theme={globalTheme}>
    <ConnectedRouter history={history}>
      { routes }
    </ConnectedRouter>
    </ThemeProvider>
  )
}

export default App
