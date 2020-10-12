import * as mongoose from 'mongoose'

interface AIDUser extends mongoose.Document {
    username: string,
    email: string,
    password: string,
    bio: string,
    homepage?: string,
    sponsorLink?: string
    avatar: string,
    admin: boolean,
}

class CreateUserDTO {
    username: string;
    password: string;
    bio: string;
    email: string;
    avatar: string;
    sponsorLink?: string;
    homepage?: string;
    admin: boolean
}

const UserSchema = new mongoose.Schema({
    avatar: String,
    username: String,
    email: String,
    password: String,
    bio: String,
    homepage: String,
    sponsorLink: String
})

export {
    AIDUser,
    CreateUserDTO,
    UserSchema
}