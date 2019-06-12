package main

import (
	"context"
	"testing"

	movie "github.com/ob-vss-ss19/blatt-4-kaiserin_king/services/movie/proto"
)

func TestMovie(t *testing.T) {
	service := MService{movie: make([]*movie.MovieData, 0), nextID: 1}

	rsp := &movie.GetMovieListResult{}
	err := service.GetMovieList(context.TODO(), &movie.GetMovieListRequest{}, rsp)

	if err == nil {
		if len(rsp.Movies) != 0 {
			t.Errorf("Expected empty List")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspCreate := &movie.CreateMovieResult{}
	err = service.CreateMovie(context.TODO(), &movie.CreateMovieRequest{Titel: "Sex and The City"}, rspCreate)

	rsp = &movie.GetMovieListResult{}
	err = service.GetMovieList(context.TODO(), &movie.GetMovieListRequest{}, rsp)

	if err == nil {
		if len(rsp.Movies) != 1 {
			t.Errorf("Expected List with one element")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete := &movie.DeleteMovieResult{}
	err = service.DeleteMovie(context.TODO(), &movie.DeleteMovieRequest{Id: rspCreate.Id}, rspDelete)

	if err == nil {
		if !rspDelete.Successful {
			t.Errorf("Expected successful deleting")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rsp = &movie.GetMovieListResult{}
	err = service.GetMovieList(context.TODO(), &movie.GetMovieListRequest{}, rsp)

	if err == nil {
		if len(rsp.Movies) != 0 {
			t.Errorf("Expected empty List")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

	rspDelete = &movie.DeleteMovieResult{}
	err = service.DeleteMovie(context.TODO(), &movie.DeleteMovieRequest{Id: 42}, rspDelete)

	if err == nil {
		if rspDelete.Successful {
			t.Errorf("Expected faling deleting")
		}
	} else {
		t.Errorf("Error %v\n", err)
	}

}