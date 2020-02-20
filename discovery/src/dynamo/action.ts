import { Guid } from 'guid-typescript';
import { isSymbol } from 'util';
import mapper from './dynamo';
import { Package, Pretrained, Registry } from './entity';

/**
 * Following will be package manipulations
 */

/**
 *
 * @param isSymbol
 * @param linkedTo
 */
function putPackage (isSymbol: boolean, linkedTo: string) {
    const toPutPackage = Object.assign(new Package, {
        id: Guid.create(),
        isSymbol,
        linkedTo
    });
    mapper.put(toPutPackage).then(objectSaved => {
    });
}
/**
 *
 * Following will be pretrained manipulations
 */

/**
 *
 * @param linkedTo
 * @param name
 */
function putPretrained (name: string, linkedTo: string) {
    const toPutPretrained = Object.assign(new Pretrained, {
        id: Guid.create(),
        name,
        linkedTo
    });
    mapper.put(toPutPretrained).then(objectSaved => {
    });
}

/**
 *
 */
async function getPretrained () {
    const results: Pretrained[] = [];
    for await (const item of mapper.scan(Pretrained)) {
        results.push(item);
    }
    return results;
}

/**
 * Following will be registry related operation
 */
async function getRegistries () {
    const results: Registry[] = [];
    for await (const item of mapper.scan(Registry)) {
        results.push(item);
    }
    return results;
}

async function addRegistry (name: string, urlPrefix: string) {
    const toAddRegistry = Object.assign(new Registry, {
        id: Guid.create(),
        name,
        urlPrefix,
    });
    mapper.put(toAddRegistry).then(objectSaved => {
    });
}

export {
    putPretrained,
    putPackage,
    getPretrained,
    addRegistry,
    getRegistries
};