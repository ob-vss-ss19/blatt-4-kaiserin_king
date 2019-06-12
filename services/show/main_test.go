package main

import (
	"context"
	"testing"

	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	"github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/srv"
)

func TestShow(t *testing.T) {
	service := srv.SService{Show: make([]*show.ShowData, 0), NextID: 1}

	rsp := &show.GetShowListResult{}
	err := service.GetShowList(context.TODO(), &show.GetShowListRequest{}, rsp)
	if err == nil {
		if len(rsp.Shows) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspCreate := &show.CreateShowResult{}
	err = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 1, HallID: 1}, rspCreate)
	if err == nil {
		if rspCreate.Id == -1 {
			t.Error("Expected successful creation")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspSeats := &show.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: rspCreate.Id}, rspSeats)
	if rspSeats.FreeSeats != 35 {
		t.Errorf("Expected 35 free seats but got %v", rspSeats.FreeSeats)
	}

	rspUpdate := &show.UpdateSeatsResult{}
	_ = service.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: rspCreate.Id, AmountSeats: 5}, rspUpdate)
	if !rspUpdate.Success {
		t.Errorf("Expected successful update")
	}

	rspUpdate = &show.UpdateSeatsResult{}
	_ = service.UpdateSeats(context.TODO(), &show.UpdateSeatsRequest{ShowID: 200, AmountSeats: 5}, rspUpdate)
	if rspUpdate.Success {
		t.Errorf("Expected failing update")
	}

	rspSeats = &show.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: rspCreate.Id}, rspSeats)
	if rspSeats.FreeSeats != 30 {
		t.Errorf("Expected 30 free seats but got %v", rspSeats.FreeSeats)
	}

	rspSeats = &show.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &show.FreeSeatsRequest{ShowID: 200}, rspSeats)
	if rspSeats.FreeSeats != -1 {
		t.Errorf("Expected 30 free seats but got %v", rspSeats.FreeSeats)
	}

	rspCreate2 := &show.CreateShowResult{}
	err = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 200, HallID: 200}, rspCreate2)
	if err == nil {
		if rspCreate2.Id != -1 {
			t.Error("Expected failing creation")
		}
	} else {
		t.Error("Error with Request!")
	}

	rsp = &show.GetShowListResult{}
	err = service.GetShowList(context.TODO(), &show.GetShowListRequest{}, rsp)
	if err == nil {
		if len(rsp.Shows) != 1 {
			t.Error("Expected List with one Element")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete := &show.DeleteShowResult{}
	err = service.DeleteShow(context.TODO(), &show.DeleteShowRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if !rspDelete.Successful {
			t.Error("Expected successful deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

	rsp = &show.GetShowListResult{}
	err = service.GetShowList(context.TODO(), &show.GetShowListRequest{}, rsp)
	if err == nil {
		if len(rsp.Shows) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete = &show.DeleteShowResult{}
	err = service.DeleteShow(context.TODO(), &show.DeleteShowRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing delete")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspCreate = &show.CreateShowResult{}
	_ = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 1, HallID: 1}, rspCreate)

	rspDeleteHall := &show.DeleteShowOfHallResult{}
	_ = service.FromHallDelete(context.TODO(), &show.DeleteShowOfHallRequest{HallID: 1}, rspDeleteHall)
	if !rspDeleteHall.Successful {
		t.Error("Expected successful deleting from Hall")
	}

	rspDeleteHall = &show.DeleteShowOfHallResult{}
	_ = service.FromHallDelete(context.TODO(), &show.DeleteShowOfHallRequest{HallID: 200}, rspDeleteHall)
	if rspDeleteHall.Successful {
		t.Error("Expected failing deleting from Hall")
	}

	rspCreate = &show.CreateShowResult{}
	_ = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 1, HallID: 1}, rspCreate)

	rspExist := &show.ExistResult{}
	_ = service.Exist(context.TODO(), &show.ExistRequest{Id: rspCreate.Id}, rspExist)
	if !rspExist.Exist {
		t.Error("Expected that show exists")
	}

	rspDeleteMovie := &show.DeleteShowOfMovieResult{}
	_ = service.FromMovieDelete(context.TODO(), &show.DeleteShowOfMovieRequest{MovieID: 1}, rspDeleteMovie)
	if !rspDeleteMovie.Successful {
		t.Error("Expected successful deleting from Movie")
	}

	rspDeleteMovie = &show.DeleteShowOfMovieResult{}
	_ = service.FromMovieDelete(context.TODO(), &show.DeleteShowOfMovieRequest{MovieID: 200}, rspDeleteMovie)
	if rspDeleteMovie.Successful {
		t.Error("Expected failing deleting from Movie")
	}

	rspExist = &show.ExistResult{}
	_ = service.Exist(context.TODO(), &show.ExistRequest{Id: rspCreate.Id}, rspExist)
	if rspExist.Exist {
		t.Error("Expected that show does not exists")
	}

}
