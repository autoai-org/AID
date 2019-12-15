import * as dotenv from 'dotenv'

dotenv.config({ path: '.env' });

interface IConfig {
    blogURL: string;
    cognitoPoolId: string;
    cognitoClientId: string;
    cognitoRegion: string;
    debugLogging: boolean;
    port: number;
    mongostring: string;
    secret: string;
    version: string;
    stripe: string;
    sentryDSN: string;
    freshdeskDomain: string;
    freshdeskToken: string;
    githubId: string;
    githubSecret: string;
}

let config :IConfig
config = {
    blogURL: process.env.BLOG_URL || 'https://write.as/api/collections/autoai',
    cognitoPoolId: process.env.COGNITO_POOL_ID || '',
    cognitoClientId: process.env.COGNITO_CLIENT_ID || '',
    cognitoRegion: process.env.COGNITO_REGION || '',
    debugLogging: process.env.NODE_ENV === 'development',
    port: Number(process.env.PORT) || 3000,
    mongostring: process.env.MONGO_STRING_TEST || 'mongodb://192.168.1.17:27017/discovery?retryWrites=true',
    secret: process.env.secret || 'your-secret-whatever',
    version: process.env.VERSION || 'v0.0.3@alpha',
    stripe: process.env.STRIPEID || '',
    sentryDSN: process.env.SENTRYDSN || '',
    freshdeskDomain: process.env.FRESHDESKDOMAIN || '',
    freshdeskToken: process.env.FRESHDESKTOKEN || '',
    githubId: process.env.GITHUBID || '',
    githubSecret: process.env.GITHUBSECRET || ''
};

export default config