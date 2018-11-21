import { BaseContext, Context } from 'koa';
import { config } from '../config';

import { getCount, updateCount } from '../parse';

export default class SystemController {
    public static async getSystemStatus(ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'status': 'running',
        };
    }
    public static async getMetaInfo (ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'version': config.version,
            'port': config.port
        };
    }
    public static async getStatInfo (ctx: Context) {
        const body = ctx.request.query;
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'results': await getCount(+body.limit)
        };
    }
    public static async putStatInfo (ctx: Context) {
        const body = ctx.request.body;
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'request': await updateCount(+body.users, +body.registries, +body.models)
        };
    }
}
