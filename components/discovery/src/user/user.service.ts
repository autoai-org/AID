import { Injectable, Inject} from '@nestjs/common';
import { CreateUserDTO, AIDUser } from './user.schema'
import { Model } from 'mongoose';

@Injectable()
export class UserService {

  constructor(@Inject('AID_User') private aidUserModel: Model<AIDUser>) { }

  async create(req: CreateUserDTO): Promise<AIDUser> {
    const createdUser = new this.aidUserModel(req);
    return createdUser.save();
  }

  async findByUsername(username: string): Promise<AIDUser | undefined> {
    return this.aidUserModel.findOne({username: username}).exec();
  }

  async count(): Promise<number> {
    return this.aidUserModel.estimatedDocumentCount().exec()
  }

  googleLogin(req) {
    if (!req.user) {
      return {
        'msg':'no user verified from google!'
      }
    }
    return {
      'user': req.user
    }
  }
}
