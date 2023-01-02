package main

import (
	"context"
	"fmt"
	"io"
	"log"

	wearablepb "mygrpc/gen/go/wearable/v1"

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
	fmt.Println("call grpc streaming method")
	client := wearablepb.NewWearableServiceClient(conn)

	// client := userpb.NewUserServiceClient(conn)

	res, err := client.BeatsPerMinute(context.Background(), &wearablepb.BeatsPerMinuteRequest{Uuid: "mario"})
	if err != nil {
		log.Fatalln("Couldn't request", err)
	}

	go func() {
		for {
			resp, err := res.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalln("Receiving", err)
				return
			}
			fmt.Printf("min: %v, value:%v \n ", resp.Minute, resp.Value)
		}
	}()

	// for {
	// 	resp, err := res.Recv()
	// 	if err == io.EOF {
	// 		return
	// 	}
	// 	if err != nil {
	// 		log.Println("Receiving", err)
	// 		return
	// 	}
	// 	fmt.Printf("min: %v, value:%v \n ", resp.GetMinute(), resp.GetValue())

	// }

	for {
		select {
		case <-res.Context().Done():
			fmt.Println("All done, possible error", res.Context().Err())
			break
		}
	}

}
