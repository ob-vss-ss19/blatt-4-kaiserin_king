package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.client"))
	service.Init()

/*	// Create new greeter client
	showC := show.NewShowService("go.micro.services.show", service.Client())

	// Call the greeter
	rsp, err := showC.GetShowList(context.TODO(), &show.GetShowListRequest{})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Printf("Show list: %v", rsp.Shows)
	fmt.Println("test")*/

	//var client client.Client
	showC := show.NewShowService("go.micro.services.show", service.Client())
	bookingC := booking.NewBookingService("go.micro.services.booking", service.Client())
	cinemaC := cinema.NewCinemaService("go.micro.services.cinema", service.Client())
	movieC := movie.NewMovieService("go.micro.services.movie", service.Client())
	userC := user.NewUserService("go.micro.services.user", service.Client())

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

	fmt.Println("##########")
	fmt.Println("##########")
	fmt.Println("")

	rspCD, err := cinemaC.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id:2})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all halls: %v \n\n", rspCD.Successful)

	fmt.Println("##########")
	fmt.Println("##########")
	fmt.Println("")

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

	rspBNC, err = bookingC.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all not confirmed Bookings: %v \n\n", rspBNC.Bookings)

	rspC, err = cinemaC.GetHallList(context.TODO(), &cinema.GetHallListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all halls: %v \n\n", rspC.CHall)

	rspM, err = movieC.GetMovieList(context.TODO(), &movie.GetMovieListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all movies: %v \n\n", rspM.Movies)

	rspU, err = userC.GetUserList(context.TODO(), &user.GetUserListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all users: %v \n\n", rspU.Users)



}
