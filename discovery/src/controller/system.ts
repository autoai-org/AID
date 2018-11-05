import { BaseContext } from 'koa';

export default class SystemController {
    public static async getSystemStatus(ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = 'ok';
    }
}
