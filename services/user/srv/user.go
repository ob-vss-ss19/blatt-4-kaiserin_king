package srv

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

type UService struct {
	User   []*user.UserData
	NextID int32
	mux    sync.Mutex
}

//Function that creats a new user, which is requested with a name.
//Return the ID of the new user.
func (us *UService) CreateUser(ctx context.Context, req *user.CreateUserRequest, rsp *user.CreateUserResult) error {
	us.mux.Lock()
	givenID := us.NextID
	us.mux.Unlock()
	us.NextID++
	us.User = append(us.User, &user.UserData{Name: req.Name, Id: givenID})
	rsp.Id = givenID

	return nil
}

//Function that deletes user by ID.
//Returns success of the operation by bool value.
func (us *UService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, rsp *user.DeleteUserResult) error {
	// delete only if no bookings
	if us.CheckBookingOfUser(req.Id) {
		// kann geloescht werden, da keine Reservierungen vorhanden für aktuellen user
		for i, v := range us.User {
			if v.Id == req.Id {
				us.mux.Lock()
				us.User = append(us.User[:i], us.User[i+1:]...)
				us.mux.Unlock()
				rsp.Successful = true
				return nil
			}
		}
	}
	rsp.Successful = false
	return nil
}

//Function that gets called if an booking the Deleted.
//The booking will be removed from the user.
func (us *UService) BookingDeleted(ctx context.Context, req *user.BookingDeletedRequest,
	rsp *user.BookingDeletedResult) error {
	us.mux.Lock()
	if !us.deleteBooking(req.UserID, req.BookingID) {
		us.deleteNotConfirmed(req.UserID, req.BookingID)
	}
	us.mux.Unlock()
	return nil
}

//Function that adds the id of a new not-confirmed-Booking.
func (us *UService) CreatedMarkedBooking(ctx context.Context, req *user.CreatedBookingRequest,
	rsp *user.CreatedBookingResult) error {
	for _, u := range us.User {
		if req.UserID == u.Id {
			u.NotConfirmed = append(u.NotConfirmed, req.BookingID)
		}
	}
	return nil
}

//Function that adds the id of a new confirmed Booking.
func (us *UService) CreatedBooking(ctx context.Context, req *user.CreatedBookingRequest,
	rsp *user.CreatedBookingResult) error {
	for _, u := range us.User {
		if u.Id == req.UserID {
			u.Bookings = append(u.Bookings, req.BookingID)
			us.deleteNotConfirmed(req.UserID, req.BookingID)
			return nil
		}
	}
	return nil
}

//Function that returns list of all users.
func (us *UService) GetUserList(ctx context.Context, req *user.GetUserListRequest, rsp *user.GetUserListResult) error {
	rsp.Users = us.User
	return nil
}

//Function that returns if a user, given by his ID does exist.
func (us *UService) Exist(ctx context.Context, req *user.ExistRequest, rsp *user.ExistResult) error {
	for _, u := range us.User {
		if u.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

//Function that deletes a not-confirmed-booking from List of user.
func (us *UService) deleteNotConfirmed(userID int32, bookingID int32) bool {
	for _, u := range us.User {
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

//Function that deletes a confirmed booking from List of a user.
func (us *UService) deleteBooking(userID int32, bookingID int32) bool {
	for _, u := range us.User {
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

//Function that checks if a user got bookings.
func (us *UService) CheckBookingOfUser(userID int32) bool {
	// look if there are bookings of userID
	for _, u := range us.User {
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

func RunService() {
	service := micro.NewService(
		micro.Name("go.micro.services.user"),
		micro.Address(fmt.Sprintf(":%v", 1035)),
	)

	service.Init()

	err := user.RegisterUserHandler(service.Server(), &UService{User: ExampleData(), NextID: 5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

//Example Data of users which are added to the Service from Beginning
func ExampleData() []*user.UserData {
	users := make([]*user.UserData, 0)
	users = append(users, &user.UserData{Id: 1, Name: "Maxi King",
		Bookings: []int32{3}, NotConfirmed: make([]int32, 0)})

	users = append(users, &user.UserData{Id: 2, Name: "Kaiserin Sissy",
		Bookings: []int32{4}, NotConfirmed: make([]int32, 0)})

	users = append(users, &user.UserData{Id: 3, Name: "Simon der Weise",
		Bookings: []int32{1}, NotConfirmed: make([]int32, 0)})

	users = append(users, &user.UserData{Id: 4, Name: "Lisa Master",
		Bookings: []int32{2}, NotConfirmed: make([]int32, 0)})

	return users
}
