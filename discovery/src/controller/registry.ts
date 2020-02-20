import { BaseContext } from 'koa';
import { addRegistry, getRegistries } from '../dynamo/action';

export default class RegistryController {
    public static async putRegistry (ctx: BaseContext) {
        const body = ctx.request.body;
        addRegistry(body.name, body.urlPrefix);
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'desc': 'success',
            'request': body
        };
    }
    public static async getRegistries (ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'results': await getRegistries()
        };
    }
}