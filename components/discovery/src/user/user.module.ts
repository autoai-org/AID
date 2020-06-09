import { Module } from '@nestjs/common';
import { UserService } from './user.service';
import { MongooseModule } from '@nestjs/mongoose';
import { UserSchema } from './user.schema';
import { UserController } from './user.controller'

@Module({
  imports: [MongooseModule.forFeature([{ name: 'AIDUser', schema: UserSchema }])],
  providers: [UserService],
  exports:[UserService],
  controllers: [UserController],
})
export class UserModule {}