import mapper from './dynamo'
import { Package } from './entity'
import { Guid } from 'guid-typescript';
/**
 * Following will be package manipulations
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
