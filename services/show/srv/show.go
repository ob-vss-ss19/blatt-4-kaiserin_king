package srv

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

type SService struct {
	Show   []*show.ShowData
	NextID int32
	mux    sync.Mutex
}

//Function that creats a new show requested with an HallId and MovieID.
//If both ID's exist the function returns a valid showID, else it returns -1.
func (shs *SService) CreateShow(ctx context.Context, req *show.CreateShowRequest, rsp *show.CreateShowResult) error {
	if shs.movieExist(req.MovieID) && shs.hallExist(req.HallID) {
		shs.mux.Lock()
		givenID := shs.NextID
		shs.mux.Unlock()
		shs.NextID++
		var client client.Client
		cinemaC := cinema.NewCinemaService("go.micro.services.cinema", client)

		resSeats, err := cinemaC.AskSeats(context.TODO(), &cinema.FreeSeatsRequest{HallID: req.HallID})
		if err != nil {
			fmt.Println(err)
		}

		shs.Show = append(shs.Show,
			&show.ShowData{HallID: req.HallID, MovieID: req.MovieID, Id: givenID, FreeSeats: resSeats.FreeSeats})
		rsp.Id = givenID
		return nil
	}
	rsp.Id = -1
	return nil
}

//Function to delete a show given by its ID.
//Return if the operation was successful by bool value.
func (shs *SService) DeleteShow(ctx context.Context, req *show.DeleteShowRequest, rsp *show.DeleteShowResult) error {
	for i, v := range shs.Show {
		if v.Id == req.Id {
			shs.mux.Lock()
			shs.delete(i, v.Id)
			shs.mux.Unlock()
			rsp.Successful = true
			return nil
		}
	}
	rsp.Successful = false
	return nil
}

//Function that get called if a hall got deleted, existing shows in this hall have to be removed.
//Gets the Id of the deleted hall and returns if the operation was successful by bool value.
func (shs *SService) FromHallDelete(ctx context.Context, req *show.DeleteShowOfHallRequest,
	rsp *show.DeleteShowOfHallResult) error {
	success := false
	//Got the Id of an Hall which no longer exists
	for i, v := range shs.Show {
		if v.HallID == req.HallID {
			shs.mux.Lock()
			shs.delete(i, v.Id)
			shs.mux.Unlock()
			success = true
		}
	}
	rsp.Successful = success
	return nil
}

//Function that get called if a movie got deleted, existing shows which are using this movie have to be removed.
//Gets the Id of the deleted movie and returns if the operation was successful by bool value.
func (shs *SService) FromMovieDelete(ctx context.Context, req *show.DeleteShowOfMovieRequest,
	rsp *show.DeleteShowOfMovieResult) error {
	success := false
	//Got the Id of an Hall which no longer exists
	for i, v := range shs.Show {
		if v.MovieID == req.MovieID {
			shs.mux.Lock()
			shs.delete(i, v.Id)
			shs.mux.Unlock()
			success = true
		}
	}
	rsp.Successful = success
	return nil
}

//Function that return the amount of current free seats in show, requested with the showID.
func (shs *SService) AskSeats(ctx context.Context, req *show.FreeSeatsRequest, rsp *show.FreeSeatsResult) error {
	for _, s := range shs.Show {
		if s.Id == req.ShowID {
			rsp.FreeSeats = s.FreeSeats
			return nil
		}
	}
	rsp.FreeSeats = -1
	return nil
}

//Function that updates amount of current free seats in show, requested with the showID and amount of booked seats.
func (shs *SService) UpdateSeats(ctx context.Context, req *show.UpdateSeatsRequest, rsp *show.UpdateSeatsResult) error {
	for _, s := range shs.Show {
		if s.Id == req.ShowID {
			shs.mux.Lock()
			s.FreeSeats -= req.AmountSeats
			shs.mux.Unlock()
			rsp.Success = true
			return nil
		}
	}
	rsp.Success = false
	return nil
}

//Function that returns the list of all shows.
func (shs *SService) GetShowList(ctx context.Context, req *show.GetShowListRequest, rsp *show.GetShowListResult) error {
	rsp.Shows = shs.Show
	return nil
}

//Function that returns if a movie given by its ID does exist.
func (shs *SService) movieExist(movieID int32) bool {
	var client client.Client
	movieC := movie.NewMovieService("go.micro.services.movie", client)

	rsp, err := movieC.Exist(context.TODO(), &movie.ExistRequest{Id: movieID})

	if err != nil {
		fmt.Println(err)
	}

	return rsp.Exist
}

//Function that returns if a hall given by its ID does exist.
func (shs *SService) hallExist(hallID int32) bool {
	var client client.Client
	cinemaC := cinema.NewCinemaService("go.micro.services.cinema", client)

	rsp, err := cinemaC.Exist(context.TODO(), &cinema.ExistRequest{Id: hallID})

	if err != nil {
		fmt.Println(err)
	}

	return rsp.Exist
}

//Function that returns if a Show given by its ID does exist.
func (shs *SService) Exist(ctx context.Context, req *show.ExistRequest, rsp *show.ExistResult) error {
	for _, s := range shs.Show {
		if s.Id == req.Id {
			rsp.Exist = true
			return nil
		}
	}
	rsp.Exist = false
	return nil
}

//Function that deletes a show from the list and informs the bookingservice about the delete.
func (shs *SService) delete(index int, showID int32) {
	shs.Show = append(shs.Show[:index], shs.Show[index+1:]...)

	var client client.Client
	bookingC := booking.NewBookingService("go.micro.services.booking", client)

	_, err := bookingC.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: showID})
	if err != nil {
		fmt.Println(err)
	}
}

func RunService() {
	service := micro.NewService(
		micro.Name("go.micro.services.show"),
		micro.Address(fmt.Sprintf(":%v", 1035)),
	)

	service.Init()
	err := show.RegisterShowHandler(service.Server(), &SService{Show: ExampleData(), NextID: 5})
	if err != nil {
		fmt.Println(err)
	}
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}

//Example Data of show which is added to Service from Beginn
func ExampleData() []*show.ShowData {
	shows := make([]*show.ShowData, 0)
	shows = append(shows, &show.ShowData{Id: 1, MovieID: 1, HallID: 2, FreeSeats: 482})
	shows = append(shows, &show.ShowData{Id: 2, MovieID: 2, HallID: 1, FreeSeats: 35})
	shows = append(shows, &show.ShowData{Id: 3, MovieID: 3, HallID: 2, FreeSeats: 482})
	shows = append(shows, &show.ShowData{Id: 4, MovieID: 4, HallID: 1, FreeSeats: 33})
	return shows
}
