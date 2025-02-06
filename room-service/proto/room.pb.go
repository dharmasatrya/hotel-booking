// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: room.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRoomByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        uint32                 `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoomByIdRequest) Reset() {
	*x = GetRoomByIdRequest{}
	mi := &file_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomByIdRequest) ProtoMessage() {}

func (x *GetRoomByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomByIdRequest.ProtoReflect.Descriptor instead.
func (*GetRoomByIdRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoomByIdRequest) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type Room struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Capacity      uint32                 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Price         uint32                 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	IsAvailable   bool                   `protobuf:"varint,5,opt,name=isAvailable,proto3" json:"isAvailable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_room_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *Room) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetCapacity() uint32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Room) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Room) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type GetAvailableRoomsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rooms         []*Room                `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAvailableRoomsResponse) Reset() {
	*x = GetAvailableRoomsResponse{}
	mi := &file_room_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvailableRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableRoomsResponse) ProtoMessage() {}

func (x *GetAvailableRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetAvailableRoomsResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *GetAvailableRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity      uint32                 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Price         uint32                 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_room_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoomRequest) GetCapacity() uint32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *CreateRoomRequest) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type EditRoomByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Capacity      uint32                 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Price         uint32                 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	IsAvailable   bool                   `protobuf:"varint,5,opt,name=isAvailable,proto3" json:"isAvailable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditRoomByIdRequest) Reset() {
	*x = EditRoomByIdRequest{}
	mi := &file_room_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditRoomByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRoomByIdRequest) ProtoMessage() {}

func (x *EditRoomByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRoomByIdRequest.ProtoReflect.Descriptor instead.
func (*EditRoomByIdRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

func (x *EditRoomByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditRoomByIdRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditRoomByIdRequest) GetCapacity() uint32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *EditRoomByIdRequest) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *EditRoomByIdRequest) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type DeleteRoomByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        uint32                 `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRoomByIdRequest) Reset() {
	*x = DeleteRoomByIdRequest{}
	mi := &file_room_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoomByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomByIdRequest) ProtoMessage() {}

func (x *DeleteRoomByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoomByIdRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRoomByIdRequest) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22,
	0x7e, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x3f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73,
	0x22, 0x59, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x13,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x2f, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x32, 0xc7, 0x02, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x39, 0x0a, 0x0c,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_room_proto_goTypes = []any{
	(*GetRoomByIdRequest)(nil),        // 0: userpb.GetRoomByIdRequest
	(*Room)(nil),                      // 1: userpb.Room
	(*GetAvailableRoomsResponse)(nil), // 2: userpb.GetAvailableRoomsResponse
	(*CreateRoomRequest)(nil),         // 3: userpb.CreateRoomRequest
	(*EditRoomByIdRequest)(nil),       // 4: userpb.EditRoomByIdRequest
	(*DeleteRoomByIdRequest)(nil),     // 5: userpb.DeleteRoomByIdRequest
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_room_proto_depIdxs = []int32{
	1, // 0: userpb.GetAvailableRoomsResponse.rooms:type_name -> userpb.Room
	0, // 1: userpb.UserService.GetRoomById:input_type -> userpb.GetRoomByIdRequest
	6, // 2: userpb.UserService.GetAvailableRooms:input_type -> google.protobuf.Empty
	3, // 3: userpb.UserService.CreateRoom:input_type -> userpb.CreateRoomRequest
	4, // 4: userpb.UserService.EditRoomById:input_type -> userpb.EditRoomByIdRequest
	5, // 5: userpb.UserService.DeleteRoomById:input_type -> userpb.DeleteRoomByIdRequest
	1, // 6: userpb.UserService.GetRoomById:output_type -> userpb.Room
	2, // 7: userpb.UserService.GetAvailableRooms:output_type -> userpb.GetAvailableRoomsResponse
	1, // 8: userpb.UserService.CreateRoom:output_type -> userpb.Room
	1, // 9: userpb.UserService.EditRoomById:output_type -> userpb.Room
	1, // 10: userpb.UserService.DeleteRoomById:output_type -> userpb.Room
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}
