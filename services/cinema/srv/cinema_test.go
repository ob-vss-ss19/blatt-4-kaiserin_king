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
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	ms "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/srv"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	shs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/srv"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
	us "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/srv"
)

func TestCinemaGetList(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := CService{CHall: make([]*cinema.CinemaHall, 0), NextID: 1}
	rsp := &cinema.GetHallListResult{}
	err := service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 0 {
			t.Errorf("Expected empty hall list!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspCreate := &cinema.CreateHallResult{}
	_ = service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 1, Cols: 1}, rspCreate)
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 1 {
			t.Errorf("Expected hall list with one entry!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete := &cinema.DeleteHallResult{}
	_ = service.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id: rspCreate.Id}, rspDelete)
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if !rspDelete.Successful {
			t.Errorf("Expected successful delete!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rsp = &cinema.GetHallListResult{}
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 0 {
			t.Errorf("Expected empty hall list!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	cancel()
}

func TestCinemaDeleteWrongID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := CService{CHall: make([]*cinema.CinemaHall, 0), NextID: 1}

	rspDelete := &cinema.DeleteHallResult{}
	err := service.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id: 42}, rspDelete)

	if err == nil {
		if rspDelete.Successful {
			t.Errorf("Expected failing delete!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	cancel()
}

func TestAskForSeats(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)
	service := CService{CHall: make([]*cinema.CinemaHall, 0), NextID: 1}

	rspCreate := &cinema.CreateHallResult{}
	_ = service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 10, Cols: 3}, rspCreate)

	rspSeats := &cinema.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &cinema.FreeSeatsRequest{HallID: rspCreate.Id}, rspSeats)

	if rspSeats.FreeSeats != 30 {
		t.Errorf("Expected 30 seats but got %v", rspSeats.FreeSeats)
	}

	rspSeats = &cinema.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &cinema.FreeSeatsRequest{HallID: 200}, rspSeats)

	if rspSeats.FreeSeats != -1 {
		t.Errorf("Expected 30 seats but got %v", rspSeats.FreeSeats)
	}
	cancel()
}

func TestAskExist(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := CService{CHall: make([]*cinema.CinemaHall, 0), NextID: 1}

	rspCreate := &cinema.CreateHallResult{}
	_ = service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 10, Cols: 3}, rspCreate)

	rspExist := &cinema.ExistResult{}
	_ = service.Exist(context.TODO(), &cinema.ExistRequest{Id: rspCreate.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected existing hall")
	}

	rspExist = &cinema.ExistResult{}
	_ = service.Exist(context.TODO(), &cinema.ExistRequest{Id: 200}, rspExist)

	if rspExist.Exist {
		t.Error("Expected not existing hall")
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
