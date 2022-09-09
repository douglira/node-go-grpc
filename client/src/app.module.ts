import { Module } from '@nestjs/common';
import { VehicleController } from './vehicle.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'VEHICLE_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'vehicle',
          protoPath: join(__dirname, '../../proto/vehicle.proto'),
        },
      },
    ]),
  ],
  controllers: [VehicleController],
  providers: [],
})
export class AppModule {}
