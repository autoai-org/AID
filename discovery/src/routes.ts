import * as Router from 'koa-router';

import controller = require('./controller');

const router = new Router();

router.get('/system/status', controller.system.getSystemStatus);
router.post('/packages', controller.package.importPackage)

export { router };
