import { BaseContext } from 'koa';
import paymentService from '../service/payment'

export default class PaywallController {
    public static async createCustomer (ctx: BaseContext) {
        const requestbody = ctx.request.body
        const customer = await paymentService.createCustomer(requestbody.email)
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'results': 'success',
            'customer': customer
        };
    }
}