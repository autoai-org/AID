import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { ModelController } from './model.controller';
import { ModelsService } from './model.service';
import { ModelSchema } from './model.schema';

@Module({
  imports: [MongooseModule.forFeature([{ name: 'AIDModel', schema: ModelSchema }])],
  controllers: [ModelController],
  providers: [ModelsService],
})
export class ModelsModule {}