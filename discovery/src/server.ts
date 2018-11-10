import * as cors from '@koa/cors';
import * as Koa from 'koa';
import * as bodyParser from 'koa-bodyparser';
import * as helmet from 'koa-helmet';
import * as winston from 'winston';

import { config } from './config';
import { router } from './routes';
import { logger } from './logging';

import * as os from 'os';

const app = new Koa();

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