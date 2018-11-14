import passport from 'passport';
import { config } from '../config';
var CognitoStrategy = require('passport-cognito')

passport.use(new CognitoStrategy({
    userPoolId: config.cognitoPoolId,
    clientId: config.cognitoClientId,
    region: config.cognitoRegion
},
    function (accessToken: string, idToken: string, refreshToken: string, user: any, cb: any) {
        process.nextTick(function() {
            cb(null, user);
        });
    }
));