import { Module } from '@nestjs/common'
import { MongooseModule } from '@nestjs/mongoose';
import { TerminusModule } from '@nestjs/terminus';
import { SystemController } from './systems.controller'
import { SystemService } from './systems.service'
import { UserService } from '../user/user.service'
import { ModelService } from '../model/model.service'
import { ModelSchema } from '../model/model.schema'
import { UserSchema } from '../user/user.schema'

@Module({
  imports: [
      MongooseModule.forFeature([{ name: 'AIDModel', schema: ModelSchema }]), MongooseModule.forFeature([{ name: 'AIDUser', schema: UserSchema }]),
      TerminusModule
  ],
  controllers: [SystemController],
  providers: [SystemService, UserService, ModelService],
})

export class SystemsModule {}