package main

import (
	"context"
	"fmt"
	"testing"

	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	"github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/srv"
	cs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/srv"
	ms "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/srv"
	shs "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/srv"
	us "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/srv"
)

func TestBookingGetList(t *testing.T) {
	cs.RunService()
	ms.RunService()
	shs.RunService()
	us.RunService()

	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}
	rsp := &booking.GetListResult{}
	err := service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected empty list! \n")
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)
	if (rspCreate.Id) == -1 {
		t.Error("Adding booking failed! \n")
	}
	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)
	err = service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}
}

func TestBookingTooMuchSeats(t *testing.T) {
	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(),
		&booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 500},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Errorf("Expected no booking because of too much seats!\n")
	}
}

func TestBookingTwoConfirmed(t *testing.T) {
	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(),
		&booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 200},
		rspCreate)

	rspCreate2 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(),
		&booking.CreateBookingRequest{UserID: 2, ShowID: 1, Seats: 400},
		rspCreate2)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)
	if !rspConfirm.Successful {
		t.Errorf("Booking expected successful")
	}

	rspConfirm2 := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate2.Id}, rspConfirm2)
	if rspConfirm2.Successful {
		t.Errorf("Booking 2 expected not successful")
	}

	rspConfirm3 := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: 42}, rspConfirm3)
	if rspConfirm3.Successful {
		t.Errorf("Booking 3 expected not successful")
	}
}

func TestBookingDelete(t *testing.T) {
	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}
	rsp := &booking.GetListResult{}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)

	rspDelete := &booking.DeleteBookingResult{}
	_ = service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id: rspCreate.Id}, rspDelete)

	err := service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rsp = &booking.GetListResult{}

	rspCreate = &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)

	rspDelete = &booking.DeleteBookingResult{}
	_ = service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id: rspCreate.Id}, rspDelete)

	err = service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspDelete = &booking.DeleteBookingResult{}
	_ = service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id: 6000}, rspDelete)
	if rspDelete.Successful {
		t.Error("Expected no booking with this ID!")
	}

}

func TestCreateWrongID(t *testing.T) {
	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 200, ShowID: 1, Seats: 1},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Error("Expected no booking")
	}

	rspCreate = &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 200, Seats: 1},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Error("Expected no booking")
	}

	rspCreate = &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 200, Seats: 200},
		rspCreate)

	if rspCreate.Id != -1 {
		t.Error("Expected no booking")
	}
}

func TestFromShowDelete(t *testing.T) {
	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate1 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate1)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate1.Id}, rspConfirm)

	rspCreate2 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 2, ShowID: 2, Seats: 1}, rspCreate2)

	rsp := &booking.GetListResult{}
	err := service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one element! \n")
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 1 {
			t.Error("Expected list with one entry! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspDelete := &booking.FromShowDeleteResult{}
	_ = service.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: 1}, rspDelete)

	err = service.GetBookingList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected empty list! \n")
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	rspDelete = &booking.FromShowDeleteResult{}
	_ = service.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: 2}, rspDelete)

	if !rspDelete.Successful {
		t.Error("Expected successful deleting!")
	}

	err = service.GetNotConfirmedList(context.TODO(), &booking.GetListRequest{}, rsp)
	if err == nil {
		if len(rsp.Bookings) != 0 {
			t.Error("Expected empty list! \n")
			println(len(rsp.Bookings))
			fmt.Printf("%v\n", rsp.Bookings)
		}
	} else {
		t.Errorf("Error for List Request: %v \n", err)
	}

	_ = service.FromShowDelete(context.TODO(), &booking.FromShowDeleteRequest{Id: 200}, rspDelete)

	if rspDelete.Successful {
		t.Error("Expected failing deleting!")
	}
}

func TestExist(t *testing.T) {
	service := srv.BService{Booking: make([]*booking.BookingData, 0),
		NotConfirmed: make([]*booking.BookingData, 0),
		NextID:       1}

	rspCreate1 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate1)

	rspConfirm := &booking.ConfirmBookingResult{}
	_ = service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate1.Id}, rspConfirm)

	rspCreate2 := &booking.CreateBookingResult{}
	_ = service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 2, ShowID: 2, Seats: 1}, rspCreate2)

	rspExist := &booking.ExistResult{}
	_ = service.Exist(context.TODO(), &booking.ExistRequest{Id: rspCreate1.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected that booking excists")
	}

	rspExist = &booking.ExistResult{}
	_ = service.Exist(context.TODO(), &booking.ExistRequest{Id: rspCreate2.Id}, rspExist)

	if !rspExist.Exist {
		t.Error("Expected that booking excists")
	}

	rspExist = &booking.ExistResult{}
	_ = service.Exist(context.TODO(), &booking.ExistRequest{Id: 200}, rspExist)

	if rspExist.Exist {
		t.Error("Expected that booking not exists")
	}

}
