import { Test } from '@nestjs/testing';
import { UserService } from './user.service';
import { UserController } from './user.controller';
import { MongooseModule } from '@nestjs/mongoose';
import { UserSchema } from './user.schema';

describe('UserService', () => {
  let service: UserService;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      imports:[MongooseModule.forFeature([{ name: 'AIDUser', schema: UserSchema }])],
      controllers: [UserController],
      providers: [UserService],
    }).compile();

    service = await moduleRef.resolve(UserService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
