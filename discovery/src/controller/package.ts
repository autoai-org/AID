import { BaseContext } from 'koa';
import { Package } from '../dynamo/model'

export default class PackageController {
    public static async importPackage (ctx: BaseContext) {
        const toImportPackage : Package = new Package();
        console.log(ctx.body)
        toImportPackage.isSymbol = true;
        toImportPackage.linkedTo = ctx.body.linkedTo;
        console.log(toImportPackage)
    }
}
