import { Field, ObjectType } from "type-graphql"
import { Typegoose } from 'typegoose'

@ObjectType()
export class Repository extends Typegoose {
    @Field(type => String)
    id: string
    @Field(type => String)
    name: string
    @Field(type => String)
    vendor: string
    @Field(type => String)
    url: string
}

export const RepositoryModel = new Repository().getModelForClass(Repository)