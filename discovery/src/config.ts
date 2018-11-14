import * as dotenv from 'dotenv';

dotenv.config({ path: '.env' });

export interface IConfig {
    port: number;
    debugLogging: boolean;
    dbsslconn: boolean;
    secret: string;
    version: string;
    cognitoPoolId: string;
    cognitoClientId: string;
    cognitoRegion: string;
}

const config: IConfig = {
    port: +process.env.PORT || 3000,
    debugLogging: process.env.NODE_ENV == 'development',
    dbsslconn: process.env.NODE_ENV != 'development',
    secret: process.env.secret || 'your-secret-whatever',
    version: process.env.VERSION || 'v0.0.3@alpha',
    cognitoPoolId: process.env.COGNITO_POOL_ID || '',
    cognitoClientId: process.env.COGNITO_CLIENT_ID || '',
    cognitoRegion: process.env.COGNITO_REGION || ''
};

export { config };