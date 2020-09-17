import { Controller, Get } from '@nestjs/common'
import { ApiTags } from '@nestjs/swagger'
import { SystemService } from './systems.service'
import { UserService } from '../user/user.service'
import { ModelService } from '../model/model.service'

@ApiTags('Model')
@Controller('System')
export class SystemController {
    constructor(
        private readonly userService: UserService,
        private readonly modelService: ModelService,
        private readonly systemService: SystemService
    ) {

    }

    @Get('stats')
    async fetchStats(): Promise<object> {
        return {
            'user_count': await this.userService.count(),
            'model_count': await this.modelService.count()
        }
    }
}