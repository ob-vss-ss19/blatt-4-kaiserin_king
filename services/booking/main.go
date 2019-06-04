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
	for i, b := range bs.booking {
		if b.Id == req.Id {
			bs.booking = append(bs.booking[:i], bs.booking[i+1:]...)
			rsp.Successful = true
			return nil
		}
	}
	for i, b := range bs.notConfirmed {
		if b.Id == req.Id {
			bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

func (bs *BService) ConfirmBooking(ctx context.Context, req *booking.ConfirmBookingRequest, rsp *booking.ConfirmBookingResult) error {
	// move booking from notConfirmed to booking list
	for i, b := range bs.notConfirmed {
		if b.Id == req.Id {
			bs.booking = append(bs.booking, b)
			// aus notConfirmed loeschen
			bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

func (bs *BService) AskBookingOfUser(ctx context.Context, req *booking.AskBookingOfUserRequest, rsp *booking.AskBookingOfUserResult) error {
	// look if there are bookings of userID
	for _, b := range bs.booking {
		if b.UserID == req.UserId {
			rsp.NoBookings = false
			return nil
		}
	}
	for _, b := range bs.notConfirmed {
		if b.UserID == req.UserId {
			rsp.NoBookings = false
			return nil
		}
	}
	rsp.NoBookings = true
	return nil
}

func (bs *BService) FromShowDelete(ctx context.Context, req *booking.FromShowDeleteRequest, rsp *booking.FromShowDeleteResult) error {
	// delete show with id
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.booking"),
	)

	service.Init()
	booking.RegisterBookingHandler(service.Server(), &BService{booking: make([]*booking.BookingData, 0), notConfirmed: make([]*booking.BookingData, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

