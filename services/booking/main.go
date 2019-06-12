package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

type BService struct {
	booking      []*booking.BookingData
	notConfirmed []*booking.BookingData
	nextID       int32
	mux          sync.Mutex
}

func (bs *BService) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest, rsp *booking.CreateBookingResult) error {
	givenID := bs.nextID
	bs.nextID++

	if bs.userExist(req.UserID) && bs.showExist(req.ShowID) {
		if bs.checkSeats(req.ShowID) >= req.Seats {
			bs.notConfirmed = append(bs.notConfirmed,
				&booking.BookingData{UserID: req.UserID, ShowID: req.ShowID, Seats: req.Seats, Id: givenID})
			rsp.Id = givenID

			bs.sendUserBooking(req.UserID, givenID, false)

			return nil
		}
		rsp.Id = -1
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
			bs.mux.Lock()
			if bs.checkSeats(b.ShowID) >= b.Seats {
				bs.booking = append(bs.booking, b)
				// aus notConfirmed loeschen
				bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
				bs.updateSeats(b.ShowID, b.Seats)
				rsp.Successful = true
				bs.sendUserBooking(b.UserID, b.Id, true)
				bs.mux.Unlock()
				return nil
			} else {
				bs.informUser(b.UserID, b.Id)
				rsp.Successful = false
				bs.mux.Unlock()
				return nil
			}
		}
	}
	rsp.Successful = false
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

func (bs *BService) GetNotConfirmedList(ctx context.Context, req *booking.GetListRequest, rsp *booking.GetListResult) error {
	rsp.Bookings = bs.notConfirmed
	return nil
}

func (bs *BService) GetBookingList(ctx context.Context, req *booking.GetListRequest, rsp *booking.GetListResult) error {
	rsp.Bookings = bs.booking
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

func (bs *BService) updateSeats(showID int32, amount int32) {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	_, err := showC.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: showID, AmountSeats: amount})
	if err != nil {
		fmt.Println(err)
	}
}

func (bs *BService) sendUserBooking(userID int32, bookingID int32, confirmed bool) {
	var client client.Client
	userC := user.NewUserService("go.micro.services.user", client)

	err := errors.New("nil")
	err = nil

	if confirmed {
		_, err = userC.CreatedBooking(context.TODO(),
			&user.CreatedBookingRequest{UserID: userID, BookingID: bookingID})
	} else {
		_, err = userC.CreatedMarkedBooking(context.TODO(),
			&user.CreatedBookingRequest{UserID: userID, BookingID: bookingID})
	}

	if err != nil {
		fmt.Println(err)
	}
}

func (bs *BService) showExist(showID int32) bool {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	rsp, err := showC.Exist(context.TODO(), &show.ExistRequest{Id: showID})

	if err != nil {
		fmt.Println(err)
	}

	return rsp.Exist
}

func (bs *BService) userExist(userID int32) bool {
	var client client.Client
	userC := user.NewUserService("go.micro.services.user", client)

	rsp, err := userC.Exist(context.TODO(), &user.ExistRequest{Id: userID})

	if err != nil {
		fmt.Println(err)
	}

	return rsp.Exist
}

func (bs *BService) Exist(ctx context.Context, req *booking.ExistRequest, rsp *booking.ExistResult) error {
	for _, b := range bs.booking {
		if b.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	for _, nc := range bs.notConfirmed {
		if nc.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.booking"),
		micro.Address(fmt.Sprintf(":%v", 1032)),
	)

	service.Init()
	booking.RegisterBookingHandler(service.Server(), &BService{booking: exampleData(), notConfirmed: make([]*booking.BookingData, 0), nextID: 5})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func exampleData() []*booking.BookingData {
	bookings := make([]*booking.BookingData, 0)
	bookings = append(bookings, &booking.BookingData{Id: 1, UserID: 3, ShowID: 4, Seats: 2})
	bookings = append(bookings, &booking.BookingData{Id: 2, UserID: 4, ShowID: 3, Seats: 2})
	bookings = append(bookings, &booking.BookingData{Id: 3, UserID: 1, ShowID: 1, Seats: 4})
	bookings = append(bookings, &booking.BookingData{Id: 4, UserID: 2, ShowID: 3, Seats: 2})
	return bookings
}
