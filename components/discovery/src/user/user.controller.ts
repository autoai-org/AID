import { Controller, Get, Post, Body, UseGuards, Req, Render } from '@nestjs/common'
import { CreateUserDTO } from './user.schema'
import { UserService } from './user.service';
import { ApiTags } from '@nestjs/swagger';
import * as bcrypt from 'bcrypt';
import { AuthGuard } from '@nestjs/passport';

@ApiTags('User')
@Controller('user')
export class UserController {
    constructor(private readonly userService: UserService) {}

    @Post()
    async create(@Body() createDto: CreateUserDTO) {
      const user = await this.userService.findByUsername(createDto.username);
      if (user) {
        return {
          'msg':"This username has been taken, please change to another one.",
          'code':"400"
        }
      } else{
        createDto.password = await bcrypt.hash(createDto.password, 10);
        return this.userService.create(createDto)  
      }
    }

    @Get('google')
    @UseGuards(AuthGuard('google'))
    async googleAuth(@Req() req) {}
  
    @Get('google/redirect')
    @UseGuards(AuthGuard('google'))
    @Render('social_login')
    googleAuthRedirect(@Req() req) {
      const user = this.userService.googleLogin(req)
      return { accessToken: user.user.accessToken, username: user.user.firstName};
    }
}