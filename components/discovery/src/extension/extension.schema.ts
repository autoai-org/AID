import * as mongoose from 'mongoose'

interface AIDExtension extends mongoose.Document {
    entrypoint: string,
    name: string,
    vendor: string,
}