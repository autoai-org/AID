import { Controller, Get, Post, Body, Param, Query } from '@nestjs/common'
import { AIDModel, CreateModelDto } from './model.schema'
import { ModelsService } from './model.service';
import { ApiTags } from '@nestjs/swagger';
@ApiTags('Model')
@Controller('model')
export class ModelController {
    constructor(private readonly modelsService: ModelsService) {}

    @Get()
    async findAll(): Promise<AIDModel[]>{
        return this.modelsService.findAll()
    }
    @Get(':keyword')
    async findByKeyword(@Param('keyword') keyword:string): Promise<AIDModel> {
        return this.modelsService.findByKeyword(keyword)
    }
    @Post()
    async create(@Body() createDto: CreateModelDto) {
      return this.modelsService.create(createDto)
    }
}