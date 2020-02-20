import { BaseContext } from 'koa';
import { freshdesk } from '../service/ticket';

export default class TicketController {
    public static async createTicket (ctx: BaseContext) {
        const body = ctx.request.body;
        await freshdesk.createTicket({
            name: body.name,
            email: body.email,
            subject: body.subject,
            description: body.description,
            status: 2,
            priority: body.priority
        }).then(function (res: any) {
            ctx.status = 200;
            ctx.body = {
                'code': 200,
                'results': res.data
            };
        }).catch(function (err: any) {
            ctx.status = 400;
            ctx.body = err.response.data;
        });
    }
}