package server

import (
	"context"
	"errors"
	pb "playground-grpc/drpc"
	"time"

	"github.com/patrickmn/go-cache"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

var (
	deafultMNOU = 4
)

type playgroundGrpc struct {
	Cache *cache.Cache
}

func NewServer() *drpcserver.Server {
	m := drpcmux.New()
	svc := playgroundGrpc{
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	}
	err := pb.DRPCRegisterPlaygroundDrpc(m, &svc)
	if err != nil {
		panic(err)
		return nil
	}
	return drpcserver.New(m)
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
