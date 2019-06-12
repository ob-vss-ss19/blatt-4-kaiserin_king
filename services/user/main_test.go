package main

import (
	"context"
	"testing"

	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

func TestShow(t *testing.T) {
	service := UService{user: make([]*user.UserData, 0), nextID: 1}

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
	_ = service.CreateUser(context.TODO(), &user.CreateUserRequest{Name: "Max Mustermann"}, rspCreate)

	rspExist := &user.ExistResult{}
	_ = service.Exist(context.TODO(), &user.ExistRequest{Id: rspCreate.Id}, rspExist)
	if !rspExist.Exist {
		t.Error("Expected user to exist")
	}

	rspExist = &user.ExistResult{}
	_ = service.Exist(context.TODO(), &user.ExistRequest{Id: 200}, rspExist)
	if rspExist.Exist {
		t.Error("Expected user not to exist")
	}

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
	err = service.DeleteUser(context.TODO(), &user.DeleteUserRequest{Id: 200}, rspDelete)
	if err == nil {
		if rspDelete.Successful {
			t.Error("Expected failing deleting")
		}
	} else {
		t.Error("Error with Request!")
	}

}

func TestShowDeleting(t *testing.T) {
	service := UService{user: exampleData(), nextID: 1}

	rspDeleteBooking := &user.BookingDeletedResult{}
	_ = service.BookingDeleted(context.TODO(), &user.BookingDeletedRequest{UserID: 4, BookingID: 2}, rspDeleteBooking)
	if !service.CheckBookingOfUser(4) {
		t.Error("Expected no bookings for id=4")
	}

	rspCreate := &user.CreateUserResult{}
	_ = service.CreateUser(context.TODO(), &user.CreateUserRequest{Name: "Max Mustermann"}, rspCreate)
	if !service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected no bookings for id=%v", rspCreate.Id)
	}

	rspMarkedB := &user.CreatedBookingResult{}
	_ = service.CreatedMarkedBooking(context.TODO(),
		&user.CreatedBookingRequest{UserID: rspCreate.Id, BookingID: 200},
		rspMarkedB)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	rspMarkedB = &user.CreatedBookingResult{}
	_ = service.CreatedBooking(context.TODO(), &user.CreatedBookingRequest{UserID: 700, BookingID: 200}, rspMarkedB)
	if !service.CheckBookingOfUser(700) {
		t.Error("Expected no bookings for id=700")
	}

	rspMarkedB = &user.CreatedBookingResult{}
	_ = service.CreatedBooking(context.TODO(), &user.CreatedBookingRequest{UserID: rspCreate.Id, BookingID: 200},
		rspMarkedB)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	_ = service.CreatedMarkedBooking(context.TODO(),
		&user.CreatedBookingRequest{UserID: rspCreate.Id, BookingID: 700},
		rspMarkedB)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

	rspDeleteBooking = &user.BookingDeletedResult{}
	_ = service.BookingDeleted(context.TODO(),
		&user.BookingDeletedRequest{UserID: rspCreate.Id, BookingID: 700},
		rspDeleteBooking)
	if service.CheckBookingOfUser(rspCreate.Id) {
		t.Errorf("Expected bookings for id=%v", rspCreate.Id)
	}

}
