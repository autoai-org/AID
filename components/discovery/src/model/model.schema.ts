import * as mongoose from 'mongoose'

interface AIDModel extends mongoose.Document {
    avatar: string,
    name: string,
    vendor: string,
    description: string,
    githubURL: string,
    downloads: number,
    star: number,
}

class CreateModelDto {
    avatar: string;
    name: string;
    vendor: string;
    description: string;
    githubURL: string;
}

const ModelSchema = new mongoose.Schema({
    avatar: String,
    name: String,
    vendor: String,
    githubURL: String,
    description: String,
    downloads: Number,
    star: Number,
})

export {
    AIDModel,
    CreateModelDto,
    ModelSchema,
}