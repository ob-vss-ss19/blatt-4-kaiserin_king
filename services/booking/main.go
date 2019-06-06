package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
	"log"

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

	if bs.checkSeats(req.ShowID) >= req.Seats {
		bs.notConfirmed = append(bs.notConfirmed,
			&booking.BookingData{UserID: req.UserID, ShowID: req.ShowID, Seats: req.Seats, Id: givenID})
		rsp.Id = givenID
		return nil
	}

	rsp.Id = -1
	return nil
}

func (bs *BService) DeleteBooking(ctx context.Context, req *booking.DeleteBookingRequest, rsp *booking.DeleteBookingResult) error {
	// TODO: Inform user that booking was deleted
	// Delete from booking or notConfirmed list
	for i, b := range bs.booking {
		if b.Id == req.Id {
			//bs.booking = append(bs.booking[:i], bs.booking[i+1:]...)
			bs.deleteFromBooking(i, b.UserID, b.Id)
			rsp.Successful = true
			return nil
		}
	}
	for i, b := range bs.notConfirmed {
		if b.Id == req.Id {
			//bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
			bs.deleteFromNotConfirmed(i, b.UserID, b.Id)
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
			if bs.checkSeats(b.ShowID) >= b.Seats {
				bs.booking = append(bs.booking, b)
				// aus notConfirmed loeschen
				bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
				bs.updateSeats(b.ShowID)
				rsp.Successful = true
				return nil
			} else {
				rsp.Successful = false
				return nil
			}
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
	success := false

	// delete show with id -> delete bookings
	for i, b := range bs.booking {
		if b.ShowID == req.Id {
			//bs.booking = append(bs.booking[:i], bs.booking[i+1:]...)
			bs.deleteFromBooking(i, b.UserID, b.Id)
			success = true
		}
	}
	for i, b := range bs.notConfirmed {
		if b.ShowID == req.Id {
			//bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
			bs.deleteFromNotConfirmed(i, b.UserID, b.Id)
			success = true
		}
	}
	rsp.Successful = success
	return nil
}

func (bs *BService) deleteFromNotConfirmed(index int, userID int32, bookingID int32) {
	bs.notConfirmed = append(bs.notConfirmed[:index], bs.notConfirmed[index+1:]...)
	bs.informUser(userID, bookingID)
}

func (bs *BService) deleteFromBooking(index int, userID int32, bookingID int32) {
	bs.booking = append(bs.booking[:index], bs.booking[index+1:]...)
	bs.informUser(userID, bookingID)
}

func (bs *BService) informUser(userID int32, bookingID int32) {
	var client client.Client
	userC := user.NewUserService("go.micro.services.user", client)

	_, err := userC.BookingDeleted(context.TODO(), &user.BookingDeletedRequest{UserID: userID, BookingID: bookingID})
	if err != nil {
		fmt.Println(err)
	}
}

func (bs *BService) checkSeats(showID int32) int32 {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	rspShow, err := showC.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: showID})
	if err != nil {
		fmt.Println(err)
	}

	return rspShow.FreeSeats
}

func (bs *BService) updateSeats(showID int32) {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	rspShow, err := showC.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: showID})
	if err != nil {
		fmt.Println(err)
	}
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

