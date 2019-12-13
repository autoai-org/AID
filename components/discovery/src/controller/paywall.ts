import { BaseContext } from 'koa'
import paymentService from '../service/payment'
import Token from '../service/database/entity/token'
import { v4 as uuid } from 'uuid'
export default class PaywallController {
    public static async createCustomer(ctx: BaseContext) {
        const requestbody = ctx.request.body;
        const customer = await paymentService.createCustomer(requestbody.email);
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'customer': customer,
            'results': 'success'
        };
    }
    public static async createCharge(ctx: BaseContext) {
        const requestbody = ctx.request.body;
        const charge = await paymentService.createCharge(requestbody.email, requestbody.subtype);
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'charges': charge,
            'results': 'success'
        };
    }
    public static async createSubscription(ctx: BaseContext) {
        const requestbody = ctx.request.body;
        const charge = await paymentService.createSubscription(requestbody.email, requestbody.subtype);
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'charges': charge,
            'results': 'success'
        };
    }
    public static async checkout(ctx: BaseContext) {

    }
    public static async addAccessToken(ctx: BaseContext) {
        const requestbody = ctx.request.body
        const email = requestbody.email
        //TODO: Check if the user has paid
        const token = new Token({
            email: email,
            token: uuid(),
            validUntil: Date.now(),
            createdAt: Date.now(),
        })
        await token.save().then(function (res:any) {
            ctx.status = 200;
            ctx.body = {
                'code': 200,
                'token': res,
                'results': 'success'
            }
        }).catch(function (err:any) {
            ctx.status = 500;
            ctx.body = {
                'code': 500,
                'token': err,
                'results': 'failed'
            }
        })
    }
    public static async deleteAccessToken(ctx: BaseContext) {
        const token = ctx.request.body.token
        await Token.findOne({
            'token': token
        }).then(function (res:any) {
            res.deactivate()
            ctx.status = 200;
            ctx.body = {
                'code': 200,
                'token': res,
                'results': 'deactivated'
            }
        }).catch(function (err:any) {
            ctx.status = 500;
            ctx.body = {
                'code': 500,
                'token': err,
                'results': 'failed'
            }
        })
    }
    public static async queryAccessTokens(ctx: BaseContext) {
        const email = ctx.request.query.email
        await Token.find({
            'email': email
        }).then(function (res:any) {
            ctx.status = 200;
            ctx.body = {
                'code': 200,
                'tokens': res.map(function(each:any){
                    return {
                        'id': each.id,
                        'validUntil': each.validUntil,
                        'createdAt': each.createdAt
                    }
                }),
                'results': 'success'
            }
        }).catch(function (err:any) {
            ctx.status = 500;
            ctx.body = {
                'code': 500,
                'token': err,
                'results': 'failed'
            }
        })
    }
}