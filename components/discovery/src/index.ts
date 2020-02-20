import {DiscoveryApplication} from './application';
import {ApplicationConfig} from '@loopback/core';

export {DiscoveryApplication};

export async function main(options: ApplicationConfig = {}) {
  const app = new DiscoveryApplication(options);
  await app.boot();
  await app.start();

  const url = app.restServer.url;
  console.log(`Server is running at ${url}`);
  return app;
}
