import { Context } from 'koa';
import { putPretrained, getPretrained } from '../dynamo/action';
export default class PackageController {
    public static async importPretrained (ctx: Context) {
        putPretrained('yolo_tiny','https://premium.file.cvtron.xyz/data/yolo_tiny.ckpt');
        ctx.status = 200;
        ctx.body = {
            'code': '200',
            'desc': 'success'
        };
    }
    public static async getAllPretrained (ctx: Context) {
        ctx.status = 200;
        ctx.body = {
            'code': '200',
            'results': await getPretrained()
        }
    }
}
