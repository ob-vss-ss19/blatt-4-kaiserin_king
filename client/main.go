package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	show "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/show/proto"
)

func main() {


	var client client.Client
	showC := show.NewShowService("go.micro.services.show", client)


	rspS , err := showC.GetShowList(context.TODO(), &show.GetShowListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of all shows: %v \n\n", rspS.Shows)

}
