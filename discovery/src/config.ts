import * as dotenv from 'dotenv'

dotenv.config({ path: '.env' });

let isDev = true

export interface IConfig {
    port: number;
    debugLogging: boolean;
    dbsslconn: boolean;
    secret: string;
    version: string;
    blogURL: string;
}

const config: IConfig = {
    port: +process.env.PORT || 3000,
    debugLogging: process.env.NODE_ENV == 'development',
    dbsslconn: process.env.NODE_ENV != 'development',
    jwtSecret: process.env.JWT_SECRET || 'your-secret-whatever',
    version: process.env.VERSION || 'v0.0.3@alpha',
    blogURL: process.env.BLOG_URL || 'https://write.as/api/collections/autoai'
};

export { config };