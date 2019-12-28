import React from 'react'
import { History } from 'history'
import { ThemeProvider } from '@material-ui/core/styles';
import globalTheme from './services/theme'
import Container from './components/container/Container'

interface AppProps {
  history: History;
}

const App = ({ history }: AppProps) => {
  return (
    <ThemeProvider theme={globalTheme}>
      <Container history={history}></Container>
    </ThemeProvider>
  )
}

export default App
