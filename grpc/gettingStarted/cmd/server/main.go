package main

import (
	"context"
	"fmt"
	"log"
	userpb "mygrpc/gen/go/user/v1"
	"net"

	"google.golang.org/grpc"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
}

func (us *userService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Println("married:", userpb.MaritalStatus_MARITAL_STATUS_MARRIED)
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:          req.Uuid,
			FullName:      "Najam Awan",
			BirthYear:     1900,
			Salary:        4000,
			MaritalStatus: userpb.MaritalStatus_MARITAL_STATUS_MARRIED,
			// MaritalStatus:
		},
	}, nil
}

func main() {
	fmt.Println("listneing on localhost:9879")
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userService{})
	grpcServer.Serve(lis)
}
