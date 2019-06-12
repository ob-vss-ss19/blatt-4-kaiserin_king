package main

import (
	"context"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
	"testing"
)

func TestShow(t *testing.T) {
	service := SService{show: make([]*show.ShowData,0), nextID:1}

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
	err = service.CreateShow(context.TODO(), &show.CreateShowRequest{MovieID: 10, HallID:1}, rspCreate)
	if err == nil {
		if rspCreate.Id == -1 {
			t.Error("Expected successful creation")
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

}