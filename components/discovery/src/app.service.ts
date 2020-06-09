import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  getHello(): string {
    return 'This is AutoAI/AID Discovery Service!';
  }
}
