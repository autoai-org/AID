import { Injectable } from '@nestjs/common';
import { UserService } from '../user/user.service';
import * as bcrypt from 'bcrypt';
import { JwtService } from '@nestjs/jwt';
@Injectable()
export class AuthService {
  constructor(
      private usersService: UserService,
      private jwtService: JwtService
      ) 
    {}

  async validateUser(username: string, password: string): Promise<any> {
    const user = await this.usersService.findByUsername(username);
    console.log(await bcrypt.hash(password, 10))
    console.log(await bcrypt.compare(password, user.password))
    console.log(user.password)
    console.log(password)
    if (user && await bcrypt.compare(password, user.password)) {
      const { ...result } = user;
      return result;
    }
    return null;
  }
  async login(user: any) {
    const payload = { username: user.username, sub: user.userId };
    return {
      'access_token': this.jwtService.sign(payload),
    };
  }
}
