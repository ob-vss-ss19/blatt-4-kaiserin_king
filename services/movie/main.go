package main


import (
	"log"
	"context"

	"github.com/micro/go-micro"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
)

type MService struct {
	movie 	[]*movie.MovieData
	nextID	int32
}

func (ms *MService) CreateMovie(ctx context.Context, req *movie.CreateMovieRequest, rsp *movie.CreateMovieResult) error {
	givenID := ms.nextID
	ms.nextID++
	ms.movie = append(ms.movie, &movie.MovieData{Titel: req.Titel, Id: givenID})
	rsp.Id = givenID

	return nil
}

func (ms *MService) DeleteMovie(ctx context.Context, req *movie.DeleteMovieRequest, rsp *movie.DeleteMovieResult) error {
	// check if movie is used for bookings or shows
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.movie"),
	)

	service.Init()
	movie.RegisterMovieHandler(service.Server(), &MService{movie: make([]*movie.MovieData, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

