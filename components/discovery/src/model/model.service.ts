import { Model } from 'mongoose';
import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { CreateModelDto, AIDModel } from './model.schema';

@Injectable()
export class ModelsService {
  constructor(@InjectModel('AIDModel') private aidModelModel: Model<AIDModel>) {}

  async create(createModelDto: CreateModelDto): Promise<AIDModel> {
    const createdModel = new this.aidModelModel(createModelDto);
    return createdModel.save();
  }

  async findAll(): Promise<AIDModel[]> {
    return this.aidModelModel.find().exec();
  }
  async findByKeyword(keyword: string): Promise<AIDModel> {
      return this.aidModelModel.findOne({name: new RegExp('^'+keyword+'$', "i")}).exec();
  }
}