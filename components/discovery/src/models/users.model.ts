// users-model.js - A mongoose model
//
// See http://mongoosejs.com/docs/models.html
// for more of what you can do here.
import { Application } from '../declarations';
import mongoose from 'mongoose'
export default function (app: Application) {
  const modelName = 'users';
  const mongooseClient = app.get('mongooseClient');
  const schema = new mongooseClient.Schema({
    username: { type: String, index: true, unique: true, required: [true, 'can\'t be blank'], match: [/^[a-zA-Z0-9_-]+$/, 'is invalid'] },
    email: { type: String, index: true, unique: true, required: [true, 'can\'t be blank'], match: [/\S+@\S+\.\S+/, 'is invalid'] },
    password: String,
    bio: String,
    image: String,
    followingList: [mongoose.Schema.Types.ObjectId],
    googleId: { type: String },
    facebookId: { type: String },
    twitterId: { type: String },
    githubId: { type: String },

  }, {
    timestamps: true
  });

  // This is necessary to avoid model compilation errors in watch mode
  // see https://mongoosejs.com/docs/api/connection.html#connection_Connection-deleteModel
  if (mongooseClient.modelNames().includes(modelName)) {
    mongooseClient.deleteModel(modelName);
  }
  return mongooseClient.model(modelName, schema);
}
