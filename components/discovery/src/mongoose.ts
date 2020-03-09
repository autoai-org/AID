import mongoose from 'mongoose';
import { Application } from './declarations';
import logger from './logger';

export default function (app: Application) {
  console.log(app.get('mongodb'))
  mongoose.connect(
    app.get('mongodb'),
    { useCreateIndex: true, useNewUrlParser: true, useUnifiedTopology: true }
  ).catch(err => {
    logger.error(err);
  });
  
  mongoose.Promise = global.Promise;

  app.set('mongooseClient', mongoose);
}
