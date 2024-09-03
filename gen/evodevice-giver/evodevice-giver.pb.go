// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: evodevice-giver.proto

package evodevice_giver

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

type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evodevice_giver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evodevice_giver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_evodevice_giver_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evodevice_giver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evodevice_giver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_evodevice_giver_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            *string   `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	StoreId         string    `protobuf:"bytes,3,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	TimezoneOffset  *int64    `protobuf:"varint,5,opt,name=timezone_offset,json=timezoneOffset,proto3,oneof" json:"timezone_offset,omitempty"`
	Imei            *string   `protobuf:"bytes,6,opt,name=imei,proto3,oneof" json:"imei,omitempty"`
	FirmwareVersion *string   `protobuf:"bytes,7,opt,name=firmware_version,json=firmwareVersion,proto3,oneof" json:"firmware_version,omitempty"`
	Location        *Location `protobuf:"bytes,8,opt,name=location,proto3,oneof" json:"location,omitempty"`
	UserId          string    `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SerialNumber    *string   `protobuf:"bytes,10,opt,name=serial_number,json=serialNumber,proto3,oneof" json:"serial_number,omitempty"`
	DeviceModel     *string   `protobuf:"bytes,11,opt,name=device_model,json=deviceModel,proto3,oneof" json:"device_model,omitempty"`
	CreatedAt       int64     `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       int64     `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evodevice_giver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_evodevice_giver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_evodevice_giver_proto_rawDescGZIP(), []int{2}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Device) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Device) GetTimezoneOffset() int64 {
	if x != nil && x.TimezoneOffset != nil {
		return *x.TimezoneOffset
	}
	return 0
}

func (x *Device) GetImei() string {
	if x != nil && x.Imei != nil {
		return *x.Imei
	}
	return ""
}

func (x *Device) GetFirmwareVersion() string {
	if x != nil && x.FirmwareVersion != nil {
		return *x.FirmwareVersion
	}
	return ""
}

func (x *Device) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Device) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Device) GetSerialNumber() string {
	if x != nil && x.SerialNumber != nil {
		return *x.SerialNumber
	}
	return ""
}

func (x *Device) GetDeviceModel() string {
	if x != nil && x.DeviceModel != nil {
		return *x.DeviceModel
	}
	return ""
}

func (x *Device) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Device) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lng float32 `protobuf:"fixed32,1,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evodevice_giver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_evodevice_giver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_evodevice_giver_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *Location) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

var File_evodevice_giver_proto protoreflect.FileDescriptor

var file_evodevice_giver_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x76, 0x6f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x25, 0x0a,
	0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x88,
	0x04, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x0f, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x69,
	0x6d, 0x65, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x69, 0x6d, 0x65,
	0x69, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x0f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06,
	0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x69, 0x6d, 0x65, 0x69, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x2e, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x32, 0x41, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x76, 0x6f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evodevice_giver_proto_rawDescOnce sync.Once
	file_evodevice_giver_proto_rawDescData = file_evodevice_giver_proto_rawDesc
)

func file_evodevice_giver_proto_rawDescGZIP() []byte {
	file_evodevice_giver_proto_rawDescOnce.Do(func() {
		file_evodevice_giver_proto_rawDescData = protoimpl.X.CompressGZIP(file_evodevice_giver_proto_rawDescData)
	})
	return file_evodevice_giver_proto_rawDescData
}

var file_evodevice_giver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_evodevice_giver_proto_goTypes = []interface{}{
	(*DeviceRequest)(nil),  // 0: auth.DeviceRequest
	(*DeviceResponse)(nil), // 1: auth.DeviceResponse
	(*Device)(nil),         // 2: auth.Device
	(*Location)(nil),       // 3: auth.Location
}
var file_evodevice_giver_proto_depIdxs = []int32{
	2, // 0: auth.DeviceResponse.devices:type_name -> auth.Device
	3, // 1: auth.Device.location:type_name -> auth.Location
	0, // 2: auth.GetStores.GetList:input_type -> auth.DeviceRequest
	1, // 3: auth.GetStores.GetList:output_type -> auth.DeviceResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_evodevice_giver_proto_init() }
func file_evodevice_giver_proto_init() {
	if File_evodevice_giver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evodevice_giver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRequest); i {
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
		file_evodevice_giver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceResponse); i {
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
		file_evodevice_giver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_evodevice_giver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
	file_evodevice_giver_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_evodevice_giver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evodevice_giver_proto_goTypes,
		DependencyIndexes: file_evodevice_giver_proto_depIdxs,
		MessageInfos:      file_evodevice_giver_proto_msgTypes,
	}.Build()
	File_evodevice_giver_proto = out.File
	file_evodevice_giver_proto_rawDesc = nil
	file_evodevice_giver_proto_goTypes = nil
	file_evodevice_giver_proto_depIdxs = nil
}
