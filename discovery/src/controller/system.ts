import { BaseContext } from 'koa';
import { config } from '../config';

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
        }
    }
}
