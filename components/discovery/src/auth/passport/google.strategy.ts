import { PassportStrategy } from '@nestjs/passport';
import { Injectable } from '@nestjs/common';
import { Strategy, VerifyCallback } from 'passport-google-oauth20';
import config from '../../config';
@Injectable()
export class GoogleStrategy extends PassportStrategy(Strategy) {

  constructor() {
    super({
      clientID: config.GOOGLE_CLIENT_ID,
      clientSecret: config.GOOGLE_SECRET,
      callbackURL: '/user/google/redirect',
      scope: ['email', 'profile'],
    });
  }

  async validate (accessToken: string, refreshToken: string, profile: any, done: VerifyCallback): Promise<any> {
    const { name, emails, photos } = profile
    const user = {
      email: emails[0].value,
      firstName: name.givenName,
      lastName: name.familyName,
      picture: photos[0].value,
      accessToken
    }
    done(null, user);
  }
}