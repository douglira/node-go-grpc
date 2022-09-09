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

  @Put()
  async update(@Body() vehicle: Vehicle): Promise<Vehicle> {
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
