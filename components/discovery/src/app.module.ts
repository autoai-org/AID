import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { ModelModule } from './model/model.module';
import { AuthModule } from './auth/auth.module';
import { UserModule } from './user/user.module';
import { SystemsModule } from './systems/systems.module'
import config from './config';

@Module({
  imports: [MongooseModule.forRoot(config.MONGO_URI),
    ModelModule,
    AuthModule,
    UserModule,
    SystemsModule],
  controllers: [AppController],
  providers: [AppService],
})

export class AppModule {}
