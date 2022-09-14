import {
  Controller,
  Post,
  Body,
  Put,
  Param,
  Get,
  Delete,
} from '@nestjs/common';
import { VehicleService } from './vehicle.service';
import { StoreVehicleDto, Vehicle } from './vehicles.dto';

@Controller('vehicles')
export class VehicleController {
  constructor(private vehicleService: VehicleService) {}

  @Post()
  async store(@Body() storeVehicleDto: StoreVehicleDto): Promise<Vehicle> {
    return this.vehicleService.store(storeVehicleDto);
  }

  @Put(':id')
  async update(
    @Param('id') id: number,
    @Body() vehicleDto: StoreVehicleDto,
  ): Promise<Vehicle> {
    const vehicle = new Vehicle();
    vehicle.id = id;
    vehicle.brand = vehicleDto.brand;
    vehicle.model = vehicleDto.model;
    vehicle.year = vehicleDto.year;
    return this.vehicleService.update(vehicle);
  }

  @Get(':id')
  async findOne(@Param() params): Promise<Vehicle> {
    return this.vehicleService.findOne({ vehicleId: params.id });
  }

  @Get()
  async findAll(): Promise<Vehicle[]> {
    return this.vehicleService.findAll();
  }

  @Delete(':id')
  async remove(@Param() params): Promise<void> {
    return this.vehicleService.remove({ vehicleId: params.id });
  }
}
