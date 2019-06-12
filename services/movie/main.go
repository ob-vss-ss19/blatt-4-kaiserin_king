package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

type MService struct {
	movie  []*movie.MovieData
	nextID int32
	mux    sync.Mutex
}

//Function to create a new Movie, requested with a name for the movie.
//Returns ID of the created movie.
func (ms *MService) CreateMovie(ctx context.Context, req *movie.CreateMovieRequest,
	rsp *movie.CreateMovieResult) error {
	ms.mux.Lock()
	givenID := ms.nextID
	ms.nextID++
	ms.mux.Unlock()
	ms.movie = append(ms.movie, &movie.MovieData{Titel: req.Titel, Id: givenID})
	rsp.Id = givenID

	return nil
}

//Function that deletes a movie which is requested by the ID.
//Returns if operation is successful by bool value.
func (ms *MService) DeleteMovie(ctx context.Context, req *movie.DeleteMovieRequest,
	rsp *movie.DeleteMovieResult) error {
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
			ms.mux.Lock()
			ms.movie = append(ms.movie[:i], ms.movie[i+1:]...)
			ms.mux.Unlock()
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

//Function that return List of all movies.
func (ms *MService) GetMovieList(ctx context.Context, req *movie.GetMovieListRequest,
	rsp *movie.GetMovieListResult) error {
	rsp.Movies = ms.movie
	return nil
}

//Function that return if a  movie, given by its ID does exist.
func (ms *MService) Exist(ctx context.Context, req *movie.ExistRequest, rsp *movie.ExistResult) error {
	for _, m := range ms.movie {
		if m.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.movie"),
		micro.Address(fmt.Sprintf(":%v", 1034)),
	)

	service.Init()
	err := movie.RegisterMovieHandler(service.Server(), &MService{movie: exampleData(), nextID: 5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

//Example Data of movies which is added to the Service
func exampleData() []*movie.MovieData {
	movies := make([]*movie.MovieData, 0)
	movies = append(movies, &movie.MovieData{Id: 1, Titel: "Deadpool"})
	movies = append(movies, &movie.MovieData{Id: 2, Titel: "Deadpool 2"})
	movies = append(movies, &movie.MovieData{Id: 3, Titel: "Avengers 4"})
	movies = append(movies, &movie.MovieData{Id: 4, Titel: "Ted"})
	return movies
}
