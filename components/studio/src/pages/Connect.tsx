import React from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LinkIcon from '@material-ui/icons/Link';
import Typography from '@material-ui/core/Typography';
import { createStyles, Theme, withStyles, WithStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';

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
  isConnected: boolean;
}
class ConnectPage extends React.Component<WithStyles<typeof useStyles>, ConnectState>{
    constructor(props: any) {
        super(props);
        props = {
            classes: useStyles,
        };
        this.state = {
          isConnected: false,
        };
    }
    
    handleClick = () => {
      this.setState({
        isConnected: !this.state.isConnected
      });
      console.log(this.state.isConnected)
    }

    copyright() {
      return (
        <Typography variant="body2" color="textSecondary" align="center">
          {'Copyright © '}
          <Link color="inherit" href="https://aid.autoai.org/">
            A.I.D
          </Link>{' '}
          {new Date().getFullYear()}
          {'.'}
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
            <form className={this.props.classes.form} noValidate>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="server"
                label="Host"
                name="server"
                autoComplete="server"
                autoFocus
              />
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
              />
              <FormControlLabel
                control={<Checkbox value="remember" color="primary" />}
                label="Remember me"
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={this.props.classes.submit}
                onClick={this.handleClick}
              >
                Connect 
              </Button>
              <Grid container>
                <Grid item xs>
                  <Link href="#" variant="body2">
                    Forgot password?
                  </Link>
                </Grid>
                <Grid item>
                  <Link href="#" variant="body2">
                    {"Don't have an account? Sign Up"}
                  </Link>
                </Grid>
              </Grid>
            </form>
          </div>
          <Box mt={8}>
            {this.copyright}
          </Box>
        </Container>
      );
    }

}


export default withStyles(useStyles)(ConnectPage);
