package main

import (
	"fmt"
	"log"
	"net"
	"playground-grpc/server"
)

var (
	defaultPort = 8080
)

func main() {
	port := defaultPort
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := server.NewServer()
	grpcServer.Serve(listener)
}
