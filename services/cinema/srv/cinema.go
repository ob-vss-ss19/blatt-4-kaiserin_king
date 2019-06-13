package srv

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

type CService struct {
	CHall  []*cinema.CinemaHall
	NextID int32
	mux    sync.Mutex
}

//Function to create a new Cinemahall by request with an name for the hall.
//Returns the id of the created hall
func (cs *CService) CreateHall(ctx context.Context, req *cinema.CreateHallRequest, rsp *cinema.CreateHallResult) error {
	cs.mux.Lock()
	givenID := cs.NextID
	cs.NextID++
	cs.mux.Unlock()
	cs.CHall = append(cs.CHall, &cinema.CinemaHall{Name: req.Name, Rows: req.Rows, Cols: req.Cols, Id: givenID})

	rsp.Id = givenID

	return nil
}

//Function to delete a Hell requested by the ID.
//Returnsn if the operation is successful by bool value.
func (cs *CService) DeleteHall(ctx context.Context, req *cinema.DeleteHallRequest, rsp *cinema.DeleteHallResult) error {
	// check if there are bookings for given id / hall
	for i, h := range cs.CHall {
		if h.Id == req.Id {

			//Send HallID to Showservice to delete all shows
			var client client.Client
			showC := show.NewShowService("go.micro.services.show", client)
			cs.mux.Lock()
			_, err := showC.FromHallDelete(context.TODO(), &show.DeleteShowOfHallRequest{HallID: req.Id})
			if err != nil {
				fmt.Println(err)
			}
			//delete Hall from CinemaService
			cs.CHall = append(cs.CHall[:i], cs.CHall[i+1:]...)
			cs.mux.Unlock()
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

//Function that returns the amount of seats that the hall given by its ID has.
func (cs *CService) AskSeats(ctx context.Context, req *cinema.FreeSeatsRequest, rsp *cinema.FreeSeatsResult) error {
	for _, h := range cs.CHall {
		if h.Id == req.HallID {
			rsp.FreeSeats = h.Cols * h.Rows
			return nil
		}
	}
	rsp.FreeSeats = -1
	return nil
}

//Function that return List of all halls.
func (cs *CService) GetHallList(ctx context.Context, req *cinema.GetHallListRequest,
	rsp *cinema.GetHallListResult) error {
	rsp.CHall = cs.CHall
	return nil
}

//Function that returns if an hall given by its ID exist.
func (cs *CService) Exist(ctx context.Context, req *cinema.ExistRequest, rsp *cinema.ExistResult) error {
	for _, c := range cs.CHall {
		if c.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

func RunService() {
	service := micro.NewService(
		micro.Name("go.micro.services.cinema"),
		micro.Address(fmt.Sprintf(":%v", 1036)),
	)

	service.Init()

	err := cinema.RegisterCinemaHandler(service.Server(), &CService{CHall: ExampleData(), NextID: 3})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

//Returns example Data of cinemahalls used in der Service
func ExampleData() []*cinema.CinemaHall {
	halls := make([]*cinema.CinemaHall, 0)
	halls = append(halls, &cinema.CinemaHall{Id: 1, Name: "Kino 9", Rows: 5, Cols: 7})
	halls = append(halls, &cinema.CinemaHall{Id: 2, Name: "Kino 11", Rows: 18, Cols: 26})
	return halls
}
