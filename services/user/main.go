package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

type UService struct {
	user         []*user.UserData
	notConfirmed []*user.CreatedBookingRequest
	bookings     []*user.CreatedBookingRequest
	nextID       int32
}

func (us *UService) CreateUser(ctx context.Context, req *user.CreateUserRequest, rsp *user.CreateUserResult) error {
	givenID := us.nextID
	us.nextID++
	us.user = append(us.user, &user.UserData{Name: req.Name, Id: givenID})
	rsp.Id = givenID

	return nil
}

func (us *UService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, rsp *user.DeleteUserResult) error {
	// delete only if no bookings

/*	var client client.Client
	userC := booking.NewBookingService("go.micro.services.booking", client)

	bookingRsp, err := userC.AskBookingOfUser(context.TODO(), &booking.AskBookingOfUserRequest{UserId: req.Id})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bookingRsp.NoBookings)*/

	if us.CheckBookingOfUser(req.Id) {
		// kann geloescht werden, da keine Reservierungen vorhanden für aktuellen user
		for i, v := range us.user {
			if v.Id == req.Id {
				us.user = append(us.user[:i], us.user[i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (us *UService) BookingDeleted(ctx context.Context, req *user.BookingDeletedRequest, rsp *user.BookingDeletedResult) error {
	if !us.deleteBooking(req.UserID, req.BookingID) {
		us.deleteNotConfirmed(req.UserID, req.BookingID)
	}
	return nil
}

func (us *UService) CreatedMarkedBooking(ctx context.Context, req *user.CreatedBookingRequest, rsp *user.CreatedBookingResult) error {
	us.notConfirmed = append(us.notConfirmed, req)
	return nil
}

func (us *UService) CreatedBooking(ctx context.Context, req *user.CreatedBookingRequest, rsp *user.CreatedBookingResult) error {
	us.bookings = append(us.bookings, req)
	us.deleteNotConfirmed(req.UserID, req.BookingID)
	return nil
}

func (us *UService) GetUserList(ctx context.Context, req *user.GetUserListRequest, rsp *user.GetUserListResult) error {
	rsp.Users = us.user
	return nil
}

func (us *UService) deleteNotConfirmed(userID int32, bookingID int32) bool {
	for i, b := range us.notConfirmed {
		if b.UserID == userID && b.BookingID == bookingID {
			us.notConfirmed = append(us.notConfirmed[:i], us.notConfirmed[i+1:]...)
			return true
		}
	}
	return false
}

func (us *UService) deleteBooking(userID int32, bookingID int32) bool {
	for i, b := range us.bookings {
		if b.UserID == userID && b.BookingID == bookingID {
			us.bookings = append(us.bookings[:i], us.bookings[i+1:]...)
			return true
		}
	}
	return false

}

func (us *UService) CheckBookingOfUser(userID int32) bool {
	// look if there are bookings of userID
	for _, b := range us.bookings {
		if b.UserID == userID {
			return false
		}
	}
	for _, b := range us.notConfirmed {
		if b.UserID == userID {
			return false
		}
	}
	return true
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.user"),
	)

	service.Init()
	user.RegisterUserHandler(service.Server(), &UService{user: make([]*user.UserData, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}
