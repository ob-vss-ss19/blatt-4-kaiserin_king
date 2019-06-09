package main

import "C"
import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

type CService struct {
	cHall  []*cinema.CinemaHall
	nextID int32
}

func (cs *CService) CreateHall(ctx context.Context, req *cinema.CreateHallRequest, rsp *cinema.CreateHallResult) error {
	givenID := cs.nextID
	cs.nextID++
	cs.cHall = append(cs.cHall, &cinema.CinemaHall{Name: req.Name, Rows: req.Rows, Cols: req.Cols, Id: givenID})
	rsp.Id = givenID

	return nil
}

func (cs *CService) DeleteHall(ctx context.Context, req *cinema.DeleteHallRequest, rsp *cinema.DeleteHallResult) error {
	// check if there are bookings for given id / hall
	for i, h := range cs.cHall {
		if h.Id == req.Id {

			//Send HallID to Showservice to delete all shows
			var client client.Client
			showC := show.NewShowService("go.micro.services.show", client)

			_, err := showC.FromHallDelete(context.TODO(), &show.DeleteShowOfHallRequest{HallID: req.Id})
			if err != nil {
				fmt.Println(err)
			}
			//delete Hall from CinemaService
			cs.cHall = append(cs.cHall[:i], cs.cHall[i+1:]...)
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

func (cs *CService) AskSeats(ctx context.Context, req *cinema.FreeSeatsRequest, rsp *cinema.FreeSeatsResult) error {
	for _, h := range cs.cHall {
		if h.Id == req.HallID {
			rsp.FreeSeats = h.Cols * h.Rows
			return nil
		}
	}
	rsp.FreeSeats = -1
	return nil
}

func (cs *CService) GetHallList(ctx context.Context, req *cinema.GetHallListRequest, rsp *cinema.GetHallListResult) error {
	rsp.CHall = cs.cHall
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.cinema"),
	)

	service.Init()
	cinema.RegisterCinemaHandler(service.Server(), &CService{cHall: exampleData(), nextID: 3})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func exampleData() []*cinema.CinemaHall {
	halls := make([]*cinema.CinemaHall, 0)
	halls = append(halls, &cinema.CinemaHall{Id: 1, Name: "Kino 9", Rows: 5, Cols: 7})
	halls = append(halls, &cinema.CinemaHall{Id: 2, Name: "Kino 11", Rows: 18, Cols: 26})
	return halls
}
