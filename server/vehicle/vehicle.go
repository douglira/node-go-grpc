package vehicle

import (
	"fmt"

	context "golang.org/x/net/context"
)

var listOfCars []*Vehicle = []*Vehicle{}

type VehicleServer struct {
	UnimplementedVehicleServiceServer
}

func (v *VehicleServer) StoreVehicle(ctx context.Context, vehicleRegistration *VehicleRegistration) (*Vehicle, error) {
	vehicle := Vehicle{
		Id:    int32(len(listOfCars)) + 1,
		Brand: vehicleRegistration.GetBrand(),
		Model: vehicleRegistration.GetModel(),
		Year:  vehicleRegistration.GetYear(),
	}
	listOfCars = append(listOfCars, &vehicle)
	fmt.Printf("new vehicle added: %v", vehicle.String())
	return &vehicle, nil
}

func (v *VehicleServer) UpdateVehicle(ctx context.Context, vehicle *Vehicle) (*Vehicle, error) {
	listOfCars[vehicle.GetId()-1] = vehicle
	return vehicle, nil
}

func (v *VehicleServer) GetVehicle(ctx context.Context, vehicleId *VehicleId) (*Vehicle, error) {
	vehicle := &listOfCars[vehicleId.GetVehicleId()-1]
	return *vehicle, nil
}

// TODO list vehicles
func (v *VehicleServer) ListVehicles(ctx context.Context, void *Void) (*VehiclesList, error) {

	for _, c := range listOfCars {
		fmt.Print(c.String())
	}
	list := VehiclesList{
		Vehicles: listOfCars,
	}

	return &list, nil
}

// TODO delete vehicles
func (v *VehicleServer) DeleteVehicle(ctx context.Context, vehicleId *VehicleId) (*Void, error) {
	i := vehicleId.GetVehicleId() - 1
	copy(listOfCars[i:], listOfCars[i+1:])
	listOfCars[len(listOfCars)-1] = &Vehicle{}
	listOfCars = listOfCars[:len(listOfCars)-1]
	return &Void{}, nil
}
