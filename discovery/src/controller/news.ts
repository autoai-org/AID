import { Context } from 'koa'
import axios from 'axios'
import { config } from '../config';

export default class NewsController {
    public static async queryNews (ctx: Context) {
        console.log('quering news')
        await axios.get(config.blogURL + '/posts').then(function (res:any) {
            console.log(res)
            ctx.status = 200;
            ctx.body = res.data.data.posts;
        }).catch(function (err) {
            console.log(err)
            ctx.status = 500;
            ctx.body = 'An error occured on the blog service';
        })
    }
}