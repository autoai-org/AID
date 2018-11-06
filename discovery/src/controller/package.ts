import { Context } from 'koa';
import { Package } from '../dynamo/entity';
import { importPackageRequest } from './entity'
export default class PackageController {
    public static async importPackage (ctx: Context) {
        const toImportPackage : Package = new Package();
        console.log(typeof(ctx.request.body));

        let ipr:importPackageRequest;
        Object.assign(ipr, ctx.request.body);
        toImportPackage.isSymbol = true;
        toImportPackage.linkedTo = ipr.linkedTo;
        console.log(toImportPackage)
    }
}
