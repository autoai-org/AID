import * as Router from 'koa-router';

import controller = require('./controller');

const router = new Router();

router.get('/', controller.system.getMetaInfo);

router.put('/tickets', controller.ticket.createTicket);

router.get('/system/status', controller.system.getSystemStatus);
router.get('/system/stats', controller.system.getStatInfo);
router.put('/system/stats', controller.system.putStatInfo);

router.post('/packages', controller.package.importPackage);

router.put('/pretrained', controller.pretrained.importPretrained);
router.get('/pretrained', controller.pretrained.getAllPretrained);

router.put('/registry', controller.registry.putRegistry);
router.get('/registries', controller.registry.getRegistries);

router.get('/news', controller.news.queryNews);

export { router };
