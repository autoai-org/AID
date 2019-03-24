import * as dotenv from 'dotenv';

dotenv.config({ path: '.env' });

export interface IConfig {
    blogURL: string;
    cognitoPoolId: string;
    cognitoClientId: string;
    cognitoRegion: string;
    debugLogging: boolean;
    port: number;
    secret: string;
    version: string;
    stripe: string;
    sentryDSN: string;
}

const config: IConfig = {
    blogURL: process.env.BLOG_URL || 'https://write.as/api/collections/autoai',
    cognitoPoolId: process.env.COGNITO_POOL_ID || '',
    cognitoClientId: process.env.COGNITO_CLIENT_ID || '',
    cognitoRegion: process.env.COGNITO_REGION || '',
    debugLogging: process.env.NODE_ENV === 'development',
    port: +process.env.PORT || 3000,
    secret: process.env.secret || 'your-secret-whatever',
    version: process.env.VERSION || 'v0.0.3@alpha',
    stripe: process.env.STRIPEID || '',
    sentryDSN: process.env.SENTRYDSN || '',
};

export { config };