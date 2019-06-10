package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

func main() {


	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)
	bookingC := booking.NewBookingService("go.micro.services.booking", client)
	cinemaC := cinema.NewCinemaService("go.micro.services.cinema", client)
	movieC := movie.NewMovieService("go.micro.services.movie", client)
	userC := user.NewUserService("go.micro.services.user", client)


	rspS , err := showC.GetShowList(context.TODO(), &show.GetShowListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all shows: %v \n\n", rspS.Shows)

	rspB, err := bookingC.GetBookingList(context.TODO(), &booking.GetListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all Bookings: %v \n\n", rspB.Bookings)

	rspBNC, err := bookingC.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all not confirmed Bookings: %v \n\n", rspBNC.Bookings)

	rspC, err := cinemaC.GetHallList(context.TODO(), &cinema.GetHallListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all halls: %v \n\n", rspC.CHall)

	rspM, err := movieC.GetMovieList(context.TODO(), &movie.GetMovieListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all movies: %v \n\n", rspM.Movies)

	rspU, err := userC.GetUserList(context.TODO(), &user.GetUserListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all users: %v \n\n", rspU.Users)


	fmt.Println("----------------")
	fmt.Println("----------------")

	//rspCD, err := cinemaC.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id:1})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("Delete was: %v \n\n", rspCD.Successful)

	rspBTNC1, err := bookingC.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 2, ShowID: 2, Seats: 20})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("First Booking ID: %v \n\n", rspBTNC1.Id)

	rspBTNC2, err := bookingC.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 2, Seats: 20})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Second Booking ID: %v \n\n", rspBTNC2.Id)

	fmt.Println("----------------")
	fmt.Println("----------------")

	rspB, err = bookingC.GetBookingList(context.TODO(), &booking.GetListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all Bookings: %v \n\n", rspB.Bookings)

	rspBNC, err = bookingC.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all not confirmed Bookings: %v \n\n", rspBNC.Bookings)

	rspU, err = userC.GetUserList(context.TODO(), &user.GetUserListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all users: %v \n\n", rspU.Users)

	fmt.Println("----------------")
	fmt.Println("----------------")

	rspBTC1, err := bookingC.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspBTNC1.Id})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Confirmed first was: %v \n\n", rspBTC1.Successful)

	rspBTC2, err := bookingC.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspBTNC2.Id})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Confirmed second was: %v \n\n", rspBTC2.Successful)

	fmt.Println("----------------")
	fmt.Println("----------------")

	rspS , err = showC.GetShowList(context.TODO(), &show.GetShowListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all shows: %v \n\n", rspS.Shows)

	rspB, err = bookingC.GetBookingList(context.TODO(), &booking.GetListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all Bookings: %v \n\n", rspB.Bookings)

	rspU, err = userC.GetUserList(context.TODO(), &user.GetUserListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all users: %v \n\n", rspU.Users)



}
