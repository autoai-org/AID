import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { ModelsModule } from './model/model.module';
import config from './config';

console.info("[db]:connecting to " + config.MONGO_URI)

@Module({
  imports: [MongooseModule.forRoot(config.MONGO_URI),
    ModelsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
