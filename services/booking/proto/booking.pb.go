// Code generated by protoc-gen-go. DO NOT EDIT.
// source: booking.proto

package booking

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DeleteBookingRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookingRequest) Reset()         { *m = DeleteBookingRequest{} }
func (m *DeleteBookingRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBookingRequest) ProtoMessage()    {}
func (*DeleteBookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{0}
}

func (m *DeleteBookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookingRequest.Unmarshal(m, b)
}
func (m *DeleteBookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookingRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookingRequest.Merge(m, src)
}
func (m *DeleteBookingRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBookingRequest.Size(m)
}
func (m *DeleteBookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookingRequest proto.InternalMessageInfo

func (m *DeleteBookingRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteBookingResult struct {
	Successful           bool     `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookingResult) Reset()         { *m = DeleteBookingResult{} }
func (m *DeleteBookingResult) String() string { return proto.CompactTextString(m) }
func (*DeleteBookingResult) ProtoMessage()    {}
func (*DeleteBookingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{1}
}

func (m *DeleteBookingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookingResult.Unmarshal(m, b)
}
func (m *DeleteBookingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookingResult.Marshal(b, m, deterministic)
}
func (m *DeleteBookingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookingResult.Merge(m, src)
}
func (m *DeleteBookingResult) XXX_Size() int {
	return xxx_messageInfo_DeleteBookingResult.Size(m)
}
func (m *DeleteBookingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookingResult.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookingResult proto.InternalMessageInfo

func (m *DeleteBookingResult) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

type CreateBookingRequest struct {
	UserID               int32    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ShowID               int32    `protobuf:"varint,2,opt,name=showID,proto3" json:"showID,omitempty"`
	Seats                int32    `protobuf:"varint,3,opt,name=seats,proto3" json:"seats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookingRequest) Reset()         { *m = CreateBookingRequest{} }
func (m *CreateBookingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookingRequest) ProtoMessage()    {}
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{2}
}

func (m *CreateBookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookingRequest.Unmarshal(m, b)
}
func (m *CreateBookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookingRequest.Marshal(b, m, deterministic)
}
func (m *CreateBookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookingRequest.Merge(m, src)
}
func (m *CreateBookingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookingRequest.Size(m)
}
func (m *CreateBookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookingRequest proto.InternalMessageInfo

func (m *CreateBookingRequest) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CreateBookingRequest) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

func (m *CreateBookingRequest) GetSeats() int32 {
	if m != nil {
		return m.Seats
	}
	return 0
}

type CreateBookingResult struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookingResult) Reset()         { *m = CreateBookingResult{} }
func (m *CreateBookingResult) String() string { return proto.CompactTextString(m) }
func (*CreateBookingResult) ProtoMessage()    {}
func (*CreateBookingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{3}
}

func (m *CreateBookingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookingResult.Unmarshal(m, b)
}
func (m *CreateBookingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookingResult.Marshal(b, m, deterministic)
}
func (m *CreateBookingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookingResult.Merge(m, src)
}
func (m *CreateBookingResult) XXX_Size() int {
	return xxx_messageInfo_CreateBookingResult.Size(m)
}
func (m *CreateBookingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookingResult.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookingResult proto.InternalMessageInfo

func (m *CreateBookingResult) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ConfirmBookingRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmBookingRequest) Reset()         { *m = ConfirmBookingRequest{} }
func (m *ConfirmBookingRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmBookingRequest) ProtoMessage()    {}
func (*ConfirmBookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{4}
}

func (m *ConfirmBookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmBookingRequest.Unmarshal(m, b)
}
func (m *ConfirmBookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmBookingRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmBookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmBookingRequest.Merge(m, src)
}
func (m *ConfirmBookingRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmBookingRequest.Size(m)
}
func (m *ConfirmBookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmBookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmBookingRequest proto.InternalMessageInfo

func (m *ConfirmBookingRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ConfirmBookingResult struct {
	Successful           bool     `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmBookingResult) Reset()         { *m = ConfirmBookingResult{} }
func (m *ConfirmBookingResult) String() string { return proto.CompactTextString(m) }
func (*ConfirmBookingResult) ProtoMessage()    {}
func (*ConfirmBookingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{5}
}

func (m *ConfirmBookingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmBookingResult.Unmarshal(m, b)
}
func (m *ConfirmBookingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmBookingResult.Marshal(b, m, deterministic)
}
func (m *ConfirmBookingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmBookingResult.Merge(m, src)
}
func (m *ConfirmBookingResult) XXX_Size() int {
	return xxx_messageInfo_ConfirmBookingResult.Size(m)
}
func (m *ConfirmBookingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmBookingResult.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmBookingResult proto.InternalMessageInfo

func (m *ConfirmBookingResult) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

type BookingData struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID               int32    `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	ShowID               int32    `protobuf:"varint,3,opt,name=showID,proto3" json:"showID,omitempty"`
	Seats                int32    `protobuf:"varint,4,opt,name=seats,proto3" json:"seats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingData) Reset()         { *m = BookingData{} }
func (m *BookingData) String() string { return proto.CompactTextString(m) }
func (*BookingData) ProtoMessage()    {}
func (*BookingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b881a7066599f9d6, []int{6}
}

func (m *BookingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingData.Unmarshal(m, b)
}
func (m *BookingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingData.Marshal(b, m, deterministic)
}
func (m *BookingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingData.Merge(m, src)
}
func (m *BookingData) XXX_Size() int {
	return xxx_messageInfo_BookingData.Size(m)
}
func (m *BookingData) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingData.DiscardUnknown(m)
}

var xxx_messageInfo_BookingData proto.InternalMessageInfo

func (m *BookingData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BookingData) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *BookingData) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

func (m *BookingData) GetSeats() int32 {
	if m != nil {
		return m.Seats
	}
	return 0
}

func init() {
	proto.RegisterType((*DeleteBookingRequest)(nil), "booking.DeleteBookingRequest")
	proto.RegisterType((*DeleteBookingResult)(nil), "booking.DeleteBookingResult")
	proto.RegisterType((*CreateBookingRequest)(nil), "booking.CreateBookingRequest")
	proto.RegisterType((*CreateBookingResult)(nil), "booking.CreateBookingResult")
	proto.RegisterType((*ConfirmBookingRequest)(nil), "booking.ConfirmBookingRequest")
	proto.RegisterType((*ConfirmBookingResult)(nil), "booking.ConfirmBookingResult")
	proto.RegisterType((*BookingData)(nil), "booking.BookingData")
}

func init() { proto.RegisterFile("booking.proto", fileDescriptor_b881a7066599f9d6) }

var fileDescriptor_b881a7066599f9d6 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x6c, 0xd2, 0xaf, 0xed, 0xa7, 0x45, 0xed, 0xc1, 0x0d, 0x28, 0xaa, 0x68, 0x85, 0x2c, 0xf1,
	0x73, 0xea, 0x01, 0x04, 0x0f, 0x40, 0x73, 0xe9, 0x05, 0x89, 0x9c, 0xb9, 0xa4, 0xe9, 0x16, 0x22,
	0x42, 0x0d, 0x59, 0x5b, 0xbc, 0x3c, 0x07, 0x94, 0xd8, 0x08, 0xc7, 0x38, 0x82, 0xe3, 0x8c, 0x77,
	0x67, 0x77, 0x76, 0x0c, 0xe3, 0x8d, 0x10, 0xcf, 0xc5, 0xfe, 0x71, 0xf9, 0x5a, 0x09, 0x29, 0xd8,
	0xc8, 0x40, 0x7e, 0x06, 0x51, 0x82, 0x25, 0x4a, 0xbc, 0xd5, 0x44, 0x8a, 0x6f, 0x0a, 0x49, 0xb2,
	0x09, 0x84, 0xc5, 0x36, 0x0e, 0x4e, 0x82, 0x8b, 0x41, 0x1a, 0x16, 0x5b, 0x7e, 0x0d, 0x53, 0xa7,
	0x8e, 0x54, 0x29, 0xd9, 0x02, 0x80, 0x54, 0x9e, 0x23, 0xd1, 0x4e, 0x95, 0x4d, 0xf9, 0xff, 0xd4,
	0x62, 0xf8, 0x03, 0x44, 0xab, 0x0a, 0xb3, 0x1f, 0xf2, 0x47, 0x30, 0x54, 0x84, 0xd5, 0x3a, 0x31,
	0x23, 0x0c, 0xaa, 0x79, 0x7a, 0x12, 0xef, 0xeb, 0x24, 0x0e, 0x35, 0xaf, 0x11, 0x8b, 0x60, 0x40,
	0x98, 0x49, 0x8a, 0xfb, 0x0d, 0xad, 0x01, 0x3f, 0x85, 0xa9, 0xa3, 0xde, 0x2c, 0xe5, 0xee, 0x7e,
	0x0e, 0x87, 0x2b, 0xb1, 0xdf, 0x15, 0xd5, 0xcb, 0x2f, 0x26, 0x6f, 0x20, 0x72, 0x0b, 0xff, 0xe4,
	0x32, 0x87, 0x03, 0xd3, 0x90, 0x64, 0x32, 0x73, 0x65, 0x2d, 0xb3, 0x61, 0x87, 0xd9, 0xbe, 0xdf,
	0xec, 0x3f, 0xcb, 0xec, 0xe5, 0x47, 0x00, 0x23, 0x33, 0x85, 0xdd, 0xc1, 0xb8, 0x95, 0x06, 0x9b,
	0x2f, 0xbf, 0xf2, 0xf5, 0xa5, 0x39, 0x3b, 0xee, 0x7a, 0xae, 0xed, 0xf1, 0x5e, 0xad, 0xd7, 0x3a,
	0xa4, 0xa5, 0xe7, 0x8b, 0xcf, 0xd2, 0xf3, 0xdc, 0x9f, 0xf7, 0xd8, 0x3d, 0x4c, 0xda, 0x87, 0x64,
	0x8b, 0xef, 0x0e, 0x5f, 0x14, 0xb3, 0x79, 0xe7, 0xbb, 0x96, 0xdc, 0x0c, 0x9b, 0x8f, 0x7b, 0xf5,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x85, 0x86, 0xee, 0x4a, 0xc9, 0x02, 0x00, 0x00,
}