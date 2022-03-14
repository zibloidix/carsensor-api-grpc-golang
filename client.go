package main

import (
	"context"
	"log"
	"time"

	"github.com/zibloidix/carsensor-api-grpc-golang/carsensorpb"
	"google.golang.org/grpc"
)

var requests = []*carsensorpb.SendPointRequest{
	&carsensorpb.SendPointRequest{
		Car:       101,
		Route:     3000,
		Latitude:  43.5000,
		Longitude: 56.6000,
	},
	&carsensorpb.SendPointRequest{
		Car:       101,
		Route:     3000,
		Latitude:  43.5010,
		Longitude: 56.6010,
	},
	&carsensorpb.SendPointRequest{
		Car:       101,
		Route:     3000,
		Latitude:  43.5020,
		Longitude: 56.6020,
	},
	&carsensorpb.SendPointRequest{
		Car:       101,
		Route:     3000,
		Latitude:  43.5030,
		Longitude: 56.6030,
	},
	&carsensorpb.SendPointRequest{
		Car:       101,
		Route:     3000,
		Latitude:  43.5040,
		Longitude: 56.6040,
	},
}

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := carsensorpb.NewCarSensorServiceClient(cc)
	stream, err := c.SendPoint(context.Background())
	if err != nil {
		log.Fatalf("")
	}
	for _, req := range requests {
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	log.Printf("Response^ %v", resp)
}
