package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	"log"

	"github.com/micro/go-micro"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

type SService struct {
	show   []*show.ShowData
	nextID int32
}

func (shs *SService) CreateShow(ctx context.Context, req *show.CreateShowRequest, rsp *show.CreateShowResult) error {
	givenID := shs.nextID
	shs.nextID++

	var client client.Client
	cinemaC := cinema.NewCinemaService("go.micro.services.cinema", client)

	resSeats, err := cinemaC.AskSeats(context.TODO(), &cinema.FreeSeatsRequest{HallID: req.HallID})
	if err != nil {
		fmt.Println(err)
	}

	shs.show = append(shs.show,
		&show.ShowData{HallID: req.HallID, MovieID: req.MovieID, Id: givenID, FreeSeats: resSeats.FreeSeats})
	rsp.Id = givenID

	return nil
}

func (shs *SService) DeleteShow(ctx context.Context, req *show.DeleteShowRequest, rsp *show.DeleteShowResult) error {
	for i, v := range shs.show {
		if v.Id == req.Id {
			shs.delete(i, v.Id)
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

func (shs *SService) FromHallDelete(ctx context.Context, req *show.DeleteShowOfHallRequest, rsp *show.DeleteShowOfHallResult) error {
	success := false
	//Got the Id of an Hall which no longer exists
	for i, v := range shs.show {
		if v.HallID == req.HallID {
			shs.delete(i, v.Id)
			success = true
		}
	}
	rsp.Successful = success
	return nil
}

func (shs *SService) FromMovieDelete(ctx context.Context, req *show.DeleteShowOfMovieRequest, rsp *show.DeleteShowOfMovieResult) error {
	success := false
	//Got the Id of an Hall which no longer exists
	for i, v := range shs.show {
		if v.MovieID == req.MovieID {
			shs.delete(i, v.Id)
			success = true
		}
	}
	rsp.Successful = success
	return nil
}

func (shs *SService) AskSeats(ctx context.Context, req *show.FreeSeatsRequest, rsp *show.FreeSeatsResult) error {
	for _, s := range shs.show {
		if s.Id == req.ShowID {
			rsp.FreeSeats = s.FreeSeats
			return nil
		}
	}
	rsp.FreeSeats = -1
	return nil
}

func (shs *SService) UpdateSeats(ctx context.Context, req *show.UpdateSeatsRequest, rsp *show.UpdateSeatsResult) error {
	for _, s := range shs.show {
		if s.Id == req.ShowID {
			s.FreeSeats = s.FreeSeats - req.AmountSeats
			rsp.Success = true
			return nil
		}
	}
	rsp.Success = false
	return nil
}

func (shs *SService) GetShowList(ctx context.Context, req *show.GetShowListRequest, rsp *show.GetShowListResult) error {
	rsp.Shows = shs.show
	return nil
}

func (shs *SService) delete(index int, showID int32) {
	shs.show = append(shs.show[:index], shs.show[index+1:]...)

	var client client.Client
	bookingC := booking.NewBookingService("go.micro.services.booking", client)

	_, err := bookingC.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: showID})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.show"),
		micro.Address(fmt.Sprintf(":%v", 1035)),
	)

	service.Init()
	show.RegisterShowHandler(service.Server(), &SService{show: exampleData(), nextID: 5})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func exampleData() []*show.ShowData {
	shows := make([]*show.ShowData, 0)
	shows = append(shows, &show.ShowData{Id: 1, MovieID: 1, HallID: 2, FreeSeats: 482})
	shows = append(shows, &show.ShowData{Id: 2, MovieID: 2, HallID: 1, FreeSeats: 35})
	shows = append(shows, &show.ShowData{Id: 3, MovieID: 3, HallID: 2, FreeSeats: 482})
	shows = append(shows, &show.ShowData{Id: 4, MovieID: 4, HallID: 1, FreeSeats: 33})
	return shows
}
