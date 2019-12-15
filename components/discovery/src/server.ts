import "reflect-metadata";
import { ApolloServer } from "apollo-server";
import { connect } from "mongoose";
import { ObjectId } from "mongodb";
import * as path from "path";
import { buildSchema } from 'type-graphql'

import { TypegooseMiddleware } from './middlewares/typegoose.mdw'
import resolvers from './resolvers'
import { ObjectIdScalar } from "./object-id.scalar";
import config from './config';

async function bootstrap() {
    try {
        const mongoose = await connect(config.mongostring)
        const schema = await buildSchema({
            resolvers: resolvers,
            emitSchemaFile: path.resolve(__dirname, "schema.gql"),
            globalMiddlewares: [TypegooseMiddleware],
            scalarsMap: [{ type: ObjectId, scalar: ObjectIdScalar }],
        })
        const server = new ApolloServer({ schema});
        const { url } = await server.listen(config.port);
        console.log(`Server is running, GraphQL Playground available at ${url}`);
    } catch (err) {
    console.error(err);
  }
}

bootstrap()