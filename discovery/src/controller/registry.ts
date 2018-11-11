import { Context } from 'koa';
import { addRegistry } from '../dynamo/action';

export default class RegistryController {
    public static async putRegistry (ctx: Context) {
        const body = ctx.request.body;
        addRegistry(body.name, body.urlPrefix);
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'desc': 'success',
            'req': body
        };
    }
}