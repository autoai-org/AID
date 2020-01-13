import { request, summary, path, body, responsesAll, tagsAll, tags } from 'koa-swagger-decorator';
import { BaseContext } from 'koa';

@responsesAll({ 200: { description: 'success'}, 400: { description: 'bad request'}, 401: { description: 'unauthorized, missing/wrong jwt token'}})
@tagsAll(['ping'])
export default class PingController {
    @request('get', '/ping')
    @summary('test if the server is running')
    public static async pong(ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = 'pong';
    }
}