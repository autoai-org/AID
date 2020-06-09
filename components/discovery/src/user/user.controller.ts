import { Controller, Post, Body } from '@nestjs/common'
import { CreateUserDTO } from './user.schema'
import { UserService } from './user.service';
import { ApiTags } from '@nestjs/swagger';
import * as bcrypt from 'bcrypt';

@ApiTags('User')
@Controller('user')
export class UserController {
    constructor(private readonly userService: UserService) {}

    @Post()
    async create(@Body() createDto: CreateUserDTO) {
      createDto.password = await bcrypt.hash(createDto.password, 10);
      return this.userService.create(createDto)
    }
}