import React from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Box from '@material-ui/core/Box';
import LinkIcon from '@material-ui/icons/Link';
import Typography from '@material-ui/core/Typography';
import { createStyles, Theme, withStyles, WithStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { connect } from '../services/apis/system.query'
const useStyles = ({ palette, spacing }: Theme) => createStyles({
  paper: {
    marginTop: spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: spacing(1),
    backgroundColor: palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: spacing(1),
  },
  submit: {
    margin: spacing(3, 0, 2),
  },
});

interface ConnectState {
  url: string;
  isConnected: boolean;
  error: string;
  sysInfo: any
}

class ConnectPage extends React.Component<any, ConnectState>{
  constructor(props: any) {
    super(props);
    props = {
      classes: useStyles,
      handler: props.handler,
    };
    this.state = {
      url: "http://127.0.0.1:10590",
      isConnected: false,
      error: "",
      sysInfo: {}
    };
  }

  handleConnect = () => {
    let self = this
    connect(this.state.url + "/ping").then(function (res) {
      self.setState({ sysInfo: res.data })
      self.setState({ isConnected: true })
    })
  }

  handleConfirm = () => {
    this.props.handler()
  }

  handleUrlChange = (e: any) => {
    this.setState({ url: e.target.value })
  }
  copyright() {
    return (
      <Typography variant="body2" color="textSecondary" align="center">
        {'You are connecting to '}
        <Link color="inherit" href="https://aid.autoai.org/">
          A.I.D
          </Link>{' '}
            Server
      </Typography>
    );
  }

  render() {
    return (
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <div className={this.props.classes.paper}>
          <Avatar className={this.props.classes.avatar}>
            <LinkIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Connect
            </Typography>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            value={this.state.url}
            id="server"
            label="Host"
            onChange={this.handleUrlChange}
            name="server"
            autoComplete="server"
            autoFocus
          />
          {!this.state.isConnected &&
            <Button
              fullWidth
              variant="contained"
              color="primary"
              className={this.props.classes.submit}
              onClick={this.handleConnect}
            >
              Connect
              </Button>
          }
          {this.state.isConnected &&
            <div>
            <p>Hostname: {this.state.sysInfo.hostname}</p>
            <Button
              fullWidth
              variant="contained"
              color="primary"
              className={this.props.classes.submit}
              onClick={this.handleConfirm}
            >
              Confirm
              </Button>
              </div>
          }
        </div>
        <Box mt={8}>
          {this.copyright}
        </Box>
      </Container>
    );
  }

}


export default withStyles(useStyles)(ConnectPage);
