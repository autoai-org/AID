import { Schema, Model, model, Document } from 'mongoose'

// interfaces
interface IToken {
    email: string;
    token: string;
    validUntil: number;
}

interface ITokenModel extends IToken, Document {
    isValid(): boolean;
}

let TokenSchema: Schema = new Schema({
    createdAt: Date,
    email: String,
    token: String,
    validUntil: Number
})

TokenSchema.methods.isValid = function ():boolean {
    let current = Date.now()
    return (current < this.validUntil)
}

export default model<ITokenModel>('Token', TokenSchema)