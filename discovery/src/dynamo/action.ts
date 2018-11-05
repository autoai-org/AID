import mapper from './dynamo'
import { Package } from './model'
import { Guid } from 'guid-typescript';
/**
 * Following will be object manipulations
 */
function putPackage (isSymbol: boolean, linkedTo: string) {
    const toPutPackage = Object.assign(new Package, {
        id: Guid.create(),
        isSymbol: isSymbol,
        linkedTo: linkedTo
    })
    mapper.put(toPutPackage).then(objectSaved => {
        console.log(objectSaved)
    })
}