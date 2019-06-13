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
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	cs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/srv"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	ms "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/srv"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	shs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/srv"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
	us "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/srv"
)

func TestBookingGetList(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}
	rsp := &booking.GetListResult{}
	err := service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected empty list! \n")
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)
	if (rspCreate.Id) == -1 {
		t.Error("Adding booking failed! \n")
	}
	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)
	err = service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	cancel()
}

func TestBookingTooMuchSeats(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(),
		&booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 500},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Errorf("Expected no booking because of too much seats!\n")
	}

	cancel()
}

func TestBookingTwoConfirmed(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(),
		&booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 200},
		rspCreate)

	rspCreate2 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(),
		&booking.CreateBookingRequest{UserID: 2, ShowID: 1, Seats: 400},
		rspCreate2)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)
	if !rspConfirm.Successful {
		t.Errorf("Booking expected successful")
	}

	rspConfirm2 := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate2.Id}, rspConfirm2)
	if rspConfirm2.Successful {
		t.Errorf("Booking 2 expected not successful")
	}

	rspConfirm3 := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: 42}, rspConfirm3)
	if rspConfirm3.Successful {
		t.Errorf("Booking 3 expected not successful")
	}

	cancel()
}

func TestBookingDelete(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}
	rsp := &booking.GetListResult{}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)

	rspDelete := &booking.DeleteBookingResult{}
	_ = service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id: rspCreate.Id}, rspDelete)

	err := service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rsp = &booking.GetListResult{}

	rspCreate = &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)

	rspDelete = &booking.DeleteBookingResult{}
	_ = service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id: rspCreate.Id}, rspDelete)

	err = service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspDelete = &booking.DeleteBookingResult{}
	_ = service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id: 6000}, rspDelete)
	if rspDelete.Successful {
		t.Error("Expected no booking with this ID!")
	}

	cancel()
}

func TestCreateWrongID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 200, ShowID: 1, Seats: 1},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Error("Expected no booking")
	}

	rspCreate = &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 200, Seats: 1},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Error("Expected no booking")
	}

	rspCreate = &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 200, Seats: 200},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Error("Expected no booking")
	}

	cancel()
}

func TestFromShowDelete(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate1 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate1)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate1.Id}, rspConfirm)

	rspCreate2 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 2, ShowID: 2, Seats: 1}, rspCreate2)

	rsp := &booking.GetListResult{}
	err := service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one element! \n")
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspDelete := &booking.FromShowDeleteResult{}
	_ = service.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: 1}, rspDelete)

	err = service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected empty list! \n")
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspDelete = &booking.FromShowDeleteResult{}
	_ = service.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: 2}, rspDelete)

	if !rspDelete.Successful {
		t.Error("Expected successful deleting!")
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected empty list! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	_ = service.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: 200}, rspDelete)

	if rspDelete.Successful {
		t.Error("Expected failing deleting!")
	}

	cancel()
}

func TestExist(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate1 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate1)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate1.Id}, rspConfirm)

	rspCreate2 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 2, ShowID: 2, Seats: 1}, rspCreate2)

	rspExist := &booking.ExistResult{}
	_ = service.Exist(context.TODO(), &booking.ExistRequest{Id: rspCreate1.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected that booking excists")
	}

	rspExist = &booking.ExistResult{}
	_ = service.Exist(context.TODO(), &booking.ExistRequest{Id: rspCreate2.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected that booking excists")
	}

	rspExist = &booking.ExistResult{}
	_ = service.Exist(context.TODO(), &booking.ExistRequest{Id: 200}, rspExist)

	if rspExist.Exist {
		t.Error("Expected that booking not exists")
	}

	cancel()
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

func RunUserService(ctx context.Context, test bool) {
	port := 0
	if test {
		rand.Seed(time.Now().UTC().UnixNano())
		port = 1024 + rand.Intn(1000) + 8
	}

	service := micro.NewService(
		micro.Name("go.micro.services.user"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(ctx),
	)

	if !test {
		service.Init()
	}

	err := user.RegisterUserHandler(service.Server(), &us.UService{User: us.ExampleData(), NextID: 5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}
