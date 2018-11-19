import { Context } from 'koa';
import { addRegistry, getRegistries } from '../dynamo/action';

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
    public static async getRegistries (ctx: Context) {
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'results': await getRegistries()
        };
    }
}