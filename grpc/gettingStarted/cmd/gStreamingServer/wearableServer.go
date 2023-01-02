package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	wearablepb "mygrpc/gen/go/wearable/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type wearableService struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (w *wearableService) ConsumeBeatsPerMinute(stream wearablepb.WearableService_ConsumeBeatsPerMinuteServer) error {
	var total uint32
	for {
		val, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wearablepb.ConsumeBeatsPerMinuteResponse{
				Total: total,
			})
		}

		if err != nil {
			return err
		}

		fmt.Printf("Min:%v, value:%v, uuid:%v \n", val.GetMinute(), val.GetValue(), val.GetUuid())
		total++
	}
}

func (w *wearableService) BeatsPerMinute(req *wearablepb.BeatsPerMinuteRequest, stream wearablepb.WearableService_BeatsPerMinuteServer) error {
	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("connection breaks")
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)
			value := 30 + rand.Int31n(80)

			if err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			}); err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}
		}
	}
}

func (w *wearableService) CalculateBeatsPerMinute(stream wearablepb.WearableService_CalculateBeatsPerMinuteServer) error {
	var count, total uint32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		total += req.GetValue()

		fmt.Println("Received", req.GetValue())

		count++

		if count%5 == 0 {
			fmt.Println("Total", total, "Sending", float32(total)/5)

			if err := stream.Send(&wearablepb.CalculateBeatsPerMinuteResponse{
				Average: float32(total) / 5,
			}); err != nil {
				return err
			}

			total = 0
		}
	}
}
