import { Context } from 'koa';
import * as passport from 'passport'
export default class UserController {
    /**
     * Create Session, a.k.a Log In
     * @param ctx 
     */
    public static async createSession(ctx: Context) {
        passport.authenticate('cognito', {
            successRedirect: '/',
            failureRedirect: '/login'        
        })
        ctx.status = 200;
        ctx.body = {
            'status': 'ok'
        }
    }
}
