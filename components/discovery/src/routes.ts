import * as Router from 'koa-router'

import controller = require('./controller');

const router = new Router();

router.post('/user/auth/social', controller.user.socialAuthenticate);

router.put('/tickets', controller.ticket.createTicket);

router.get('/system/status', controller.system.getSystemStatus);
router.get('/system/stats', controller.system.getStatInfo);
router.put('/system/stats', controller.system.putStatInfo);

router.post('/packages', controller.package.importPackage);

router.put('/pretrained', controller.pretrained.importPretrained);
router.get('/pretrained', controller.pretrained.getAllPretrained);

router.put('/registry', controller.registry.putRegistry);
router.get('/registries', controller.registry.getRegistries);

router.post('/session', controller.user.validateAccessToken);

router.get('/news', controller.news.queryNews);

router.post('/paywall/customer', controller.paywall.createCustomer);
router.post('/paywall/charge', controller.paywall.createCharge);
router.post('/paywall/subscription', controller.paywall.createSubscription);
router.put('/paywall/token', controller.paywall.addAccessToken);
router.delete('/paywall/token', controller.paywall.deleteAccessToken);
router.get('/paywall/tokens', controller.paywall.queryAccessTokens);

router.get('/', controller.system.getMetaInfo);

export { router };
