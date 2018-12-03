import { Context } from 'koa';
import { getUserInfo } from '../service/auth'
export default class UserController {
    /**
     * Validate Access Token, a.k.a Get User information
     * @param ctx
     */
    public static async validateAccessToken(ctx: Context) {
        const body = ctx.request.body;
        getUserInfo(body.accessToken)
    }
}
