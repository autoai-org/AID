import { BaseContext } from 'koa';
import paymentService from '../service/payment';

export default class PaywallController {
    public static async createCustomer (ctx: BaseContext) {
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
    public static async createCharge (ctx: BaseContext) {
        const requestbody = ctx.request.body;
        const charge = await paymentService.createCharge(requestbody.email, requestbody.subtype);
        ctx.status = 200;
        ctx.body = {
            'code': 200,
            'charges': charge,
            'results': 'success'
        };
    }
}