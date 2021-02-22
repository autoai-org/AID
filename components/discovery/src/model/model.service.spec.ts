import { Test } from '@nestjs/testing';
import { ModelService } from './model.service';
import { ModelController } from './model.controller';
import { DatabaseModule } from '../database/database.module';
import { ModelProvider } from './model.provider'

describe('ModelService', () => {
  let service: ModelService;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      imports:[DatabaseModule],
      controllers: [ModelController],
      providers: [ModelService, ...ModelProvider,],
    }).compile();

    service = await moduleRef.resolve(ModelService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
