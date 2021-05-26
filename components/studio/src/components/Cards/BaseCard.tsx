// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import React from 'react';
import { WithStyles, withStyles, createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

const useStyles = (theme: Theme) => createStyles({
    root: {
      maxWidth: 275,
      padding: theme.spacing(2),
      display: 'flex',
      overflow: 'auto',
      flexDirection: 'column',
      margin: theme.spacing(2),
    },
    title: {
      fontSize: 14,
    },
    pos: {
      marginBottom: 12,
    },
  });


interface IProps extends WithStyles<typeof useStyles> {
    status: number;
}

class Basecard extends React.Component<IProps> {
    constructor(props: any){
        super(props);
        props = {
            status: 1,
        }
    }
  
    render() {
        return (
            <Card className={this.props.classes.root}>
            <CardContent>
                <Typography className={this.props.classes.title} color="textSecondary" gutterBottom>
                Word of the Day
                </Typography>
                <Typography variant="h5" component="h2">
                benevolent
                </Typography>
                <Typography className={this.props.classes.pos} color="textSecondary">
                adjective
                </Typography>
                <Typography variant="body2" component="p">
                well meaning and kindly.
                <br />
                {'"a benevolent smile"'}
                </Typography>
            </CardContent>
            <CardActions>
                <Button size="small">Learn More</Button>
            </CardActions>
            </Card>
        );
    }
}

export default withStyles(useStyles)(Basecard);