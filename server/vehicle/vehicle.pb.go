// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: vehicle.proto

package vehicle

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{0}
}

type VehicleId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId int32 `protobuf:"varint,1,opt,name=vehicleId,proto3" json:"vehicleId,omitempty"`
}

func (x *VehicleId) Reset() {
	*x = VehicleId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleId) ProtoMessage() {}

func (x *VehicleId) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleId.ProtoReflect.Descriptor instead.
func (*VehicleId) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleId) GetVehicleId() int32 {
	if x != nil {
		return x.VehicleId
	}
	return 0
}

type VehicleRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand string `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Year  int32  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *VehicleRegistration) Reset() {
	*x = VehicleRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleRegistration) ProtoMessage() {}

func (x *VehicleRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleRegistration.ProtoReflect.Descriptor instead.
func (*VehicleRegistration) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleRegistration) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *VehicleRegistration) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *VehicleRegistration) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Year  int32  `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{3}
}

func (x *Vehicle) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vehicle) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Vehicle) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Vehicle) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type VehiclesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicles []*Vehicle `protobuf:"bytes,1,rep,name=vehicles,proto3" json:"vehicles,omitempty"`
}

func (x *VehiclesList) Reset() {
	*x = VehiclesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehiclesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehiclesList) ProtoMessage() {}

func (x *VehiclesList) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehiclesList.ProtoReflect.Descriptor instead.
func (*VehiclesList) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{4}
}

func (x *VehiclesList) GetVehicles() []*Vehicle {
	if x != nil {
		return x.Vehicles
	}
	return nil
}

var File_vehicle_proto protoreflect.FileDescriptor

var file_vehicle_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x29, 0x0a, 0x09, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x22, 0x59, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x3c, 0x0a,
	0x0c, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x08, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x08, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x32, 0xa3, 0x02, 0x0a, 0x0e,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1c,
	0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x33,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x10, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x1a, 0x10, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x12, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x15, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x12,
	0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicle_proto_rawDescOnce sync.Once
	file_vehicle_proto_rawDescData = file_vehicle_proto_rawDesc
)

func file_vehicle_proto_rawDescGZIP() []byte {
	file_vehicle_proto_rawDescOnce.Do(func() {
		file_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicle_proto_rawDescData)
	})
	return file_vehicle_proto_rawDescData
}

var file_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vehicle_proto_goTypes = []interface{}{
	(*Void)(nil),                // 0: vehicle.Void
	(*VehicleId)(nil),           // 1: vehicle.VehicleId
	(*VehicleRegistration)(nil), // 2: vehicle.VehicleRegistration
	(*Vehicle)(nil),             // 3: vehicle.Vehicle
	(*VehiclesList)(nil),        // 4: vehicle.VehiclesList
}
var file_vehicle_proto_depIdxs = []int32{
	3, // 0: vehicle.VehiclesList.vehicles:type_name -> vehicle.Vehicle
	2, // 1: vehicle.VehicleService.StoreVehicle:input_type -> vehicle.VehicleRegistration
	3, // 2: vehicle.VehicleService.UpdateVehicle:input_type -> vehicle.Vehicle
	1, // 3: vehicle.VehicleService.GetVehicle:input_type -> vehicle.VehicleId
	0, // 4: vehicle.VehicleService.ListVehicles:input_type -> vehicle.Void
	1, // 5: vehicle.VehicleService.DeleteVehicle:input_type -> vehicle.VehicleId
	3, // 6: vehicle.VehicleService.StoreVehicle:output_type -> vehicle.Vehicle
	3, // 7: vehicle.VehicleService.UpdateVehicle:output_type -> vehicle.Vehicle
	3, // 8: vehicle.VehicleService.GetVehicle:output_type -> vehicle.Vehicle
	4, // 9: vehicle.VehicleService.ListVehicles:output_type -> vehicle.VehiclesList
	0, // 10: vehicle.VehicleService.DeleteVehicle:output_type -> vehicle.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vehicle_proto_init() }
func file_vehicle_proto_init() {
	if File_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vehicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vehicle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleRegistration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vehicle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vehicle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehiclesList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vehicle_proto_goTypes,
		DependencyIndexes: file_vehicle_proto_depIdxs,
		MessageInfos:      file_vehicle_proto_msgTypes,
	}.Build()
	File_vehicle_proto = out.File
	file_vehicle_proto_rawDesc = nil
	file_vehicle_proto_goTypes = nil
	file_vehicle_proto_depIdxs = nil
}
