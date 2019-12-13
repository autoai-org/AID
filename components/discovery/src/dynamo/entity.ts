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
    public id: string;

    @rangeKey({defaultProvider: () => new Date()})
    public createdAt: Date;

    @attribute()
    public isSymbol: boolean;

    @attribute()
    public linkedTo: string;
}

@table('cvpm-pretrained')
class Pretrained {
    @hashKey()
    public id: string;

    @rangeKey({defaultProvider: () => new Date()})
    public createdAt: Date;

    @attribute()
    public linkedTo: string;

    @attribute()
    public name: string;
}

@table('cvpm-registry')
class Registry {
    @hashKey()
    public id: string;

    @rangeKey({defaultProvider: () => new Date()})
    public createdAt: Date;

    @attribute()
    public name: string;

    @attribute()
    public urlPrefix: string;
}

export {
    Package,
    Pretrained,
    Registry
};