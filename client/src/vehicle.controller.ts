import { Controller, Get, Post, Put, Delete, Body } from '@nestjs/common';
import { StoreVehicleDto, Vehicle } from './vehicles.dto';

@Controller('vehicles')
export class VehicleController {
  @Post()
  store(@Body() storeVehicleDto: StoreVehicleDto): Vehicle {
    console.log(storeVehicleDto);
    const vehicle = new Vehicle();
    vehicle.id = Math.floor(Math.random() * 10);
    vehicle.brand = storeVehicleDto.brand;
    vehicle.model = storeVehicleDto.model;
    vehicle.year = storeVehicleDto.year;
    return vehicle;
  }
}
