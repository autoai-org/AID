import * as mongoose from 'mongoose'
import { config } from '../../config'
function connect() {
    console.log('connecting database...')
    mongoose.connect(config.mongostring, {
        useNewUrlParser: true
    }).then(function (res) {
        console.log('connected!')
    }).catch(function (err) {
        console.error('cannot connect to database ' + config.mongostring)
        console.error(err)
    })
}

export {
    connect
}