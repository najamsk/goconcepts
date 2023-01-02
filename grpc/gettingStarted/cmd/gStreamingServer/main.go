package main

import (
	"fmt"
	"log"
	wearablepb "mygrpc/gen/go/wearable/v1"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("server doing server side grpc streaming on localhost:9879")
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	wearablepb.RegisterWearableServiceServer(grpcServer, &wearableService{})
	reflection.Register(grpcServer)
	// userpb.RegisterUserServiceServer(grpcServer, &userService{})
	grpcServer.Serve(lis)
}
