// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { createStyles, makeStyles, Theme } from '@material-ui/core';
import Basecard from '../components/Cards/BaseCard'
import RepoTable from '../components/Tables/RepoTable'

const useStyles = makeStyles((theme: Theme)=> 
    createStyles({
        container: {
            paddingTop: theme.spacing(4),
            paddingBottom: theme.spacing(4),
            display: 'flex',
            flexWrap: 'wrap',
        },
    })
)

export default function HomePage() {
    const classes = useStyles();
    return (
        <div>
        <Container maxWidth="lg" className={classes.container}>
                    <Grid item xs={12} md={4} lg={3}>
                        <Basecard status={2} ></Basecard>
                    </Grid>
                    <Grid item xs={12} md={4} lg={3}>
                        <Basecard status={3} ></Basecard>
                    </Grid>
                    <Grid item xs={12} md={4} lg={3}>
                        <Basecard status={3} ></Basecard>
                    </Grid>
                    <Grid item xs={12} md={4} lg={3}>
                        <Basecard status={3} ></Basecard>
                    </Grid>
                    <Grid item xs={12} md={4} lg={3}>
                        <Basecard status={3} ></Basecard>
                    </Grid>
                </Container>
                <RepoTable/>
                </div>
                
    )
}