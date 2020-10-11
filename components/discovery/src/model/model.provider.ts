import { Connection } from 'mongoose';
import { ModelSchema } from './model.schema';

export const ModelProvider = [
  {
    provide: 'AID_Model',
    useFactory: (connection: Connection) => connection.model('aidmodels', ModelSchema),
    inject: ['DATABASE_CONNECTION'],
  },
];