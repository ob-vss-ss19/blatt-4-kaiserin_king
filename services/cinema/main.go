package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

type CService struct {
	cHall 	[]*cinema.CinemaHall
	nextID 	int32
	me	micro.Service
}

func (cs *CService) CreateHall(ctx context.Context, req *cinema.CreateHallRequest, rsp *cinema.CreateHallResult) error {
	givenID := cs.nextID
	cs.nextID++
	cs.cHall = append(cs.cHall, &cinema.CinemaHall{Name: req.Name, Rows: req.Rows, Cols: req.Cols, Id: givenID})
	rsp.Id = givenID

	// Create new greeter client
	greeter := user.NewUserService("go.micro.services.user", cs.me.Client())

	// Call the greeter
	rsp2, err := greeter.CreateUser(context.TODO(), &user.CreateUserRequest{Name: "test"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp2.Id)

	return nil
}

func (cs *CService) DeleteHall(ctx context.Context, req *cinema.DeleteHallRequest, rsp *cinema.DeleteHallResult) error {
	// check if there are bookings for given id / hall
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
	cinema.RegisterCinemaHandler(service.Server(), &CService{cHall: make([]*cinema.CinemaHall, 0), nextID: 0, me: service})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}