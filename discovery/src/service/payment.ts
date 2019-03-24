import * as Stripe from 'stripe';
import { config } from '../config';

const stripe = new Stripe(config.stripe);

class PaymentService {
    public async createCustomer(email: string) {
        const customer = await stripe.customers.create({
            email
        }, function(err, customer) {
        });
        return customer;
    }
    public async createCharge(amount: number, customerId: string) {
        const charges = await stripe.charges.create({
            amount,
            currency: 'hkd',
            customer: customerId
        });
    }
}
const paymentService = new PaymentService();
export default paymentService;
