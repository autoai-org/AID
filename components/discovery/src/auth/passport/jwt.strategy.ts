// Copyright (c) 2020 Xiaozhe Yao & AutoAI.org
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { Strategy, ExtractJwt } from 'passport-jwt';
import { Injectable } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { AuthService } from '../auth.service';
import config from '../../config'

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(private readonly authService: AuthService) {
    super({
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
      passReqToCallback: true,
      secretOrKey: config.JWT_SECRET,
    })
  }
  async validate(payload:any) {
    return { userId: payload.sub, username: payload.username };
  } 
}