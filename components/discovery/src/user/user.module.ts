import { Module } from '@nestjs/common';
import { UserService } from './user.service';
import { UserController } from './user.controller'
import { UserProvider } from './user.provider'
import { DatabaseModule } from '../database/database.module';

@Module({
  imports: [DatabaseModule],
  controllers: [UserController],
  exports: [UserService],
  providers: [UserService, ...UserProvider],
})
export class UserModule { }