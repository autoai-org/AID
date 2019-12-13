import { BaseContext } from 'koa';
import { getUserInfo } from '../service/auth';
import { githubService } from '../service/github'
export default class UserController {
    /**
     * Validate Access Token, a.k.a Get User information
     * @param ctx
     */
    public static async validateAccessToken(ctx: BaseContext) {
        const body = ctx.request.body;
        getUserInfo(body.accessToken);
    }
    public static async socialAuthenticate(ctx: BaseContext) {
        let requestBody = ctx.request.body;
        let res =  await githubService.finishAuth(requestBody.code,requestBody.redirect)
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'body': res.data,
            'results': 'success'
        }
    }
}
