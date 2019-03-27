import { BaseContext } from 'koa'
import paymentService from '../service/payment'
import Token, { ITokenModel } from '../service/database/entity/token'
import { v4 as uuid } from 'uuid'
export default class PaywallController {
    public static async createCustomer(ctx: BaseContext) {
        const requestbody = ctx.request.body;
        const customer = await paymentService.createCustomer(requestbody.email);
        console.log(customer)
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
        const token: ITokenModel = new Token({
            email: email,
            token: uuid(),
            validUntil: Date.now()
        })
        token.save()
    }
}