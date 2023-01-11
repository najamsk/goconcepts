package main

import (
	"context"
	"fmt"
	"log"

	userpb "mygrpc/gen/go/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("will be connecting to grpc server")
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "1-2-3-4-5-6",
	})
	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	fmt.Printf("user:%+v\n", res.User)
	fmt.Println("---------------------")
	fmt.Printf("Fullname:%v\n", res.User.FullName)
	fmt.Printf("Salary:%v\n", res.User.Salary)
	fmt.Printf("Status:%v\n", res.User.MaritalStatus)
}
