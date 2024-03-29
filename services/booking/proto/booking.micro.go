// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: booking.proto

package booking

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Booking service

type BookingService interface {
	DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...client.CallOption) (*DeleteBookingResult, error)
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...client.CallOption) (*CreateBookingResult, error)
	ConfirmBooking(ctx context.Context, in *ConfirmBookingRequest, opts ...client.CallOption) (*ConfirmBookingResult, error)
	FromShowDelete(ctx context.Context, in *FromShowDeleteRequest, opts ...client.CallOption) (*FromShowDeleteResult, error)
	GetBookingList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResult, error)
	GetNotConfirmedList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResult, error)
	Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResult, error)
}

type bookingService struct {
	c    client.Client
	name string
}

func NewBookingService(name string, c client.Client) BookingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "booking"
	}
	return &bookingService{
		c:    c,
		name: name,
	}
}

func (c *bookingService) DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...client.CallOption) (*DeleteBookingResult, error) {
	req := c.c.NewRequest(c.name, "Booking.DeleteBooking", in)
	out := new(DeleteBookingResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingService) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...client.CallOption) (*CreateBookingResult, error) {
	req := c.c.NewRequest(c.name, "Booking.CreateBooking", in)
	out := new(CreateBookingResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingService) ConfirmBooking(ctx context.Context, in *ConfirmBookingRequest, opts ...client.CallOption) (*ConfirmBookingResult, error) {
	req := c.c.NewRequest(c.name, "Booking.ConfirmBooking", in)
	out := new(ConfirmBookingResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingService) FromShowDelete(ctx context.Context, in *FromShowDeleteRequest, opts ...client.CallOption) (*FromShowDeleteResult, error) {
	req := c.c.NewRequest(c.name, "Booking.FromShowDelete", in)
	out := new(FromShowDeleteResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingService) GetBookingList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResult, error) {
	req := c.c.NewRequest(c.name, "Booking.GetBookingList", in)
	out := new(GetListResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingService) GetNotConfirmedList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResult, error) {
	req := c.c.NewRequest(c.name, "Booking.GetNotConfirmedList", in)
	out := new(GetListResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingService) Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResult, error) {
	req := c.c.NewRequest(c.name, "Booking.Exist", in)
	out := new(ExistResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Booking service

type BookingHandler interface {
	DeleteBooking(context.Context, *DeleteBookingRequest, *DeleteBookingResult) error
	CreateBooking(context.Context, *CreateBookingRequest, *CreateBookingResult) error
	ConfirmBooking(context.Context, *ConfirmBookingRequest, *ConfirmBookingResult) error
	FromShowDelete(context.Context, *FromShowDeleteRequest, *FromShowDeleteResult) error
	GetBookingList(context.Context, *GetListRequest, *GetListResult) error
	GetNotConfirmedList(context.Context, *GetListRequest, *GetListResult) error
	Exist(context.Context, *ExistRequest, *ExistResult) error
}

func RegisterBookingHandler(s server.Server, hdlr BookingHandler, opts ...server.HandlerOption) error {
	type booking interface {
		DeleteBooking(ctx context.Context, in *DeleteBookingRequest, out *DeleteBookingResult) error
		CreateBooking(ctx context.Context, in *CreateBookingRequest, out *CreateBookingResult) error
		ConfirmBooking(ctx context.Context, in *ConfirmBookingRequest, out *ConfirmBookingResult) error
		FromShowDelete(ctx context.Context, in *FromShowDeleteRequest, out *FromShowDeleteResult) error
		GetBookingList(ctx context.Context, in *GetListRequest, out *GetListResult) error
		GetNotConfirmedList(ctx context.Context, in *GetListRequest, out *GetListResult) error
		Exist(ctx context.Context, in *ExistRequest, out *ExistResult) error
	}
	type Booking struct {
		booking
	}
	h := &bookingHandler{hdlr}
	return s.Handle(s.NewHandler(&Booking{h}, opts...))
}

type bookingHandler struct {
	BookingHandler
}

func (h *bookingHandler) DeleteBooking(ctx context.Context, in *DeleteBookingRequest, out *DeleteBookingResult) error {
	return h.BookingHandler.DeleteBooking(ctx, in, out)
}

func (h *bookingHandler) CreateBooking(ctx context.Context, in *CreateBookingRequest, out *CreateBookingResult) error {
	return h.BookingHandler.CreateBooking(ctx, in, out)
}

func (h *bookingHandler) ConfirmBooking(ctx context.Context, in *ConfirmBookingRequest, out *ConfirmBookingResult) error {
	return h.BookingHandler.ConfirmBooking(ctx, in, out)
}

func (h *bookingHandler) FromShowDelete(ctx context.Context, in *FromShowDeleteRequest, out *FromShowDeleteResult) error {
	return h.BookingHandler.FromShowDelete(ctx, in, out)
}

func (h *bookingHandler) GetBookingList(ctx context.Context, in *GetListRequest, out *GetListResult) error {
	return h.BookingHandler.GetBookingList(ctx, in, out)
}

func (h *bookingHandler) GetNotConfirmedList(ctx context.Context, in *GetListRequest, out *GetListResult) error {
	return h.BookingHandler.GetNotConfirmedList(ctx, in, out)
}

func (h *bookingHandler) Exist(ctx context.Context, in *ExistRequest, out *ExistResult) error {
	return h.BookingHandler.Exist(ctx, in, out)
}
