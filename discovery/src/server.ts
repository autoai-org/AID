import * as cors from '@koa/cors';
import * as Koa from 'koa';
import bodyParser from 'koa-bodyparser-ts';
import * as helmet from 'koa-helmet';
import * as passport from 'koa-passport';
import * as session from 'koa-session';

import { config } from './config';
import { Sentry } from './logging';
import { initParse } from './parse';
import { router } from './routes';

import * as os from 'os';

export const app = new Koa();

/**
 * Initialize Parse independently
 */
initParse();
/**
 * Passport Settings
 */
app.keys = [config.secret];
app.use(session(app));
app.use(passport.initialize());
app.use(passport.session());

/**
 * Global Headers
 */
app.use(async (ctx: Koa.Context, next) => {
    ctx.set('CVPM-DISCOVERY-SERVED-BY', os.hostname());
    ctx.set('CVPM-DISCOVERY-VERSION', config.version);
    await next();
});

console.log(config)

app.use(cors());
app.use(helmet());
app.use(bodyParser());
app.use(router.routes()).use(router.allowedMethods());
app.on('error', err => {
    Sentry.captureException(err);
});
app.listen(config.port);