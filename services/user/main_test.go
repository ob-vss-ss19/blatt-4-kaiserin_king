package main

import (
	"context"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
	"testing"


)

func TestShow(t *testing.T) {
	service := UService{user: make([]*user.UserData, 0) , nextID: 1}

	rsp := &user.GetUserListResult{}
	err := service.GetUserList(context.TODO(), &user.GetUserListRequest{}, rsp)
	if err == nil {
		if len(rsp.Users) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspCreate := &user.CreateUserResult{}
	err = service.CreateUser(context.TODO(), &user.CreateUserRequest{Name: "Max Mustermann"}, rspCreate)

	rsp = &user.GetUserListResult{}
	err = service.GetUserList(context.TODO(), &user.GetUserListRequest{}, rsp)
	if err == nil {
		if len(rsp.Users) != 1 {
			t.Error("Expected List with one Element")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete := &user.DeleteUserResult{}
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if !rspDelete.Successful {
			t.Error("Expected successful deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

	rsp = &user.GetUserListResult{}
	err = service.GetUserList(context.TODO(), &user.GetUserListRequest{}, rsp)
	if err == nil {
		if len(rsp.Users) != 0 {
			t.Error("Expected empty List")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete = &user.DeleteUserResult{}
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: rspCreate.Id}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

	rspDelete = &user.DeleteUserResult{}
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: 1}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

}