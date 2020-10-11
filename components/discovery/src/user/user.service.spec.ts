import { Test } from '@nestjs/testing';
import { UserService } from './user.service';
import { UserController } from './user.controller';
import { DatabaseModule } from '../database/database.module'
import { UserProvider } from './user.provider'
describe('UserService', () => {
  let service: UserService;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      imports:[DatabaseModule],
      controllers: [UserController],
      providers: [UserService, ...UserProvider],
    }).compile();

    service = await moduleRef.resolve(UserService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
