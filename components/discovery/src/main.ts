import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import config from './config'
import * as helmet from 'helmet'


async function bootstrap() {
  if (config.MONGO_URI==="") {
    console.error("[db]: connection string is undefined!")
  } else {
    console.info("[db]:connecting to " + config.MONGO_URI)
    const app = await NestFactory.create(AppModule);
    app.use(helmet());
    const options = new DocumentBuilder()
      .setTitle('AID Discovery')
      .setDescription('The AID Discovery API')
      .setVersion('1.0')
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
