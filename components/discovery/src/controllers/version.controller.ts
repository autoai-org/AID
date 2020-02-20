// Uncomment these imports to begin using these cool features!

// import {inject} from '@loopback/context';

import {get} from '@loopback/rest';

export class VersionController {
  @get('/version')
  hello(): string {
    return 'v1.0.0@dev';
  }
}
