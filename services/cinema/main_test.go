package main

import (
	"context"
	"testing"

	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
)

func TestCinemaGetList(t *testing.T) {
	service := CService{cHall: make([]*cinema.CinemaHall, 0), nextID: 1}
	rsp := &cinema.GetHallListResult{}
	err := service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 0 {
			t.Errorf("Expected empty hall list!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspCreate := &cinema.CreateHallResult{}
	_ = service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 1, Cols: 1}, rspCreate)
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 1 {
			t.Errorf("Expected hall list with one entry!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete := &cinema.DeleteHallResult{}
	_ = service.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id: rspCreate.Id}, rspDelete)
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if !rspDelete.Successful {
			t.Errorf("Expected successful delete!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rsp = &cinema.GetHallListResult{}
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 0 {
			t.Errorf("Expected empty hall list!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}
}

func TestCinemaDeleteWrongID(t *testing.T) {
	service := CService{cHall: make([]*cinema.CinemaHall, 0), nextID: 1}

	rspDelete := &cinema.DeleteHallResult{}
	err := service.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id: 42}, rspDelete)

	if err == nil {
		if rspDelete.Successful {
			t.Errorf("Expected failing delete!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

}

func TestAskForSeats(t *testing.T) {
	service := CService{cHall: make([]*cinema.CinemaHall, 0), nextID: 1}

	rspCreate := &cinema.CreateHallResult{}
	_ = service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 10, Cols: 3}, rspCreate)

	rspSeats := &cinema.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &cinema.FreeSeatsRequest{HallID: rspCreate.Id}, rspSeats)

	if rspSeats.FreeSeats != 30 {
		t.Errorf("Expected 30 seats but got %v", rspSeats.FreeSeats)
	}

	rspSeats = &cinema.FreeSeatsResult{}
	_ = service.AskSeats(context.TODO(), &cinema.FreeSeatsRequest{HallID: 200}, rspSeats)

	if rspSeats.FreeSeats != -1 {
		t.Errorf("Expected 30 seats but got %v", rspSeats.FreeSeats)
	}
}

func TestAskExist(t *testing.T) {
	service := CService{cHall: make([]*cinema.CinemaHall, 0), nextID: 1}

	rspCreate := &cinema.CreateHallResult{}
	_ = service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 10, Cols: 3}, rspCreate)

	rspExist := &cinema.ExistResult{}
	_ = service.Exist(context.TODO(), &cinema.ExistRequest{Id: rspCreate.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected existing hall")
	}

	rspExist = &cinema.ExistResult{}
	_ = service.Exist(context.TODO(), &cinema.ExistRequest{Id: 200}, rspExist)

	if rspExist.Exist {
		t.Error("Expected not existing hall")
	}

}
