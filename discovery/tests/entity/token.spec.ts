import mongoose from 'mongoose'
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
})