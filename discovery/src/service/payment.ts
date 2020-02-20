import * as Stripe from 'stripe';
import { config } from '../config';

const stripe = new Stripe(config.stripe);

export default class PaymentService {
    public async createCustomer(email:string) {
        const customer: Promise<Stripe.customers.ICustomer> = stripe.customers.create({
            email: email
        })
        return customer
    }
}
