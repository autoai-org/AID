import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import ContainerMui from '@material-ui/core/Container';
import Toolbar from '../toolbar/Toolbar'
import { ConnectedRouter } from 'connected-react-router'
import routes from '../../services/routes'
import { History } from 'history'
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';

interface ContainerProps {
    history: History;
}

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        content: {
            paddingLeft:'240px',
            paddingTop:'100px',
            flexGrow: 1,
            padding: theme.spacing(3),
        },
    }),
);

const Container = ({ history }: ContainerProps) => {
    const classes = useStyles();
    return (
        <React.Fragment>
            <CssBaseline />
            <ContainerMui>
                <Toolbar></Toolbar>
                <main className={classes.content}>
                    <ConnectedRouter history={history}>
                        {routes}
                    </ConnectedRouter>
                </main>
            </ContainerMui>
        </React.Fragment>
    )
}

export default Container