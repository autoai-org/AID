import {
    attribute,
    hashKey,
    rangeKey,
    table,
} from '@aws/dynamodb-data-mapper-annotations';

/**
 * Following will be object definitions
 */
@table('cvpm-package')
class Package {
    @hashKey()
    id: string;

    @rangeKey({defaultProvider: () => new Date()})
    createdAt: Date;

    @attribute()
    isSymbol: boolean;

    @attribute()
    linkedTo: string;
}

export {
    Package
}