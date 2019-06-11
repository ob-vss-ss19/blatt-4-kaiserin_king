package main

import (
	"context"
	"fmt"
	booking "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/booking/proto"
	"testing"
)

func TestBookingGetList(t *testing.T) {
	service := BService{booking: make([]*booking.BookingData, 0), notConfirmed: make([]*booking.BookingData, 0), nextID: 1}
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
	service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)
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
	service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)
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
	service := BService{booking: make([]*booking.BookingData, 0), notConfirmed: make([]*booking.BookingData, 0), nextID: 1}

	rspCreate := &booking.CreateBookingResult{}
	service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 500}, rspCreate)

	if rspCreate.Id != -1 {
		t.Errorf("Expected no booking because of too much seats!\n")
	}
}

func TestBookingTwoConfirmed(t *testing.T) {
	service := BService{booking: make([]*booking.BookingData, 0), notConfirmed: make([]*booking.BookingData, 0), nextID: 1}

	rspCreate := &booking.CreateBookingResult{}
	service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 200}, rspCreate)

	rspCreate2 := &booking.CreateBookingResult{}
	service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 2, ShowID: 1, Seats: 400}, rspCreate2)

	rspConfirm := &booking.ConfirmBookingResult{}
	service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)
	if !rspConfirm.Successful {
		t.Errorf("Booking expected successful")
	}

	rspConfirm2:= &booking.ConfirmBookingResult{}
	service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate2.Id}, rspConfirm2)
	if rspConfirm2.Successful {
		t.Errorf("Booking 2 expected not successful")
	}

	rspConfirm3:= &booking.ConfirmBookingResult{}
	service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: 42}, rspConfirm3)
	if rspConfirm3.Successful {
		t.Errorf("Booking 3 expected not successful")
	}
}

func TestBookingDelete(t *testing.T) {
	service := BService{booking: make([]*booking.BookingData, 0), notConfirmed: make([]*booking.BookingData, 0), nextID: 1}
	rsp := &booking.GetListResult{}

	rspCreate := &booking.CreateBookingResult{}
	service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)

	rspConfirm := &booking.ConfirmBookingResult{}
	service.ConfirmBooking(context.TODO(), &booking.ConfirmBookingRequest{Id: rspCreate.Id}, rspConfirm)

	rspDelete := &booking.DeleteBookingResult{}
	service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id:rspCreate.Id}, rspDelete)

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
	service.CreateBooking(context.TODO(), &booking.CreateBookingRequest{UserID: 1, ShowID: 1, Seats: 1}, rspCreate)

	rspDelete = &booking.DeleteBookingResult{}
	service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id:rspCreate.Id}, rspDelete)

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
	service.DeleteBooking(context.TODO(), &booking.DeleteBookingRequest{Id:6000}, rspDelete)
	if rspDelete.Successful{
		t.Error("Expected no booking with this ID!")
	}
}