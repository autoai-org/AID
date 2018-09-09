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

@table('cvpm-pretrained')
class Pretrained {
    @hashKey()
    id: string;

    @rangeKey({defaultProvider: () => new Date()})
    createdAt: Date;

    @attribute()
    linkedTo: string;

    @attribute()
    name: string;
}

@table('cvpm-registry')
class Registry {
    @hashKey()
    id: string;

    @rangeKey({defaultProvider: () => new Date()})
    createdAt: Date;

    @attribute()
    name: string;

    @attribute()
    urlPrefix: string;
}

export {
    Package,
    Pretrained,
    Registry
};