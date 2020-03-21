import * as feathersAuthentication from '@feathersjs/authentication';
// Don't remove this comment. It's needed to format import lines nicely.

const { authenticate } = feathersAuthentication.hooks;

import {hooks} from '@feathersjs/authentication-local/lib/';


export default {
  before: {
    all: [],
    find: [ authenticate('jwt'), hooks.protect("password") ],
    get: [ authenticate('jwt'), hooks.protect("password") ],
    create: [ hooks.hashPassword("password"), hooks.protect("password") ],
    update: [ hooks.hashPassword("password"), authenticate('jwt'), hooks.protect("password") ],
    patch: [  hooks.hashPassword("password"), authenticate('jwt'), hooks.protect("password") ],
    remove: [ authenticate('jwt'), hooks.protect("password") ]
  },

  after: {
    all: [ 
    ],
    find: [],
    get: [],
    create: [],
    update: [],
    patch: [],
    remove: []
  },

  error: {
    all: [],
    find: [],
    get: [],
    create: [],
    update: [],
    patch: [],
    remove: []
  }
};
