import { Controller, Get, Post, Body, UseGuards, Req } from '@nestjs/common'
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
      createDto.password = await bcrypt.hash(createDto.password, 10);
      return this.userService.create(createDto)
    }

    @Get('google')
    @UseGuards(AuthGuard('google'))
    async googleAuth(@Req() req) {}
  
    @Get('google/redirect')
    @UseGuards(AuthGuard('google'))
    googleAuthRedirect(@Req() req) {
      return this.userService.googleLogin(req)
    }
  
}