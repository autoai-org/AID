import * as mongoose from 'mongoose'
import { config } from '../../src/config';
import Token from '../../src/service/database/entity/token'

describe('Token Model', () => {
    beforeAll(async () => {
        await mongoose.connect(config.mongostring, {
            useNewUrlParser: true
        })
    })
    afterAll(async () => {
        mongoose.connection.close();
    })
    it('Should throw validation errors', () => {
        const token = new Token()
        expect(token.validate).toThrow()
    })
    // more tests
    // ref: https://medium.com/@tomanagle/strongly-typed-models-with-mongoose-and-typescript-7bc2f7197722
    it('should save a token', async () => {
        expect.assertions(3);
        const token = new Token({
            email: 'test@test.em',
            token: '1234',
            validUntil: 100000,
            createdAt: Date.now(),
        })
        const spy = jest.spyOn(token, 'save')
        token.save()
        expect(spy).toHaveBeenCalled()
        expect(token).toMatchObject({
            email: expect.any(String),
            token: expect.any(String),
            validUntil: expect.any(Number),
            createdAt: expect.any(Date),
        })
        expect(token.email).toBe('test@test.em')
    })
    it('should deactive a token', async () => {
        expect.assertions(1);
        const token = new Token({
            email: 'test@test.em',
            token: '1234',
            validUntil: 100000,
            createdAt: Date.now(),
        })
        token.deactivate()
        expect(token.validUntil).toBe(0)
        expect(token.isValid()).toBe(false)
    })
})