import { Context } from 'koa';
import { Package } from '../dynamo/entity';
import { importPackageRequest } from './entity'
export default class PackageController {
    public static async importPackage (ctx: Context) {
        const toImportPackage:Package = new Package();
        
        console.log(toImportPackage)
    }
}
