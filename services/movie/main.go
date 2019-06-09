package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	"log"

	"github.com/micro/go-micro"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
)

type MService struct {
	movie  []*movie.MovieData
	nextID int32
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
	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)

	_, err := showC.FromMovieDelete(context.TODO(), &show.DeleteShowOfMovieRequest{MovieID: req.Id})
	if err != nil {
		fmt.Println(err)
	}
	//delete Movie from MovieService
	for i, v := range ms.movie {
		if v.Id == req.Id {
			ms.movie = append(ms.movie[:i], ms.movie[i+1:]...)
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

func (ms *MService) GetMovieList(ctx context.Context, req *movie.GetMovieListRequest, rsp *movie.GetMovieListResult) error {
	rsp.Movies = ms.movie
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.movie"),
	)

	service.Init()
	movie.RegisterMovieHandler(service.Server(), &MService{movie: exampleData(), nextID: 5})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

func exampleData() []*movie.MovieData {
	movies := make([]*movie.MovieData, 0)
	movies = append(movies, &movie.MovieData{Id: 1, Titel: "Deadpool"})
	movies = append(movies, &movie.MovieData{Id: 2, Titel: "Deadpool 2"})
	movies = append(movies, &movie.MovieData{Id: 3, Titel: "Avengers 4"})
	movies = append(movies, &movie.MovieData{Id: 4, Titel: "Ted"})
	return movies
}
