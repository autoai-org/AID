// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { createStyles, Theme, makeStyles } from '@material-ui/core';
import Solvercard from '../Cards/SolverCard'
import { useQuery } from '@apollo/client'
import { ALL_SOLVERS } from '../../services/apis/queries'
import { Solver } from '../../services/interfaces/entities'

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        container: {
            paddingTop: theme.spacing(4),
            paddingBottom: theme.spacing(4),
            display: 'flex',
            flexWrap: 'wrap',
        },
    })
)

export default function SolverStatus() {
    const { loading, error, data } = useQuery(ALL_SOLVERS);
    if (loading) return <p>Loaindg</p>;
    if (error) return <p>Error :(</p>;
    console.log(data);
    return (
        <div>
            {
                <Container>
                    <Grid item xs={12} md={4} lg={3}>
                        {data.solvers.map((value:Solver, index:number) => {
                            return <Solvercard key={index} status={value.status} title={value.name} package={value.repository.name}/>
                        })}
                    </Grid>
                </Container>
            }
        </div>
    )
}