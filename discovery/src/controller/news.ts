import axios from 'axios';
import { BaseContext } from 'koa';
import { config } from '../config';

export default class NewsController {
    public static async queryNews (ctx: BaseContext) {
        await axios.get(config.blogURL + '/posts').then(function (res: any) {
            ctx.status = 200;
            ctx.body = res.data.data.posts;
        }).catch(function (err) {
            ctx.status = 500;
            ctx.body = 'An error occured on the blog service';
        });
    }
}