package main

import (
	"log"
	"net"

	pb "github.com/MYK12397/grpc-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", l.Addr())
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
