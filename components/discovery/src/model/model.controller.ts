import { Controller, Get, Post, Body, Param } from '@nestjs/common'
import { AIDModel, CreateModelDto } from './model.schema'
import { ModelService } from './model.service';
import { ApiTags } from '@nestjs/swagger';
import { fetchGithubSummary } from '../services/github'

@ApiTags('Model')
@Controller('model')
export class ModelController {
    constructor(private readonly modelService: ModelService) {}

    @Get()
    async findAll(): Promise<AIDModel[]>{
        return this.modelService.findAll()
    }
    @Get("/keyword/:keyword")
    async findByKeyword(@Param("keyword") keyword:string): Promise<AIDModel[]> {
        return this.modelService.findByKeyword(keyword)
    }
    @Get("/meta/:packageId")
    async fetchMetaInformation(@Param("packageId") packageId: string): Promise<any> {
        const model = await this.modelService.findById(packageId)
        if (model.githubURL!=="") {
            const infos = model.githubURL.split("/")
            return fetchGithubSummary(infos[infos.length-2], infos[infos.length-1])
        } else {
            return {}
        }
    }
    @Post()
    async create(@Body() createDto: CreateModelDto) {
      return this.modelService.create(createDto)
    }
}