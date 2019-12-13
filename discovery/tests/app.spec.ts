import * as supertest from 'supertest';
import { app } from '../src/server';

const server = app.listen();

const request = supertest.agent(server);

describe('ROOT', () => {
    describe('GET /', () => {
        it('should return meta info', () => {
            return request.get('/').expect(200);
        });
    });
    describe('GET /system/status', () => {
        it('should return system status', () => {
            return request.get('/system/status').expect(200);
        });
    });
    server.close();
});