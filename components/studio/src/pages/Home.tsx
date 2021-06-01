// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import Container from '@material-ui/core/Container';
import React from 'react';
import { createStyles, Theme, withStyles } from '@material-ui/core';
import RepoTable from '../components/Tables/RepoTable'
import SolverStatus from '../components/Grids/SolverStatus'

const useStyles = ({ palette, spacing }: Theme) => createStyles({
    container: {
        paddingTop: spacing(4),
        paddingBottom: spacing(4),
        display: 'flex',
        flexWrap: 'wrap',
    },
})

interface HomeState {
    
}

class HomePage extends React.Component<any, HomeState> {
    constructor(props: any) {
        super(props);
        props = {
            classes: useStyles,
            handler: props.handler,
        };
        this.state = {
            solvers: []
        }
    }
    render() {
        return (
            <div>
                <SolverStatus />
                <RepoTable />
            </div >
        );
    }
}

export default withStyles(useStyles)(HomePage);