package main


import (
	"log"
	"context"

	"github.com/micro/go-micro"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
)

type BService struct {
	booking 	[]*booking.BookingData
	notConfirmed []*booking.BookingData
	nextID	int32
}

func (bs *BService) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest, rsp *booking.CreateBookingResult) error {
	givenID := bs.nextID
	bs.nextID++
	bs.notConfirmed = append(bs.notConfirmed, &booking.BookingData{UserID: req.UserID, ShowID: req.ShowID, Seats: req.Seats, Id: givenID})
	rsp.Id = givenID

	return nil
}

func (bs *BService) DeleteBooking(ctx context.Context, req *booking.DeleteBookingRequest, rsp *booking.DeleteBookingResult) error {
	// Delete from booking or notConfirmed list
	return nil
}

func (bs *BService) ConfirmBooking(ctx context.Context, req *booking.ConfirmBookingRequest, rsp *booking.ConfirmBookingResult) error {
	// move booking from notConfirmed to booking list
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.booking"),
	)

	service.Init()
	booking.RegisterBookingHandler(service.Server(), &BService{booking: make([]*booking.BookingData, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

