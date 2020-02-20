import { BaseContext } from 'koa';
import { Package } from '../dynamo/entity';
import { ImportPackageRequest } from './entity';
export default class PackageController {
    public static async importPackage (ctx: BaseContext) {
        const toImportPackage: Package = new Package();
    }
}
