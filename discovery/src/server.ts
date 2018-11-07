import * as cors from '@koa/cors';
import * as Koa from 'koa';
import * as bodyParser from 'koa-bodyparser';
import * as helmet from 'koa-helmet';
import * as winston from 'winston';

import { config } from './config';
import { router } from './routes';
import { logger } from './logging';

const app = new Koa();

app.use(cors());
app.use(helmet());
app.use(logger(winston))
app.use(bodyParser());
app.use(router.routes()).use(router.allowedMethods());

app.listen(config.port);
console.log(`Server running on port ${config.port}`);
