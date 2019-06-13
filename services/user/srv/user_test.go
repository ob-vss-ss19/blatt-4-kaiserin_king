package srv

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/micro/go-micro"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	bs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/srv"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	cs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/srv"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	ms "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/srv"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	shs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/srv"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

func TestUser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)

	service := UService{User: make([]*user.UserData, 0), NextID: 1}

	rsp := &user.GetUserListResult{}
	err := service.GetUserList(context.TODO(), &user.GetUserListRequest{}, rsp)
	if err == nil {
		if len(rsp.Users) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspCreate := &user.CreateUserResult{}
	_ = service.CreateUser(context.TODO(), &user.CreateUserRequest{Name: "Max Mustermann"}, rspCreate)

	rspExist := &user.ExistResult{}
	_ = service.Exist(context.TODO(), &user.ExistRequest{Id: rspCreate.Id}, rspExist)
	if !rspExist.Exist {
		t.Error("Expected user to exist")
	}

	rspExist = &user.ExistResult{}
	_ = service.Exist(context.TODO(), &user.ExistRequest{Id: 200}, rspExist)
	if rspExist.Exist {
		t.Error("Expected user not to exist")
	}

	rsp = &user.GetUserListResult{}
	err = service.GetUserList(context.TODO(), &user.GetUserListRequest{}, rsp)
	if err == nil {
		if len(rsp.Users) != 1 {
			t.Error("Expected List with one Element")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete := &user.DeleteUserResult{}
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if !rspDelete.Successful {
			t.Error("Expected successful deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

	rsp = &user.GetUserListResult{}
	err = service.GetUserList(context.TODO(), &user.GetUserListRequest{}, rsp)
	if err == nil {
		if len(rsp.Users) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete = &user.DeleteUserResult{}
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete = &user.DeleteUserResult{}
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: 200}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing deleting")
		}
	} else {
		t.Error("Error with Request!")
	}
	cancel()
}

func TestShowDeleting(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)

	service := UService{User: ExampleData(), NextID: 5}

	rspDeleteBooking := &user.BookingDeletedResult{}
	_ = service.BookingDeleted(context.TODO(), &user.BookingDeletedRequest{UserID: 4, BookingID: 2}, rspDeleteBooking)
	if !service.CheckBookingOfUser(4) {
		t.Error("Expected no bookings for id=4")
	}

	rspCreate := &user.CreateUserResult{}
	_ = service.CreateUser(context.TODO(), &user.CreateUserRequest{Name: "Max Mustermann"}, rspCreate)
	if !service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected no bookings for id=%v", rspCreate.Id)
	}

	rspMarkedB := &user.CreatedBookingResult{}
	_ = service.CreatedMarkedBooking(context.TODO(),
		&user.CreatedBookingRequest{UserID: rspCreate.Id, BookingID: 200},
		rspMarkedB)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	rspMarkedB = &user.CreatedBookingResult{}
	_ = service.CreatedBooking(context.TODO(), &user.CreatedBookingRequest{UserID: 700, BookingID: 200}, rspMarkedB)
	if !service.CheckBookingOfUser(700) {
		t.Error("Expected no bookings for id=700")
	}

	rspMarkedB = &user.CreatedBookingResult{}
	_ = service.CreatedBooking(context.TODO(), &user.CreatedBookingRequest{UserID: rspCreate.Id, BookingID: 200},
		rspMarkedB)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	_ = service.CreatedMarkedBooking(context.TODO(),
		&user.CreatedBookingRequest{UserID: rspCreate.Id, BookingID: 700},
		rspMarkedB)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	rspDeleteBooking = &user.BookingDeletedResult{}
	_ = service.BookingDeleted(context.TODO(),
		&user.BookingDeletedRequest{UserID: rspCreate.Id, BookingID: 700},
		rspDeleteBooking)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	cancel()
}

func RunBookingService(ctx context.Context, test bool) {
	port := 0
	if test {
		rand.Seed(time.Now().UTC().UnixNano())
		port = 1024 + rand.Intn(1000) + 8
	}

	service := micro.NewService(
		micro.Name("go.micro.services.booking"),
		micro.Address(fmt.Sprintf(":%v", port)),
	)

	if !test {
		service.Init()
	}

	err := booking.RegisterBookingHandler(service.Server(),
		&bs.BService{Booking: bs.ExampleData(),
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

func RunCinemaService(ctx context.Context, test bool) {
	port := 0
	if test {
		rand.Seed(time.Now().UTC().UnixNano())
		port = 1024 + rand.Intn(1000) + 8
	}

	service := micro.NewService(
		micro.Name("go.micro.services.cinema"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(ctx),
	)

	if !test {
		service.Init()
	}

	err := cinema.RegisterCinemaHandler(service.Server(), &cs.CService{CHall: cs.ExampleData(), NextID: 3})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func RunMovieService(ctx context.Context, test bool) {
	port := 0
	if test {
		rand.Seed(time.Now().UTC().UnixNano())
		port = 1024 + rand.Intn(1000) + 8
	}

	service := micro.NewService(
		micro.Name("go.micro.services.movie"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(ctx),
	)

	if !test {
		service.Init()
	}

	err := movie.RegisterMovieHandler(service.Server(), &ms.MService{Movie: ms.ExampleData(), NextID: 5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func RunShowService(ctx context.Context, test bool) {
	port := 0
	if test {
		rand.Seed(time.Now().UTC().UnixNano())
		port = 1024 + rand.Intn(1000) + 8
	}

	service := micro.NewService(
		micro.Name("go.micro.services.show"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(ctx),
	)

	if !test {
		service.Init()
	}

	err := show.RegisterShowHandler(service.Server(), &shs.SService{Show: shs.ExampleData(), NextID: 5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}
