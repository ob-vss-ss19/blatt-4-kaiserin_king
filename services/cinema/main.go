package cinema

import (
	"context"
	"log"
)

type CService struct {
	cHall 	[]*cinema.CinemaHall
	nextID 	int32
}

func (cs *CService) CreateHall(ctx context.Context, req *cinema.CreateHallRequest, rsp *cinema.CreateHallResult) error {
	givenID := cs.nextID
	cs.nextID++
	cs.cHall = append(cs.halls, &cinema.CinemaHall{Name: req.Name, Rows: req.Rows, Cols: req.Cols, Id: givenID})
	rsp.Id = givenID

	return nil
}

func (cs *CService) DeleteHall(ctx context.Context, req *cinema.DeleteHallRequest, rsp *cinema.DeleteHallResult) error {
	// check if there are bookings for given id / hall
	return nil
}

func (cs *CService) GetHallList(ctx context.Context, req *cinema.GetHallListRequest, rsp *cinema.GetHallListResult) error {
	rsp.cHall = cs.cHall
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.user"),
	)

	service.Init()
	proto.RegisterCinemaHandler(service.Server(), &CService{cHall: make([]*cinema.CinemaHall, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}