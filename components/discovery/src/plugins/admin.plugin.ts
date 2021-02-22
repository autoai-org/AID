import { INestApplication } from '@nestjs/common';
import AdminBro from 'admin-bro';
import AdminBroExpress from 'admin-bro-expressjs';
import AdminBroMongoose from '@admin-bro/mongoose'
import { AIDModel } from '../model/model.schema';
import config from '../config';
import mongoose from "mongoose";

export async function setupAdminPanel(app: INestApplication): Promise<void> {
    const mongooseDb = await mongoose.connect(config.MONGO_URI)
    AdminBro.registerAdapter(AdminBroMongoose)
    /** Create adminBro instance */
    const adminBro = new AdminBro({
        databases: [mongooseDb],
        resources: [],
        rootPath: '/admin',
        branding: {
            companyName: 'AutoAI.org',
        },
    });


    /** Create router */
    const router = AdminBroExpress.buildRouter(adminBro);

    /** Bind routing */
    app.use(adminBro.options.rootPath, router);
}