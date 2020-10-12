import mongoose from "mongoose";
import config from '../config';

export const databaseProviders = [
    {
        provide: 'DATABASE_CONNECTION',
        useFactory: (): Promise<typeof mongoose> =>
            mongoose.connect(config.MONGO_URI),
    },
];