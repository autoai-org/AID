import * as dotenv from 'dotenv';

dotenv.config({ path: '.env' });

export interface IConfig {
    port: number;
    debugLogging: boolean;
    secret: string;
    version: string;
    blogURL: string;
    cognitoPoolId: string;
    cognitoClientId: string;
    cognitoRegion: string;
    stripe: string;
    sentryDSN: string;
}

const config: IConfig = {
    port: +process.env.PORT || 3000,
    debugLogging: process.env.NODE_ENV == 'development',
    secret: process.env.secret || 'your-secret-whatever',
    version: process.env.VERSION || 'v0.0.3@alpha',
    blogURL: process.env.BLOG_URL || 'https://write.as/api/collections/autoai',
    cognitoPoolId: process.env.COGNITO_POOL_ID || '',
    cognitoClientId: process.env.COGNITO_CLIENT_ID || '',
    cognitoRegion: process.env.COGNITO_REGION || '',
    stripe: process.env.STRIPEID || '',
    sentryDSN: process.env.SENTRYDSN || '',
};

export { config };