import * as Stripe from 'stripe';
import { config } from '../config';

const stripe = new Stripe(config.stripe);

class PaymentService {
    public async createCustomer(email: string) {
        return new Promise((resolve, reject) => {
            stripe.customers.create({
                email
            }, function (err, customer) {
                resolve(customer)
            });
        })
    }
    public async createCharge(amount: number, customerId: string) {
        const charges = await stripe.charges.create({
            amount,
            currency: 'usd',
            customer: customerId
        });
    }
    public async createSubscription(customerId: string, subscriptionType: string) {
        return new Promise((resolve, reject) => {
            let amount
            if (subscriptionType === 'pro') {
                amount = 9
            } else if (subscriptionType === 'enterprise') {
                amount = 99
            } else {
                return {
                    'code': '400',
                    'help': 'not supported subscription type, only free, pro and enterprise is supported'
                }
            }
            stripe.charges.create({
                amount: amount,
                currency: 'usd',
                customer: customerId
            }, function(err, charge) {
                resolve(charge)
            })
            
        })
    }
}
const paymentService = new PaymentService();
export default paymentService;
