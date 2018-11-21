import * as cors from '@koa/cors';
import * as Koa from 'koa';
import bodyParser from 'koa-bodyparser-ts';
import * as helmet from 'koa-helmet';
import * as winston from 'winston';
import * as passport from 'koa-passport';
import * as session from 'koa-session';

import { initParse } from './parse';
import { config } from './config';
import { router } from './routes';
import { logger } from './logging';

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

app.use(cors());
app.use(helmet());
app.use(logger(winston));
app.use(bodyParser());
app.use(router.routes()).use(router.allowedMethods());

app.listen(config.port);
winston.log('info', `Server running on port ${config.port}`);