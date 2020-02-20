import { Schema, Model, model, Document } from 'mongoose'

// interfaces
interface IToken {
    email: string;
    token: string;
    validUntil: number;
    createdAt: Date;
}

interface ITokenModel extends IToken, Document {
    isValid(): boolean;
    deactivate(): boolean;
}

let TokenSchema: Schema = new Schema({
    createdAt: Date,
    email: String,
    token: String,
    validUntil: Number,
})

TokenSchema.methods.isValid = function ():boolean {
    let current = Date.now()
    return (current < this.validUntil)
}

TokenSchema.methods.deactivate = function ():boolean {
    this.validUntil = 0
    this.save().catch(function(err:any) {
        console.error(err)
    })
    return true
}

export default model<ITokenModel>('Token', TokenSchema)