import { request, summary, path, body, responsesAll, tagsAll, tags } from 'koa-swagger-decorator';
import { BaseContext } from 'koa';

@responsesAll({ 200: { description: 'Success'}, 500: { description: 'Server Error'}})
@tagsAll(['version'])
export default class PingController {
    @request('get', '/version')
    @summary('Get Current Version of Discovery Service')
    public static async pong(ctx: BaseContext) {
        ctx.status = 200;
        ctx.body = {'version':'1.0.0'};
    }
}