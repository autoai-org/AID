import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import config from './config'
import helmet from 'helmet';
import rateLimit from 'express-rate-limit';
import { join } from 'path';
import { NestExpressApplication } from '@nestjs/platform-express';

async function bootstrap() {
  if (config.MONGO_URI==="") {
    console.error("[db]: connection string is undefined!")
  } else {
    console.info("[db]:connecting to " + config.MONGO_URI)
    const app = await NestFactory.create<NestExpressApplication>(AppModule);
    app.use(helmet());
    app.use(
      rateLimit({
        windowMs: 60 * 1000, // 15 minutes
        max: 100, // limit each IP to 100 requests per windowMs
      }),
    );
    app.setBaseViewsDir(join(__dirname, '..', 'src', 'views'));
    app.setViewEngine('hbs');
    const options = new DocumentBuilder()
      .setTitle('AID Discovery')
      .setDescription('The AID Discovery API')
      .setVersion('0.0.1')
      .addTag('Model')
      .addTag('User')
      .addTag('Extension')
      .build();
    const document = SwaggerModule.createDocument(app, options);
    SwaggerModule.setup('api', app, document);
    console.info("[discovery]: will listen on port " + config.PORT)
    app.enableCors();
    await app.listen(config.PORT);  
  }
}
bootstrap();
