package main

import (
	"context"
	cinema "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/proto"
	"testing"
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
	service.CreateHall(context.TODO(), &cinema.CreateHallRequest{Name: "Kino1", Rows: 1, Cols: 1}, rspCreate)
	err = service.GetHallList(context.TODO(), &cinema.GetHallListRequest{}, rsp)

	if err == nil {
		if len(rsp.CHall) != 1 {
			t.Errorf("Expected hall list with one entry!")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete := &cinema.DeleteHallResult{}
	service.DeleteHall(context.TODO(), &cinema.DeleteHallRequest{Id:rspCreate.Id}, rspDelete)
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
