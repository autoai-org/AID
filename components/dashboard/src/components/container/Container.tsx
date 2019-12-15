import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';

import Toolbar from '../toolbar/Toolbar'
export default class Container extends React.Component {
    render() {
        return (
            <React.Fragment>
                <CssBaseline />
                <Toolbar></Toolbar>
            </React.Fragment>
        )
    }
}