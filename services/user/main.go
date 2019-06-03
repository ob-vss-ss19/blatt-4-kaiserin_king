package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	user "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/user/proto"
)

type UService struct {
	user 	[]*user.UserData
	nextID	int32
}

func (us *UService) CreateUser(ctx context.Context, req *user.CreateUserRequest, rsp *user.CreateUserResult) error {
	givenID := us.nextID
	us.nextID++
	us.user = append(us.user, &user.UserData{Name: req.Name, Id: givenID})
	rsp.Id = givenID
	fmt.Printf("got name: %v", req.Name)

	return nil
}

func (us *UService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, rsp *user.DeleteUserResult) error {
	// delete only if no bookings
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.services.user"),
	)

	service.Init()
	user.RegisterUserHandler(service.Server(), &UService{user: make([]*user.UserData, 0), nextID: 0})
	r := service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}
