import assert from 'assert';
import app from '../src/app';

describe('authentication', () => {
  it('registered the authentication service', () => {
    assert.ok(app.service('authentication'));
  });
  
});
