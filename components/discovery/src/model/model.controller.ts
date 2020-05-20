import { Controller, Get, Post, Body, Param } from '@nestjs/common'
import { AIDModel, CreateModelDto } from './model.schema'
import { ModelsService } from './model.service';
import { ApiTags } from '@nestjs/swagger';
import { fetch_github_summary } from '../services/github'

@ApiTags('Model')
@Controller('model')
export class ModelController {
    constructor(private readonly modelsService: ModelsService) {}

    @Get()
    async findAll(): Promise<AIDModel[]>{
        return this.modelsService.findAll()
    }
    @Get("/keyword/:keyword")
    async findByKeyword(@Param("keyword") keyword:string): Promise<AIDModel[]> {
        return this.modelsService.findByKeyword(keyword)
    }
    @Get("/meta/:packageId")
    async fetchMetaInformation(@Param("packageId") packageId: string): Promise<any> {
        let model = await this.modelsService.findById(packageId)
        if (model.githubURL!=="") {
            let infos = model.githubURL.split("/")
            return fetch_github_summary(infos[infos.length-2], infos[infos.length-1])
        } else {
            return {}
        }
    }
    @Post()
    async create(@Body() createDto: CreateModelDto) {
      return this.modelsService.create(createDto)
    }
}