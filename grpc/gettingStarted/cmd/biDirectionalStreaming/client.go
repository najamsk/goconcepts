package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// stream, err := client.ConsumeBeatsPerMinute(context.Background())
	stream, err := client.CalculateBeatsPerMinute(context.Background())

	if err != nil {
		log.Fatalln("Couldn't request", err)
	}

	for i := 0; i < 10; i++ {
		err = stream.Send(&wearablepb.CalculateBeatsPerMinuteRequest{
			Uuid:   "najam awa",
			Value:  uint32(i),
			Minute: uint32(i * 2),
		})

		if err != nil {
			log.Fatalln("stream send error:", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalln("close error:", err)
	}
	// fmt.Println("Total:", resp.GetTotal())
	for {

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("receving  stream error:", err)
		}

		fmt.Println("Average:", res.GetAverage())
	}

}
