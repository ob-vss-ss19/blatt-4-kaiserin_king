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
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	shs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/srv"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
	us "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/srv"
)

func TestMovie(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunShowService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)

	service := MService{Movie: make([]*movie.MovieData, 0), NextID: 1}

	rsp := &movie.GetMovieListResult{}
	err := service.GetMovieList(context.TODO(), &movie.GetMovieListRequest{}, rsp)

	if err == nil {
		if len(rsp.Movies) != 0 {
			t.Errorf("Expected empty List")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspCreate := &movie.CreateMovieResult{}
	_ = service.CreateMovie(context.TODO(), &movie.CreateMovieRequest{Titel: "Sex and The City"}, rspCreate)

	rspExist := &movie.ExistResult{}
	_ = service.Exist(context.TODO(), &movie.ExistRequest{Id: rspCreate.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected movie to exist")
	}

	rspExist = &movie.ExistResult{}
	_ = service.Exist(context.TODO(), &movie.ExistRequest{Id: 200}, rspExist)

	if rspExist.Exist {
		t.Error("Expected movie not to exist")
	}

	rsp = &movie.GetMovieListResult{}
	err = service.GetMovieList(context.TODO(), &movie.GetMovieListRequest{}, rsp)

	if err == nil {
		if len(rsp.Movies) != 1 {
			t.Errorf("Expected List with one element")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete := &movie.DeleteMovieResult{}
	err = service.DeleteMovie(context.TODO(), &movie.DeleteMovieRequest{Id: rspCreate.Id}, rspDelete)

	if err == nil {
		if !rspDelete.Successful {
			t.Errorf("Expected successful deleting")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rsp = &movie.GetMovieListResult{}
	err = service.GetMovieList(context.TODO(), &movie.GetMovieListRequest{}, rsp)

	if err == nil {
		if len(rsp.Movies) != 0 {
			t.Errorf("Expected empty List")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete = &movie.DeleteMovieResult{}
	err = service.DeleteMovie(context.TODO(), &movie.DeleteMovieRequest{Id: 42}, rspDelete)

	if err == nil {
		if rspDelete.Successful {
			t.Errorf("Expected faling deleting")
		}
	} else {
		t.Errorf("Error %v\n", err)
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
