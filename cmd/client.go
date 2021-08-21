package main

import (
	"context"
	"fmt"
	"net"

	pb "playground-grpc/drpc"

	"storj.io/drpc/drpcconn"
)

func main() {
	rawconn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	conn := drpcconn.New(rawconn)
	defer conn.Close()

	client := pb.NewDRPCPlaygroundDrpcClient(conn)
	ctx := context.Background()
	creatTodo, err := client.CreateTodo(ctx, &pb.TodoData{
		Id:    "1",
		Title: "homework",
		Body:  "math",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(creatTodo)
	readTodo, err := client.ReadTodo(ctx, &pb.TodoData{Id: "1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(readTodo)
	updateTodo, err := client.UpdateTodo(ctx, &pb.TodoData{Title: "updated title", Body: "updated body"})
	if err != nil {
		panic(err)
	}
	fmt.Println(updateTodo)
	deleteTodo, err := client.DeleteTodo(ctx, &pb.TodoData{Id: "1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteTodo)
}
