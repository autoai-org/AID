import { AppService } from './app.service';
import { Controller, Get, Request, UseGuards, Post } from '@nestjs/common';
import { LocalAuthGuard } from './auth/guard/local-auth.guard';
import { AuthService } from './auth/auth.service';
import { ApiTags } from '@nestjs/swagger';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService, private readonly authService: AuthService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @ApiTags('Authentication')
  @UseGuards(LocalAuthGuard)
  @Post('auth/login')
  async login(@Request() req) {
    return this.authService.login(req.user);
  }
  @Get("version")
  getVersion(): string {
    return "0.0.1"
  }
}
