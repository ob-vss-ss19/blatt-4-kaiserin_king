// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: movie.proto

package movie

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Movie service

type MovieService interface {
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*DeleteMovieResult, error)
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...client.CallOption) (*CreateMovieResult, error)
	GetMovieList(ctx context.Context, in *GetMovieListRequest, opts ...client.CallOption) (*GetMovieListResult, error)
	Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResult, error)
}

type movieService struct {
	c    client.Client
	name string
}

func NewMovieService(name string, c client.Client) MovieService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "movie"
	}
	return &movieService{
		c:    c,
		name: name,
	}
}

func (c *movieService) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*DeleteMovieResult, error) {
	req := c.c.NewRequest(c.name, "Movie.DeleteMovie", in)
	out := new(DeleteMovieResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...client.CallOption) (*CreateMovieResult, error) {
	req := c.c.NewRequest(c.name, "Movie.CreateMovie", in)
	out := new(CreateMovieResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) GetMovieList(ctx context.Context, in *GetMovieListRequest, opts ...client.CallOption) (*GetMovieListResult, error) {
	req := c.c.NewRequest(c.name, "Movie.GetMovieList", in)
	out := new(GetMovieListResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResult, error) {
	req := c.c.NewRequest(c.name, "Movie.Exist", in)
	out := new(ExistResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Movie service

type MovieHandler interface {
	DeleteMovie(context.Context, *DeleteMovieRequest, *DeleteMovieResult) error
	CreateMovie(context.Context, *CreateMovieRequest, *CreateMovieResult) error
	GetMovieList(context.Context, *GetMovieListRequest, *GetMovieListResult) error
	Exist(context.Context, *ExistRequest, *ExistResult) error
}

func RegisterMovieHandler(s server.Server, hdlr MovieHandler, opts ...server.HandlerOption) error {
	type movie interface {
		DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *DeleteMovieResult) error
		CreateMovie(ctx context.Context, in *CreateMovieRequest, out *CreateMovieResult) error
		GetMovieList(ctx context.Context, in *GetMovieListRequest, out *GetMovieListResult) error
		Exist(ctx context.Context, in *ExistRequest, out *ExistResult) error
	}
	type Movie struct {
		movie
	}
	h := &movieHandler{hdlr}
	return s.Handle(s.NewHandler(&Movie{h}, opts...))
}

type movieHandler struct {
	MovieHandler
}

func (h *movieHandler) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *DeleteMovieResult) error {
	return h.MovieHandler.DeleteMovie(ctx, in, out)
}

func (h *movieHandler) CreateMovie(ctx context.Context, in *CreateMovieRequest, out *CreateMovieResult) error {
	return h.MovieHandler.CreateMovie(ctx, in, out)
}

func (h *movieHandler) GetMovieList(ctx context.Context, in *GetMovieListRequest, out *GetMovieListResult) error {
	return h.MovieHandler.GetMovieList(ctx, in, out)
}

func (h *movieHandler) Exist(ctx context.Context, in *ExistRequest, out *ExistResult) error {
	return h.MovieHandler.Exist(ctx, in, out)
}
