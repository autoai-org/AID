import { Connection } from 'mongoose';
import { UserSchema } from './user.schema';

export const UserProvider = [
  {
    provide: 'AID_User',
    useFactory: (connection: Connection) => connection.model('aidusers', UserSchema),
    inject: ['DATABASE_CONNECTION'],
  },
];