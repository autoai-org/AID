import { config } from './config';

const Sentry = require('@sentry/node');
Sentry.init({});

export {
    Sentry
};