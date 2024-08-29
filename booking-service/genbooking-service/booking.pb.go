// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: booking.proto

package genbooking_service

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

type UserBookingEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserBookingEmail) Reset() {
	*x = UserBookingEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBookingEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBookingEmail) ProtoMessage() {}

func (x *UserBookingEmail) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBookingEmail.ProtoReflect.Descriptor instead.
func (*UserBookingEmail) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *UserBookingEmail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PostBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	HotelId      string `protobuf:"bytes,2,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	RoomType     string `protobuf:"bytes,3,opt,name=roomType,proto3" json:"roomType,omitempty"`
	CheckInDate  string `protobuf:"bytes,4,opt,name=checkInDate,proto3" json:"checkInDate,omitempty"`
	CheckOutDate string `protobuf:"bytes,5,opt,name=checkOutDate,proto3" json:"checkOutDate,omitempty"`
}

func (x *PostBookingRequest) Reset() {
	*x = PostBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBookingRequest) ProtoMessage() {}

func (x *PostBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBookingRequest.ProtoReflect.Descriptor instead.
func (*PostBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *PostBookingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostBookingRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *PostBookingRequest) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *PostBookingRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *PostBookingRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	HotelId      string `protobuf:"bytes,2,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	RoomType     string `protobuf:"bytes,3,opt,name=roomType,proto3" json:"roomType,omitempty"`
	CheckInDate  string `protobuf:"bytes,4,opt,name=checkInDate,proto3" json:"checkInDate,omitempty"`
	CheckOutDate string `protobuf:"bytes,5,opt,name=checkOutDate,proto3" json:"checkOutDate,omitempty"`
	TotalAmount  int32  `protobuf:"varint,6,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	BookingId    string `protobuf:"bytes,7,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	Status       string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *BookingResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingResponse) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *BookingResponse) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *BookingResponse) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *BookingResponse) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *BookingResponse) GetTotalAmount() int32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *BookingResponse) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *BookingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PutBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckInDate  string `protobuf:"bytes,1,opt,name=checkInDate,proto3" json:"checkInDate,omitempty"`
	CheckOutDate string `protobuf:"bytes,2,opt,name=checkOutDate,proto3" json:"checkOutDate,omitempty"`
	Status       string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	BookingId    string `protobuf:"bytes,5,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
}

func (x *PutBookingRequest) Reset() {
	*x = PutBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutBookingRequest) ProtoMessage() {}

func (x *PutBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutBookingRequest.ProtoReflect.Descriptor instead.
func (*PutBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *PutBookingRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *PutBookingRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *PutBookingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PutBookingRequest) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId string `protobuf:"bytes,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *BookingRequest) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

type UserBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserBookingRequest) Reset() {
	*x = UserBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBookingRequest) ProtoMessage() {}

func (x *UserBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBookingRequest.ProtoReflect.Descriptor instead.
func (*UserBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *UserBookingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	BookingId string `protobuf:"bytes,2,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
}

func (x *DeleteBookingResponse) Reset() {
	*x = DeleteBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookingResponse) ProtoMessage() {}

func (x *DeleteBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookingResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteBookingResponse) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x22, 0xfd, 0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x32, 0xd2, 0x02, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0a, 0x50, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x50,
	0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x0f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_booking_proto_goTypes = []any{
	(*UserBookingEmail)(nil),      // 0: UserBookingEmail
	(*PostBookingRequest)(nil),    // 1: PostBookingRequest
	(*BookingResponse)(nil),       // 2: BookingResponse
	(*PutBookingRequest)(nil),     // 3: PutBookingRequest
	(*BookingRequest)(nil),        // 4: BookingRequest
	(*UserBookingRequest)(nil),    // 5: UserBookingRequest
	(*DeleteBookingResponse)(nil), // 6: DeleteBookingResponse
}
var file_booking_proto_depIdxs = []int32{
	1, // 0: BookingService.AddBooking:input_type -> PostBookingRequest
	4, // 1: BookingService.GetBooking:input_type -> BookingRequest
	3, // 2: BookingService.PutBooking:input_type -> PutBookingRequest
	4, // 3: BookingService.DelBooking:input_type -> BookingRequest
	5, // 4: BookingService.GetUserBooking:input_type -> UserBookingRequest
	5, // 5: BookingService.GetUserEmail:input_type -> UserBookingRequest
	2, // 6: BookingService.AddBooking:output_type -> BookingResponse
	2, // 7: BookingService.GetBooking:output_type -> BookingResponse
	2, // 8: BookingService.PutBooking:output_type -> BookingResponse
	6, // 9: BookingService.DelBooking:output_type -> DeleteBookingResponse
	2, // 10: BookingService.GetUserBooking:output_type -> BookingResponse
	0, // 11: BookingService.GetUserEmail:output_type -> UserBookingEmail
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserBookingEmail); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PostBookingRequest); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BookingResponse); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PutBookingRequest); i {
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
		file_booking_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BookingRequest); i {
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
		file_booking_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UserBookingRequest); i {
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
		file_booking_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBookingResponse); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
