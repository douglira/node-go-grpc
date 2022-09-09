import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import { Vehicle, StoreVehicleDto } from './vehicles.dto';

interface VehicleClient {
  StoreVehicle(storeVehicleDto: StoreVehicleDto): Observable<Vehicle>;
  UpdateVehicle(vehicle: Vehicle): Observable<Vehicle>;
  GetVehicle(vehicleId: { vehicleId: number }): Observable<Vehicle>;
  ListVehicles(): Observable<Vehicle[]>;
  DeleteVehicle(vehicleId: { vehicleId: number }): Observable<void>;
}
@Injectable()
export class VehicleService implements OnModuleInit {
  private vehicleClient: VehicleClient;

  constructor(@Inject('VEHICLE_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.vehicleClient =
      this.client.getService<VehicleClient>('VehicleService');
  }

  store(storeVehicleDto: StoreVehicleDto): Observable<Vehicle> {
    return this.vehicleClient.StoreVehicle(storeVehicleDto);
  }
}
