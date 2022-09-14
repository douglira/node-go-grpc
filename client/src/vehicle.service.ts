import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Vehicle, StoreVehicleDto } from './vehicles.dto';

interface IVehicleClient {
  StoreVehicle(storeVehicleDto: StoreVehicleDto): Promise<Vehicle>;
  UpdateVehicle(vehicle: Vehicle): Promise<Vehicle>;
  GetVehicle(data: { vehicleId: number }): Promise<Vehicle>;
  ListVehicles(data: any): Promise<Vehicle[]>;
  DeleteVehicle(data: { vehicleId: number }): Promise<void>;
}
@Injectable()
export class VehicleService implements OnModuleInit {
  private vehicleClient: IVehicleClient;

  constructor(@Inject('VEHICLE_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.vehicleClient =
      this.client.getService<IVehicleClient>('VehicleService');
  }

  store(storeVehicleDto: StoreVehicleDto): Promise<Vehicle> {
    return this.vehicleClient.StoreVehicle(storeVehicleDto);
  }

  update(vehicle: Vehicle): Promise<Vehicle> {
    return this.vehicleClient.UpdateVehicle(vehicle);
  }

  findOne(data: { vehicleId: number }): Promise<Vehicle> {
    return this.vehicleClient.GetVehicle(data);
  }

  findAll(): Promise<Vehicle[]> {
    return this.vehicleClient.ListVehicles({});
  }

  remove(data: { vehicleId: number }): Promise<void> {
    return this.vehicleClient.DeleteVehicle(data);
  }
}
