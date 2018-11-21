const Parse = require('parse/node').Parse;

import { CountResult } from './entity';

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
        result[i].models = results[i]['models'];
        result[i].registries = results[i]['registries'];
        result[i].users = results[i]['users'];
    }
    return result;
}

export {
    initParse,
    updateCount,
    getCount
};