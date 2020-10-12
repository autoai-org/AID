import { Module } from '@nestjs/common'
import { TerminusModule } from '@nestjs/terminus';
import { SystemController } from './systems.controller'
import { SystemService } from './systems.service'
import { UserService } from '../user/user.service'
import { ModelService } from '../model/model.service'
import { ModelProvider } from '../model/model.provider'
import { UserProvider } from '../user/user.provider'
import { DatabaseModule } from '../database/database.module';

@Module({
  imports: [
    DatabaseModule,
    TerminusModule
  ],
  controllers: [SystemController],
  providers: [SystemService, UserService, ModelService, ...ModelProvider, ...UserProvider],
})

export class SystemsModule { }