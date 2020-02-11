import { request, summary, path, body, responsesAll, tagsAll, tags } from 'koa-swagger-decorator';
import { BaseContext } from 'koa';

@responsesAll({ 200: { description: 'Success'}, 500: { description: 'Server Error'}})
@tagsAll(['ping'])
export default class PingController {
    @request('get', '/ping')
    @summary('test if the server is running')
    public static async pong(ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = 'pong';
    }
}