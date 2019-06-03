package main


import (
	"log"
	"context"

	"github.com/micro/go-micro"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

type SService struct {
	show 	[]*show.ShowData
	nextID	int32
}

func (shs *SService) CreateShow(ctx context.Context, req *show.CreateShowRequest, rsp *show.CreateShowResult) error {
	givenID := shs.nextID
	shs.nextID++
	shs.show = append(shs.show, &show.ShowData{HallID: req.HallID, MovieID: req.MovieID, Id: givenID})
	rsp.Id = givenID

	return nil
}

func (shs *SService) DeleteShow(ctx context.Context, req *show.DeleteShowRequest, rsp *show.DeleteShowResult) error {
	// delete movie and hall ?
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.show"),
	)

	service.Init()
	show.RegisterShowHandler(service.Server(), &SService{show: make([]*show.ShowData, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

