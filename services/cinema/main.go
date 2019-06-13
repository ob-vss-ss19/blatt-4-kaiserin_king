package main

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/cinema/srv"
)

func main() {
	srv.RunService(context.TODO(), false)
}
