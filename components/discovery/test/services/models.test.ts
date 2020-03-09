import assert from 'assert';
import app from '../../src/app';

describe('\'models\' service', () => {
  it('registered the service', () => {
    const service = app.service('models');

    assert.ok(service, 'Registered the service');
  });
});
