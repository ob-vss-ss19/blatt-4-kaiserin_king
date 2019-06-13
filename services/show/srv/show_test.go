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
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
	us "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/srv"
)

func TestShow(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go RunCinemaService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunMovieService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunBookingService(ctx, true)
	time.Sleep(time.Second * 3)
	go RunUserService(ctx, true)
	time.Sleep(time.Second * 3)
	service := SService{Show: make([]*show.ShowData, 0), NextID: 1}

	rsp := &show.GetShowListResult{}
	err := service.GetShowList(context.TODO(), &show.GetShowListRequest{}, rsp)
	if err == nil {
		if len(rsp.Shows) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspCreate := &show.CreateShowResult{}
	err = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 1, HallID: 1}, rspCreate)
	if err == nil {
		if rspCreate.Id == -1 {
			t.Error("Expected successful creation")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspSeats := &show.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: rspCreate.Id}, rspSeats)
	if rspSeats.FreeSeats != 35 {
		t.Errorf("Expected 35 free seats but got %v", rspSeats.FreeSeats)
	}

	rspUpdate := &show.UpdateSeatsResult{}
	_ = service.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: rspCreate.Id, AmountSeats: 5}, rspUpdate)
	if !rspUpdate.Success {
		t.Errorf("Expected successful update")
	}

	rspUpdate = &show.UpdateSeatsResult{}
	_ = service.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: 200, AmountSeats: 5}, rspUpdate)
	if rspUpdate.Success {
		t.Errorf("Expected failing update")
	}

	rspSeats = &show.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: rspCreate.Id}, rspSeats)
	if rspSeats.FreeSeats != 30 {
		t.Errorf("Expected 30 free seats but got %v", rspSeats.FreeSeats)
	}

	rspSeats = &show.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: 200}, rspSeats)
	if rspSeats.FreeSeats != -1 {
		t.Errorf("Expected 30 free seats but got %v", rspSeats.FreeSeats)
	}

	rspCreate2 := &show.CreateShowResult{}
	err = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 200, HallID: 200}, rspCreate2)
	if err == nil {
		if rspCreate2.Id != -1 {
			t.Error("Expected failing creation")
		}
	} else {
		t.Error("Error with Request!")
	}

	rsp = &show.GetShowListResult{}
	err = service.GetShowList(context.TODO(), &show.GetShowListRequest{}, rsp)
	if err == nil {
		if len(rsp.Shows) != 1 {
			t.Error("Expected List with one Element")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete := &show.DeleteShowResult{}
	err = service.DeleteShow(context.TODO(), &show.DeleteShowRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if !rspDelete.Successful {
			t.Error("Expected successful deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

	rsp = &show.GetShowListResult{}
	err = service.GetShowList(context.TODO(), &show.GetShowListRequest{}, rsp)
	if err == nil {
		if len(rsp.Shows) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete = &show.DeleteShowResult{}
	err = service.DeleteShow(context.TODO(), &show.DeleteShowRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing delete")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspCreate = &show.CreateShowResult{}
	_ = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 1, HallID: 1}, rspCreate)

	rspDeleteHall := &show.DeleteShowOfHallResult{}
	_ = service.FromHallDelete(context.TODO(), &show.DeleteShowOfHallRequest{HallID: 1}, rspDeleteHall)
	if !rspDeleteHall.Successful {
		t.Error("Expected successful deleting from Hall")
	}

	rspDeleteHall = &show.DeleteShowOfHallResult{}
	_ = service.FromHallDelete(context.TODO(), &show.DeleteShowOfHallRequest{HallID: 200}, rspDeleteHall)
	if rspDeleteHall.Successful {
		t.Error("Expected failing deleting from Hall")
	}

	rspCreate = &show.CreateShowResult{}
	_ = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 1, HallID: 1}, rspCreate)

	rspExist := &show.ExistResult{}
	_ = service.Exist(context.TODO(), &show.ExistRequest{Id: rspCreate.Id}, rspExist)
	if !rspExist.Exist {
		t.Error("Expected that show exists")
	}

	rspDeleteMovie := &show.DeleteShowOfMovieResult{}
	_ = service.FromMovieDelete(context.TODO(), &show.DeleteShowOfMovieRequest{MovieID: 1}, rspDeleteMovie)
	if !rspDeleteMovie.Successful {
		t.Error("Expected successful deleting from Movie")
	}

	rspDeleteMovie = &show.DeleteShowOfMovieResult{}
	_ = service.FromMovieDelete(context.TODO(), &show.DeleteShowOfMovieRequest{MovieID: 200}, rspDeleteMovie)
	if rspDeleteMovie.Successful {
		t.Error("Expected failing deleting from Movie")
	}

	rspExist = &show.ExistResult{}
	_ = service.Exist(context.TODO(), &show.ExistRequest{Id: rspCreate.Id}, rspExist)
	if rspExist.Exist {
		t.Error("Expected that show does not exists")
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
