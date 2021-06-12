package server

import (
	"context"
	"errors"
	pb "playground-grpc/grpc"
	"time"

	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	deafultMNOU = 4
)

type playgroundGrpc struct {
	Cache *cache.Cache
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	svc := playgroundGrpc{
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	}
	pb.RegisterPlaygroundGrpcServer(grpcServer, &svc)
	reflection.Register(grpcServer)
	return grpcServer
}

func (p *playgroundGrpc) CreateTodo(c context.Context, todo *pb.TodoData) (*pb.TodoData, error) {
	p.Cache.Set(todo.Id, todo, cache.NoExpiration)
	return todo, nil
}

func (p *playgroundGrpc) ReadTodo(c context.Context, todo *pb.TodoData) (*pb.TodoData, error) {
	got, ok := p.Cache.Get(todo.Id)
	if !ok {
		return nil, errors.New("not found")
	}
	todo, ok = got.(*pb.TodoData)
	if !ok {
		return nil, errors.New("unexpected value")
	}
	return todo, nil
}

func (p *playgroundGrpc) UpdateTodo(c context.Context, todo *pb.TodoData) (*pb.TodoData, error) {
	p.Cache.Set(todo.Id, todo, cache.NoExpiration)
	return todo, nil
}

func (p *playgroundGrpc) DeleteTodo(c context.Context, todo *pb.TodoData) (*pb.TodoData, error) {
	p.Cache.Delete(todo.Id)
	return todo, nil
}
