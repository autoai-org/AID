import * as Router from 'koa-router';

import controller = require('./controller');

const router = new Router();

router.get('/system/status', controller.system.getSystemStatus);

router.post('/packages', controller.package.importPackage);

router.put('/pretrained', controller.pretrained.importPretrained);
router.get('/pretrained', controller.pretrained.getAllPretrained);

router.put('/registry', controller.registry.putRegistry);

router.get('/news', controller.news.queryNews);

export { router };
