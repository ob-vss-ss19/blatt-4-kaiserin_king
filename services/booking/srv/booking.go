package srv

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

//Struct for a bookingservice
type BService struct {
	Booking      []*booking.BookingData
	NotConfirmed []*booking.BookingData
	NextID       int32
	Mux          sync.Mutex
}

//Function to create a new "Booking".
//Gets an request with an UserID and ShowID, if both ot the exist the Result will return a valid Id,
//else it will return -1. The "booking" will be added as not confirmed
func (bs *BService) CreateBooking(ctx context.Context,
	req *booking.CreateBookingRequest,
	rsp *booking.CreateBookingResult) error {

	if bs.userExist(req.UserID) && bs.showExist(req.ShowID) {
		if bs.checkSeats(req.ShowID) >= req.Seats {
			bs.Mux.Lock()
			givenID := bs.NextID
			bs.NextID++
			bs.Mux.Unlock()
			bs.NotConfirmed = append(bs.NotConfirmed,
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

//Function to delete a "booking" by Request with an ID.
//Return if the operation was successful by a bool value.
func (bs *BService) DeleteBooking(ctx context.Context, req *booking.DeleteBookingRequest,
	rsp *booking.DeleteBookingResult) error {
	// Delete from booking or notConfirmed list
	for i, b := range bs.Booking {
		if b.Id == req.Id {
			//bs.booking = append(bs.booking[:i], bs.booking[i+1:]...)
			bs.Mux.Lock()
			bs.deleteFromBooking(i, b.UserID, b.Id)
			bs.Mux.Unlock()
			rsp.Successful = true
			return nil
		}
	}
	for i, b := range bs.NotConfirmed {
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

//Function to Confirm a Booking by ID.
//Return if the operation was successful by a bool value.
func (bs *BService) ConfirmBooking(ctx context.Context, req *booking.ConfirmBookingRequest,
	rsp *booking.ConfirmBookingResult) error {
	// move booking from notConfirmed to booking list
	for i, b := range bs.NotConfirmed {
		if b.Id == req.Id {
			if bs.checkSeats(b.ShowID) >= b.Seats {
				bs.Mux.Lock()
				bs.Booking = append(bs.Booking, b)
				// aus notConfirmed loeschen
				bs.NotConfirmed = append(bs.NotConfirmed[:i], bs.NotConfirmed[i+1:]...)
				bs.updateSeats(b.ShowID, b.Seats)
				bs.Mux.Unlock()
				rsp.Successful = true
				bs.sendUserBooking(b.UserID, b.Id, true)
				return nil
			}
			bs.informUser(b.UserID, b.Id)
			rsp.Successful = false
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

//This Functions gets called when an Show got deleted, maybe there are bookings with this show which should be removed.
//Gets the ID of the delete show and returns if the operation was successful by a bool value.
func (bs *BService) FromShowDelete(ctx context.Context, req *booking.FromShowDeleteRequest,
	rsp *booking.FromShowDeleteResult) error {
	success := false

	// delete show with id -> delete bookings
	for i, b := range bs.Booking {
		if b.ShowID == req.Id {
			//bs.booking = append(bs.booking[:i], bs.booking[i+1:]...)
			bs.Mux.Lock()
			bs.deleteFromBooking(i, b.UserID, b.Id)
			bs.Mux.Unlock()
			success = true
		}
	}
	for i, b := range bs.NotConfirmed {
		if b.ShowID == req.Id {
			//bs.notConfirmed = append(bs.notConfirmed[:i], bs.notConfirmed[i+1:]...)
			bs.Mux.Lock()
			bs.deleteFromNotConfirmed(i, b.UserID, b.Id)
			bs.Mux.Unlock()
			success = true
		}
	}
	rsp.Successful = success
	return nil
}

//Function which return List of all not confirmed bookings.
func (bs *BService) GetNotConfirmedList(ctx context.Context, req *booking.GetListRequest,
	rsp *booking.GetListResult) error {
	rsp.Bookings = bs.NotConfirmed
	return nil
}

//Function which return List of all confirmed bookings.
func (bs *BService) GetBookingList(ctx context.Context, req *booking.GetListRequest, rsp *booking.GetListResult) error {
	rsp.Bookings = bs.Booking
	return nil
}

//Function to delete an booking from the not confirmed List.
//Also informs the User of the Booking that his booking go deleted
func (bs *BService) deleteFromNotConfirmed(index int, userID int32, bookingID int32) {
	bs.NotConfirmed = append(bs.NotConfirmed[:index], bs.NotConfirmed[index+1:]...)
	bs.informUser(userID, bookingID)
}

//Function to delete an booking from the confirmed Booking List.
//Also informs the User of the Booking that his booking go deleted
func (bs *BService) deleteFromBooking(index int, userID int32, bookingID int32) {
	bs.Booking = append(bs.Booking[:index], bs.Booking[index+1:]...)
	bs.informUser(userID, bookingID)
}

//This Functions implements the logic of informing the user about the delete of his booking.
func (bs *BService) informUser(userID int32, bookingID int32) {
	var client client.Client
	userC := user.NewUserService("go.micro.services.user", client)

	_, err := userC.BookingDeleted(context.TODO(), &user.BookingDeletedRequest{UserID: userID, BookingID: bookingID})
	if err != nil {
		fmt.Println(err)
	}
}

//Function that return the amounts of seat which are currently free in a show.
func (bs *BService) checkSeats(showID int32) int32 {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	rspShow, err := showC.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: showID})
	if err != nil {
		fmt.Println(err)
	}
	return rspShow.FreeSeats
}

//Function to update the amount of seat that are free in a show.
func (bs *BService) updateSeats(showID int32, amount int32) {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	_, err := showC.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: showID, AmountSeats: amount})
	if err != nil {
		fmt.Println(err)
	}
}

//Function that send a user the ID of his new created Booking.
//The Parameter confirmed choose if the user gets informed for a new not-confirmed-booking or a confirmed one.
func (bs *BService) sendUserBooking(userID int32, bookingID int32, confirmed bool) {
	var client client.Client
	userC := user.NewUserService("go.micro.services.user", client)

	if confirmed {
		_, err := userC.CreatedBooking(context.TODO(),
			&user.CreatedBookingRequest{UserID: userID, BookingID: bookingID})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err := userC.CreatedMarkedBooking(context.TODO(),
			&user.CreatedBookingRequest{UserID: userID, BookingID: bookingID})
		if err != nil {
			fmt.Println(err)
		}
	}

}

//Function that return if a show given by its ID does exist.
func (bs *BService) showExist(showID int32) bool {
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	rsp, err := showC.Exist(context.TODO(), &show.ExistRequest{Id: showID})

	if err != nil {
		fmt.Println(err)
	}

	return rsp.Exist
}

//Function that return if a user given by its ID does exist.
func (bs *BService) userExist(userID int32) bool {
	var client client.Client
	userC := user.NewUserService("go.micro.services.user", client)

	rsp, err := userC.Exist(context.TODO(), &user.ExistRequest{Id: userID})

	if err != nil {
		fmt.Println(err)
	}

	return rsp.Exist
}

//Function that return if a booking given by its ID does exist.
func (bs *BService) Exist(ctx context.Context, req *booking.ExistRequest, rsp *booking.ExistResult) error {
	for _, b := range bs.Booking {
		if b.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	for _, nc := range bs.NotConfirmed {
		if nc.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

func RunService() {
	service := micro.NewService(
		micro.Name("go.micro.services.booking"),
		micro.Address(fmt.Sprintf(":%v", 1032)),
	)

	service.Init()
	err := booking.RegisterBookingHandler(service.Server(),
		&BService{Booking: ExampleData(),
			NotConfirmed: make([]*booking.BookingData, 0),
			NextID:       5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

//Function which return example Data of bookings which are added to the Service from the Beginning.
func ExampleData() []*booking.BookingData {
	bookings := make([]*booking.BookingData, 0)
	bookings = append(bookings, &booking.BookingData{Id: 1, UserID: 3, ShowID: 4, Seats: 2})
	bookings = append(bookings, &booking.BookingData{Id: 2, UserID: 4, ShowID: 3, Seats: 2})
	bookings = append(bookings, &booking.BookingData{Id: 3, UserID: 1, ShowID: 1, Seats: 4})
	bookings = append(bookings, &booking.BookingData{Id: 4, UserID: 2, ShowID: 3, Seats: 2})
	return bookings
}
