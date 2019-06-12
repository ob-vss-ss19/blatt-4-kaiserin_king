package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

type UService struct {
	user   []*user.UserData
	nextID int32
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
	if us.CheckBookingOfUser(req.Id) {
		// kann geloescht werden, da keine Reservierungen vorhanden für aktuellen user
		for i, v := range us.user {
			if v.Id == req.Id {
				us.user = append(us.user[:i], us.user[i+1:]...)
				rsp.Successful = true
				return nil
			}
		}
	}
	rsp.Successful = false
	return nil
}

func (us *UService) BookingDeleted(ctx context.Context, req *user.BookingDeletedRequest, rsp *user.BookingDeletedResult) error {
	if !us.deleteBooking(req.UserID, req.BookingID) {
		us.deleteNotConfirmed(req.UserID, req.BookingID)
	}
	return nil
}

func (us *UService) CreatedMarkedBooking(ctx context.Context, req *user.CreatedBookingRequest, rsp *user.CreatedBookingResult) error {
	for _, u := range us.user {
		if req.UserID == u.Id {
			u.NotConfirmed = append(u.NotConfirmed, req.BookingID)
		}
	}
	return nil
}

func (us *UService) CreatedBooking(ctx context.Context, req *user.CreatedBookingRequest, rsp *user.CreatedBookingResult) error {
	for _, u := range us.user {
		if u.Id == req.UserID {
			u.Bookings = append(u.Bookings, req.BookingID)
			us.deleteNotConfirmed(req.UserID, req.BookingID)
			return nil
		}
	}
	return nil
}

func (us *UService) GetUserList(ctx context.Context, req *user.GetUserListRequest, rsp *user.GetUserListResult) error {
	rsp.Users = us.user
	return nil
}

func (us *UService) Exist(ctx context.Context, req *user.ExistRequest, rsp *user.ExistResult) error {
	for _, u := range us.user {
		if u.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

func (us *UService) deleteNotConfirmed(userID int32, bookingID int32) bool {
	for _, u := range us.user {
		if u.Id == userID {
			for i, b := range u.NotConfirmed {
				if b == bookingID {
					u.NotConfirmed = append(u.NotConfirmed[:i], u.NotConfirmed[i+1:]...)
					return true
				}
			}
		}
	}
	return false
}

func (us *UService) deleteBooking(userID int32, bookingID int32) bool {
	for _, u := range us.user {
		if u.Id == userID {
			for i, b := range u.Bookings {
				if b == bookingID {
					u.Bookings = append(u.Bookings[:i], u.Bookings[i+1:]...)
					return true
				}
			}
		}
	}
	return false
}

func (us *UService) CheckBookingOfUser(userID int32) bool {
	// look if there are bookings of userID
	for _, u := range us.user {
		if u.Id == userID {
			if len(u.Bookings) != 0 {
				return false
			}
		}
		if u.Id == userID {
			if len(u.NotConfirmed) != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.user"),
		micro.Address(fmt.Sprintf(":%v",1036)),
	)

	service.Init()
	_ = user.RegisterUserHandler(service.Server(), &UService{user: exampleData(), nextID: 5})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func exampleData() []*user.UserData {
	users := make([]*user.UserData, 0)
	users = append(users, &user.UserData{Id: 1, Name: "Maxi King",
		Bookings:[]int32{3}, NotConfirmed: make([]int32, 0)})

	users = append(users, &user.UserData{Id: 2, Name: "Kaiserin Sissy",
		Bookings: []int32{4}, NotConfirmed: make([]int32, 0)})

	users = append(users, &user.UserData{Id: 3, Name: "Simon der Weise",
		Bookings: []int32{1}, NotConfirmed: make([]int32, 0)})

	users = append(users, &user.UserData{Id: 4, Name: "Lisa Master",
		Bookings: []int32{2}, NotConfirmed: make([]int32, 0)})

	return users
}
