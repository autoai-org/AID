import {Repository, RepositoryModel} from '../schemas/repository'
import { Resolver, Query } from 'type-graphql';

@Resolver(of => Repository)
export default class RepositoryResolver{
    @Query(returns => [Repository], {nullable: true})
    async repositories(): Promise<Repository[]> {
        return await RepositoryModel.find({});
    }
}