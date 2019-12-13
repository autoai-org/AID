const Parse = require('parse/node').Parse;

import { CountResult, PojoResult } from './entity';

function initParse() {
    Parse.initialize(process.env.PARSE_ID, process.env.PARSE_TOKEN, process.env.PARSE_MASTER_KEY);
    (Parse as any).serverURL = process.env.PARSE_URL;
}

async function updateCount(users: number, registries: number, models: number) {
    const Meta = Parse.Object.extend('Meta');
    const meta = new Meta();
    meta.set('user', users);
    meta.set('registry', registries);
    meta.set('model', models);
    const results: any[] = await meta.save();
    return results;
}

async function getCount(limit: number) {
    const Meta = Parse.Object.extend('Meta');
    const query = new Parse.Query(Meta);
    query.descending('createdAt');
    query.limit(limit);
    const results: any[] = await query.find();
    const result: CountResult[] = [];
    for (let i = 0; i < results.length; i++) {
        const pojo_result: PojoResult = results[i].toJSON();
        const each_result: CountResult = {
            user: pojo_result.user,
            model: pojo_result.model,
            registry: pojo_result.registry,
            updatedAt: pojo_result.updatedAt,
        };
        result.push(each_result);
    }
    return result;
}

export {
    initParse,
    updateCount,
    getCount
};