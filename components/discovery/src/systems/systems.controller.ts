import { Controller, Get } from '@nestjs/common'
import { ApiTags } from '@nestjs/swagger'
import { SystemService } from './systems.service'
import { UserService } from '../user/user.service'
import { ModelService } from '../model/model.service'
import { HealthCheckService, DNSHealthIndicator, HealthCheck, MongooseHealthIndicator } from '@nestjs/terminus'
import config from '../config'
@ApiTags('Model')
@Controller('System')
export class SystemController {
    constructor(
        private readonly userService: UserService,
        private readonly modelService: ModelService,
        private readonly systemService: SystemService,
        private readonly health: HealthCheckService,
        private readonly dns: DNSHealthIndicator,
        private readonly mongo: MongooseHealthIndicator,
    ) {

    }

    @Get('stats')
    async fetchStats(): Promise<object> {
        return {
            'user_count': await this.userService.count(),
            'model_count': await this.modelService.count()
        }
    }
    
    @Get('health')
    @HealthCheck()
    check() {
        return this.health.check([
            () => this.dns.pingCheck('aid-docs', 'https://aid.autoai.org')
            //() => this.mongo.pingCheck('aid-database', { connection: config.MONGO_URI })
        ]);
    }
}