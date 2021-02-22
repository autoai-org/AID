import { Module } from '@nestjs/common';
import { ModelController } from './model.controller';
import { ModelService } from './model.service';
import { DatabaseModule } from '../database/database.module';
import { ModelProvider } from './model.provider'

@Module({
  imports: [DatabaseModule],
  controllers: [ModelController],
  exports: [ModelService],
  providers: [ModelService, ...ModelProvider,],
})
export class ModelModule {}